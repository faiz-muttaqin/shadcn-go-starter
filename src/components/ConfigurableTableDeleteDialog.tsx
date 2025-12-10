import { useState } from 'react'
import { AlertTriangle } from 'lucide-react'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { ConfirmDialog } from '@/components/ConfirmDialog'
import type { TableSettings } from '@/types/auth'
import {apiClient} from '@/lib/api/client'

type ConfigurableTableDeleteDialogProps = {
    open: boolean
    onOpenChange: (open: boolean) => void
    tableConfig: TableSettings
    rowData: Record<string, string | number | Record<string, unknown>>
    onDelete?: () => void
}

export function ConfigurableTableDeleteDialog({
    open,
    onOpenChange,
    tableConfig,
    rowData,
    onDelete,
}: ConfigurableTableDeleteDialogProps) {
    const [confirmValue, setConfirmValue] = useState('')
    const [isDeleting, setIsDeleting] = useState(false)
    // Try common name fields
    const nameFields = ['name', 'username', 'email', 'title', 'id', 'ID', 'Id']
    // Find a unique identifier column for confirmation (prefer id, then first visible column)
    const confirmColumn =
        // cari berdasarkan urutan priority
        nameFields
            .map(key => tableConfig.column.find(col => col.data === key))
            .find(Boolean)
        // fallback: kolom yang visible pertama
        || tableConfig.column.find(col => col.visible);

    const confirmFieldValue = confirmColumn
        ? String(rowData[confirmColumn.data] || '')
        : ''

    // Get display name for the record
    const getDisplayName = (): string => {

        for (const field of nameFields) {
            const value = rowData[field]
            if (value) {
                if (typeof value === 'string' || typeof value === 'number') {
                    return String(value)
                }
                if (typeof value === 'object' && 'name' in value) {
                    return String(value.name)
                }
            }
        }
        return 'this record'
    }

    const handleDelete = async () => {
        if (confirmValue.trim() !== confirmFieldValue) return

        setIsDeleting(true)
        try {
            // Assuming DELETE request with ID in URL
            const id = rowData.id || rowData.ID
            const deleteUrl = id ? `${tableConfig.url}/${id}` : tableConfig.url

            await apiClient.delete(deleteUrl, { data: rowData })

            if (onDelete) onDelete()
            onOpenChange(false)
            setConfirmValue('')
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : 'Failed to delete'
            alert(errorMessage)
        } finally {
            setIsDeleting(false)
        }
    }

    const handleOpenChange = (newOpen: boolean) => {
        if (!newOpen) {
            setConfirmValue('')
        }
        onOpenChange(newOpen)
    }

    return (
        <ConfirmDialog
            open={open}
            onOpenChange={handleOpenChange}
            handleConfirm={handleDelete}
            disabled={confirmValue.trim() !== confirmFieldValue || isDeleting}
            title={
                <span className='text-destructive'>
                    <AlertTriangle
                        className='stroke-destructive me-1 inline-block'
                        size={18}
                    />{' '}
                    Delete {tableConfig.table_name}
                </span>
            }
            desc={
                <div className='space-y-4'>
                    <p className='mb-2'>
                        Are you sure you want to delete{' '}
                        <span className='font-bold'>{getDisplayName()}</span>?
                        <br />
                        This action will permanently remove this record from the system.
                        This cannot be undone.
                    </p>

                    {confirmColumn && (
                        <Label className='my-2'>
                            {confirmColumn.name}:
                            <Input
                                value={confirmValue}
                                onChange={(e) => setConfirmValue(e.target.value)}
                                placeholder={`Enter ${confirmColumn.name.toLowerCase()} to confirm deletion`}
                                disabled={isDeleting}
                            />
                        </Label>
                    )}

                    <Alert variant='destructive'>
                        <AlertTitle>Warning!</AlertTitle>
                        <AlertDescription>
                            Please be careful, this operation cannot be rolled back.
                        </AlertDescription>
                    </Alert>
                </div>
            }
            confirmText={isDeleting ? 'Deleting...' : 'Delete'}
            destructive
        />
    )
}
