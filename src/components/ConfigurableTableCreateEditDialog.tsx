import { useState, useEffect } from 'react'
import { Button } from '@/components/ui/button'
import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { useFilterOptions,type FilterOption } from '@/hooks/use-filter-options'
import {apiClient} from '@/lib/api/client'
import type { TableColumn, TableSettings } from '@/types/auth'

type ConfigurableTableEditDialogProps = {
    open: boolean
    onOpenChange: (open: boolean) => void
    tableConfig: TableSettings
    rowData: Record<string, string | number | Record<string, unknown>> | null
    onSave?: () => void
}

export function ConfigurableTableCreateEditDialog({ 
    open, 
    onOpenChange, 
    tableConfig, 
    rowData, 
    onSave 
}: ConfigurableTableEditDialogProps) {
    const [formData, setFormData] = useState<Record<string, string | number | Record<string, unknown>>>(rowData || {})
    const [isSubmitting, setIsSubmitting] = useState(false)

    useEffect(() => {
        setFormData(rowData || {})
    }, [rowData])

    const handleChange = (key: string, value: string | number | Record<string, unknown>) => {
        setFormData((prev: Record<string, string | number | Record<string, unknown>>) => ({ ...prev, [key]: value }))
    }

    // Helper function to extract display value from data
    const getDisplayValue = (value: string | number | Record<string, unknown>): string | number => {
        if (value == null) return ''
        if (typeof value === 'string' || typeof value === 'number') return value
        if (typeof value === 'object' && !Array.isArray(value)) {
            // Extract value from object (e.g., role.name, role.id)
            const obj = value as Record<string, unknown>
            if ('name' in obj && (typeof obj.name === 'string' || typeof obj.name === 'number')) return obj.name
            if ('title' in obj && (typeof obj.title === 'string' || typeof obj.title === 'number')) return obj.title
            if ('label' in obj && (typeof obj.label === 'string' || typeof obj.label === 'number')) return obj.label
            if ('id' in obj && (typeof obj.id === 'string' || typeof obj.id === 'number')) return obj.id
        }
        return String(value)
    }

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        setIsSubmitting(true)
        try {
            if (rowData === null) {
                // Create new record
                await apiClient.post(tableConfig.url, formData)
            } else {
                // Update existing record
                await apiClient.put(tableConfig.url, formData)
            }
            if (onSave) onSave()
            onOpenChange(false)
        } catch (error) {
            // Handle error - you might want to show a toast notification here
            const errorMessage = error instanceof Error ? error.message : 'Failed to update'
            alert(errorMessage) // Replace with your preferred error notification system
        } finally {
            setIsSubmitting(false)
        }
    }

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="sm:max-w-lg max-h-[80vh]">
                <DialogHeader>
                    <DialogTitle>Edit {tableConfig.table_name}</DialogTitle>
                </DialogHeader>
                <form onSubmit={handleSubmit} className="space-y-4 overflow-y-auto px-1">
                    {tableConfig.column
                        .filter((col) => col.visible)
                        .map((col) => {
                            const rawValue = formData[col.data]
                            const displayValue = getDisplayValue(rawValue)
                            const editable = !!col.editable

                            // Selection dropdown
                            if (col.selection && col.selection !== '') {
                                return <SelectionField key={col.data} col={col} value={displayValue} editable={editable} onChange={handleChange} />
                            }

                            // Text/number/email/datetime input
                            let inputType = 'text'
                            if (col.type === 'number') inputType = 'number'
                            if (col.type === 'email') inputType = 'email'
                            if (col.type === 'datetime') inputType = 'datetime-local'

                            return (
                                <div key={col.data} className="grid grid-cols-4 items-center gap-4">
                                    <Label htmlFor={col.data} className="text-right">
                                        {col.name}
                                    </Label>
                                    <Input
                                        id={col.data}
                                        type={inputType}
                                        value={displayValue}
                                        onChange={(e) => handleChange(col.data, e.target.value)}
                                        disabled={!editable}
                                        className="col-span-3"
                                    />
                                </div>
                            )
                        })}
                    <DialogFooter>
                        <Button type="button" variant="outline" onClick={() => onOpenChange(false)}>
                            Cancel
                        </Button>
                        <Button type="submit" disabled={isSubmitting}>
                            {isSubmitting ? 'Saving...' : 'Save changes'}
                        </Button>
                    </DialogFooter>
                </form>
            </DialogContent>
        </Dialog>
    )
}

// Helper component to fetch and render selection dropdown
function SelectionField({ 
    col, 
    value, 
    editable, 
    onChange 
}: { 
    col: TableColumn
    value: string | number | Record<string, unknown>
    editable: boolean
    onChange: (key: string, value: string | number | Record<string, unknown>) => void
}) {
    const { data: optionsData, isLoading } = useFilterOptions(col.selection!)
    
    const items = optionsData?.options.map((opt: FilterOption) => ({
        label: opt.label,
        value: opt.value,
    })) || []

    // Ensure value is a string for the Select component
    const currentValue = String(value || '')

    return (
        <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor={col.data} className="text-right">
                {col.name}
            </Label>
            <Select
                value={currentValue}
                onValueChange={(v: string) => onChange(col.data, v)}
                disabled={!editable || isLoading}
            >
                <SelectTrigger className="col-span-3">
                    <SelectValue placeholder={isLoading ? 'Loading...' : `Select ${col.name}`} />
                </SelectTrigger>
                <SelectContent>
                    {items.map((item: { label: string; value: string }) => (
                        <SelectItem key={item.value} value={item.value}>
                            {item.label}
                        </SelectItem>
                    ))}
                </SelectContent>
            </Select>
        </div>
    )
}
