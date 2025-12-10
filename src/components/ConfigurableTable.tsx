import { useEffect, useState, useMemo } from 'react'
import {
    type ColumnDef,
    type SortingState,
    type VisibilityState,
    flexRender,
    getCoreRowModel,
    getFacetedRowModel,
    getFacetedUniqueValues,
    getFilteredRowModel,
    getPaginationRowModel,
    getSortedRowModel,
    useReactTable,
} from '@tanstack/react-table'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuShortcut,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { cn } from '@/lib/utils'
import { type NavigateFn } from '@/hooks/use-table-url-state'
import { Trash2, UserPen,RefreshCw, AlertCircle } from 'lucide-react'
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/ui/table'
import { DataTablePagination, DataTableToolbar } from '@/components/data-table'
import { Button } from '@/components/ui/button'
import { StorageHelper } from '@/lib/api/storage'
import type { TableSettings, TableColumn } from '@/types/auth'
import { useQuery, useQueries } from '@tanstack/react-query'
import {apiClient} from '@/lib/api/client'
import { type FilterOption } from '@/hooks/use-filter-options'
import { type IconName } from 'lucide-react/dynamic'
import { Checkbox } from './ui/checkbox'
import { ConfigurableTableCreateEditDialog } from './ConfigurableTableCreateEditDialog'
import { ConfigurableTableDeleteDialog } from './ConfigurableTableDeleteDialog'

type ConfigurableTableProps = {
    tableName: string // e.g., "users" (without "t_" prefix)
    search?: Record<string, unknown>
    navigate?: NavigateFn
    mode?: 'client' | 'server' // Default: client
}

// Helper: Generate TanStack Table columns from config
function generateColumns(
    tableConfig: TableSettings | null,
    onEdit?: (row: Record<string, unknown>) => void,
    onDelete?: (row: Record<string, unknown>) => void
): ColumnDef<Record<string, unknown>>[] {
    if (!tableConfig) return []
    const baseColumns: ColumnDef<Record<string, unknown>>[] = tableConfig?.column
        .filter((col) => col.visible && col.data)
        .map((col) => ({
            accessorKey: col.data,
            header: col.name,
            enableSorting: col.sortable,
            enableColumnFilter: col.filterable,
            cell: (ctx) => {
                const value = ctx.getValue()
                // ...existing code...
                if (value == null) return '-'
                if (col.type === 'datetime') return new Date(value as string).toLocaleString()
                if (col.type === 'email') return <a href={`mailto:${value}`} className="text-blue-500 hover:underline">{String(value)}</a>
                if (col.type === 'image') return <img src={String(value)} alt={col.name} className="h-10 w-10 rounded object-cover" />
                if (col.type === 'avatar') return <img src={String(value)} alt={col.name} className="size-8 rounded-full object-cover" />
                if (col.type === 'badge') {
                    const badgeValue = String(value).toLowerCase()
                    const badgeColors: Record<string, string> = {
                        active: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300',
                        inactive: 'bg-gray-100 text-gray-800 dark:bg-gray-800 dark:text-gray-300',
                        pending: 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300',
                        suspended: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300',
                        invited: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300',
                        success: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300',
                        warning: 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300',
                        error: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300',
                        info: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300',
                    }
                    const colorClass = badgeColors[badgeValue] || 'bg-gray-100 text-gray-800 dark:bg-gray-800 dark:text-gray-300'
                    return <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${colorClass}`}>{String(value)}</span>
                }
                if (col.passwordable) return '••••••••'
                if (typeof value === 'object' && value !== null) {
                    if (Array.isArray(value)) return `${value.length} items`
                    const obj = value as Record<string, unknown>
                    if (obj.icon && obj.title && obj.name) {
                        return <span className="inline-flex items-center gap-1"><i className={obj.icon as string}></i><span>{String(obj.title)}</span></span>
                    }
                    if (obj.name) return String(obj.name)
                    if (obj.role_name) return String(obj.role_name)
                    if (obj.alias) return String(obj.alias)
                    if (obj.title) return String(obj.title)
                    if (obj.label) return String(obj.label)
                    const json = JSON.stringify(value)
                    return json.length > 50 ? `${json.substring(0, 47)}...` : json
                }
                return String(value)
            },
        }))

    // Add checkbox column if checkable
    let finalColumns = [...baseColumns]
    if (tableConfig?.checkable) {
        finalColumns = [
            {
                id: 'select',
                header: ({ table }) => (
                    <Checkbox
                        checked={
                            table.getIsAllPageRowsSelected() ||
                            (table.getIsSomePageRowsSelected() && 'indeterminate')
                        }
                        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
                        aria-label='Select all'
                        className='translate-y-[2px]'
                    />
                ),
                meta: {
                    className: cn('max-md:sticky start-0 z-10 rounded-tl-[inherit]'),
                },
                cell: ({ row }) => (
                    <Checkbox
                        checked={row.getIsSelected()}
                        onCheckedChange={(value) => row.toggleSelected(!!value)}
                        aria-label='Select row'
                        className='translate-y-[2px]'
                    />
                ),
                enableSorting: false,
                enableHiding: false,
            },
            ...finalColumns,
        ]
    }

    // Add actions column if editable or deletable
    if (tableConfig?.editable || tableConfig?.deletable) {
        finalColumns = [
            ...finalColumns,
            {
                id: 'actions',
                cell: ({ row }) => (
                    <>
                        <DropdownMenu modal={false}>
                            <DropdownMenuTrigger asChild>
                                <Button
                                    variant='ghost'
                                    className='data-[state=open]:bg-muted flex h-8 w-8 p-0'
                                >
                                    <DotsHorizontalIcon className='h-4 w-4' />
                                    <span className='sr-only'>Open menu</span>
                                </Button>
                            </DropdownMenuTrigger>
                            <DropdownMenuContent align='end' className='w-[160px]'>
                                {tableConfig.editable && (
                                    <>
                                        <DropdownMenuItem
                                            onClick={() => {
                                                if (onEdit) onEdit(row.original)
                                            }}
                                        >
                                            Edit
                                            <DropdownMenuShortcut>
                                                <UserPen size={16} />
                                            </DropdownMenuShortcut>
                                        </DropdownMenuItem>
                                        <DropdownMenuSeparator />
                                    </>
                                )}
                                {tableConfig.deletable && (
                                    <DropdownMenuItem
                                        onClick={() => {
                                            if (onDelete) onDelete(row.original)
                                        }}
                                        className='text-red-500!'
                                    >
                                        Delete
                                        <DropdownMenuShortcut>
                                            <Trash2 size={16} />
                                        </DropdownMenuShortcut>
                                    </DropdownMenuItem>
                                )}
                            </DropdownMenuContent>
                        </DropdownMenu>
                    </>
                ),
                enableSorting: false,
                enableHiding: false,
            },
        ]
    }
    return finalColumns

}

// Helper: Parse sort string "id desc" to TanStack format
function parseSortString(sortString: string): SortingState {
    const [id, direction] = sortString.split(' ')
    return [{ id, desc: direction?.toLowerCase() === 'desc' }]
}

// Component to fetch and manage filter options for a single column
function useColumnFilters(columns: TableColumn[]) {
    // Get all columns that need selection filters (filterable + selection URL exists)
    const selectionColumns = useMemo(
        () => columns.filter((col) => col.filterable && col.selection && col.selection !== ''),
        [columns]
    )

    // Use react-query's useQueries to fetch options for each selection column
    const queries = useQueries({
        queries: selectionColumns.map((col) => ({
            queryKey: ['filter-options', col.selection],
            queryFn: async () => {
                const response = await apiClient.get(col.selection!)
                return response.data
            },
            enabled: !!col.selection,
            staleTime: 5 * 60 * 1000,
        })),
    })

    // Build filters array from fetched data (zip selectionColumns with queries)
    const filters = selectionColumns
        .map((col, idx) => {
            const q = queries[idx]
            if (!q) return null
            if (!q.isSuccess || !q.data) return null
            return {
                data: col.data,
                title: q.data!.title || col.name,
                options: q.data!.options.map((opt: FilterOption) => ({
                    label: opt.label,
                    value: opt.value,
                    icon: opt.icon as IconName | undefined,
                })),
            }
        })
        .filter(Boolean) as {
        data: string
        title: string
        options: { label: string; value: string; icon?: IconName }[]
    }[]

    // Check if any queries are still loading
    const isLoadingFilters = queries.some((q) => q.isLoading)

    // Get text search columns (filterable but no selection URL)
    const textSearchColumns = useMemo(
        () => columns.filter((col) => col.filterable && (!col.selection || col.selection === '')),
        [columns]
    )

    return { filters, isLoadingFilters, textSearchColumns }
}

export function ConfigurableTable({
    tableName,
    search = {},
    navigate: _navigate,
    mode = 'client'
}: ConfigurableTableProps) {
    // Load table config from localStorage (stored under key `t_<tableName>`)
    const rawTable = StorageHelper.getTableData(`t_${tableName}`)
    const tableConfig: TableSettings | null = rawTable ? (rawTable as TableSettings) : null

    // Edit dialog state
    const [createEditDialogOpen, setCreateEditDialogOpen] = useState(false)
    const [currentEditRow, setCurrentEditRow] = useState<Record<string, string | number | Record<string, unknown>> | null>(null)

    // Delete dialog state
    const [deleteDialogOpen, setDeleteDialogOpen] = useState(false)
    const [currentDeleteRow, setCurrentDeleteRow] = useState<Record<string, string | number | Record<string, unknown>> | null>(null)



    const handleCreate = () => {
        setCurrentEditRow(null)
        setCreateEditDialogOpen(true)
    }
    const handleEdit = (row: Record<string, unknown>) => {
        // console.log('Edit row:', row)
        setCurrentEditRow(row as Record<string, string | number | Record<string, unknown>> | null)
        setCreateEditDialogOpen(true)
    }

    const handleDelete = (row: Record<string, unknown>) => {
        setCurrentDeleteRow(row as Record<string, string | number | Record<string, unknown>> | null)
        setDeleteDialogOpen(true)
    }

    // Get dynamic filters from table config
    const { filters, isLoadingFilters, textSearchColumns } = useColumnFilters(tableConfig?.column || [])

    // Determine search configuration
    // If there's only one text search column, use it as searchKey
    // Otherwise, use global search (searchKey = undefined)
    // const searchKey = textSearchColumns.length > 0 ? textSearchColumns[0].data : ""
    const searchKey =
        textSearchColumns.length > 0
            ? textSearchColumns
                .filter(col => col.visible)       // keep only visible ones
                .map(col => col.data)             // then extract the data
            : [""]
    // console.log("textSearchColumns:", textSearchColumns)
    // console.log("isLoadingFilters:", isLoadingFilters)
    // console.log("SEARCH filters:", filters)
    // const searchPlaceholder = searchKey
    //     ? `Search by ${textSearchColumns[0].name}...`
    //     : `Search ${tableConfig?.table_name || 'table'}...`

    // Local UI states
    const [rowSelection, setRowSelection] = useState({})
    const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({})
    const [sorting, setSorting] = useState<SortingState>(() =>
        tableConfig?.sort ? parseSortString(tableConfig.sort) : []
    )
    const [pagination, setPagination] = useState({
        pageIndex: 0,
        pageSize: tableConfig?.row || 10,
    })
    const [globalFilter, setGlobalFilter] = useState('')
    const [drawCounter, setDrawCounter] = useState(1)

    // Build query key based on mode
    const queryKey = [
        'table',
        mode,
        tableName,
        tableConfig?.url,
        // Server-specific dependencies
        ...(mode === 'server' ? [
            pagination.pageIndex,
            pagination.pageSize,
            sorting,
            globalFilter,
            drawCounter
        ] : [])
    ]

    // Data fetching (both client and server modes)
    const { data: fetchedData, isLoading, error, refetch } = useQuery({
        queryKey,
        queryFn: async () => {
            if (!tableConfig?.url) return null

            if (mode === 'client') {
                // Client mode: fetch all data once
                const response = await apiClient.get(tableConfig.url)
                return response.data
            } else {
                // Server mode: fetch with query params
                const params = new URLSearchParams()

                // DataTables format params
                params.append('draw', String(drawCounter))
                params.append('start', String(pagination.pageIndex * pagination.pageSize))
                params.append('length', String(pagination.pageSize))
                params.append('search[value]', globalFilter)

                // Sorting
                if (sorting.length > 0) {
                    const sortCol = sorting[0]
                    const colIndex = tableConfig.column.findIndex(c => c.data === sortCol.id)
                    if (colIndex !== -1) {
                        params.append('order[0][column]', String(colIndex))
                        params.append('order[0][dir]', sortCol.desc ? 'desc' : 'asc')
                    }
                }

                // Column search filters
                tableConfig.column.forEach((col, index) => {
                    params.append(`columns[${index}][data]`, col.data)
                    // Add search value if exists in URL search params
                    const searchValue = search[col.data]
                    if (searchValue) {
                        params.append(`columns[${index}][search][value]`, String(searchValue))
                    }
                })

                const response = await apiClient.get(`${tableConfig.url}?${params.toString()}`)
                return response.data
            }
        },
        enabled: !!tableConfig?.url,
    })

    // Increment draw counter when server-side params change
    useEffect(() => {
        if (mode === 'server') {
            setDrawCounter(prev => prev + 1)
        }
    }, [pagination.pageIndex, pagination.pageSize, sorting, globalFilter, mode])

    // Reset draw counter when switching modes
    useEffect(() => {
        setDrawCounter(1)
    }, [mode, tableName])

    // Generate columns from config
    const columns = generateColumns(tableConfig, handleEdit, handleDelete)

    // Data source based on mode
    const data = mode === 'server'
        ? (fetchedData?.data || [])
        : (fetchedData?.data || [])

    const table = useReactTable({
        data,
        columns,
        state: {
            sorting,
            pagination,
            rowSelection,
            columnVisibility,
            globalFilter,
        },
        enableRowSelection: true,
        manualPagination: mode === 'server',
        manualSorting: mode === 'server',
        manualFiltering: mode === 'server',
        pageCount: mode === 'server' ? Math.ceil((fetchedData?.recordsFiltered || 0) / pagination.pageSize) : undefined,
        onPaginationChange: setPagination,
        onRowSelectionChange: setRowSelection,
        onSortingChange: setSorting,
        onColumnVisibilityChange: setColumnVisibility,
        onGlobalFilterChange: setGlobalFilter,
        getCoreRowModel: getCoreRowModel(),
        getFilteredRowModel: mode === 'client' ? getFilteredRowModel() : undefined,
        getPaginationRowModel: mode === 'client' ? getPaginationRowModel() : undefined,
        getSortedRowModel: mode === 'client' ? getSortedRowModel() : undefined,
        getFacetedRowModel: getFacetedRowModel(),
        getFacetedUniqueValues: getFacetedUniqueValues(),
    })

    // Update pagination when config changes
    useEffect(() => {
        if (tableConfig?.row && table.getState().pagination.pageSize !== tableConfig.row) {
            table.setPageSize(tableConfig.row)
        }
    }, [tableConfig?.row, table])

    // Config not found - show error
    if (!tableConfig) {
        return (
            <div className="flex flex-col items-center justify-center h-64 gap-4">
                <AlertCircle className="size-12 text-destructive" />
                <div className="text-center">
                    <h3 className="text-lg font-semibold">Table Configuration Not Found</h3>
                    <p className="text-muted-foreground text-sm mt-1">
                        Configuration for table "<code className="font-mono">{tableName}</code>" is not available.
                    </p>
                    <p className="text-muted-foreground text-sm">
                        Please sign in again to load table settings.
                    </p>
                </div>
                <Button onClick={() => window.location.reload()} variant="outline">
                    <RefreshCw className="mr-2 size-4" />
                    Reload Page
                </Button>
            </div>
        )
    }

    // Loading state
    if (mode === 'server' && isLoading) {
        return (
            <div className="flex items-center justify-center h-64">
                <RefreshCw className="size-8 animate-spin text-muted-foreground" />
            </div>
        )
    }

    // Error state
    if (mode === 'server' && error) {
        return (
            <div className="flex flex-col items-center justify-center h-64 gap-4">
                <AlertCircle className="size-12 text-destructive" />
                <div className="text-center">
                    <h3 className="text-lg font-semibold">Failed to Load Data</h3>
                    <p className="text-muted-foreground text-sm mt-1">
                        {error instanceof Error ? error.message : 'Unknown error occurred'}
                    </p>
                </div>
                <Button onClick={() => refetch()} variant="outline">
                    <RefreshCw className="mr-2 size-4" />
                    Retry
                </Button>
            </div>
        )
    }

    return (
        <div className={cn('flex flex-1 flex-col gap-4')}>
            <DataTableToolbar
                table={table}
                searchPlaceholder="Search..."
                searchKey={searchKey}
                filters={filters}
                onCreate={() => handleCreate()}
            />

            <div className="overflow-hidden rounded-md border">
                <Table>
                    <TableHeader>
                        {table.getHeaderGroups().map((headerGroup) => (
                            <TableRow key={headerGroup.id} className="group/row">
                                {headerGroup.headers.map((header) => (
                                    <TableHead
                                        key={header.id}
                                        colSpan={header.colSpan}
                                        className="bg-background group-hover/row:bg-muted"
                                    >
                                        {header.isPlaceholder
                                            ? null
                                            : flexRender(header.column.columnDef.header, header.getContext())}
                                    </TableHead>
                                ))}
                            </TableRow>
                        ))}
                    </TableHeader>
                    <TableBody>
                        {table.getRowModel().rows?.length ? (
                            table.getRowModel().rows.map((row) => (
                                <TableRow
                                    key={row.id}
                                    data-state={row.getIsSelected() && 'selected'}
                                    className="group/row"
                                >
                                    {row.getVisibleCells().map((cell) => (
                                        <TableCell
                                            key={cell.id}
                                            className="bg-background group-hover/row:bg-muted group-data-[state=selected]/row:bg-muted"
                                        >
                                            {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                        </TableCell>
                                    ))}
                                </TableRow>
                            ))
                        ) : (
                            <TableRow>
                                <TableCell colSpan={columns.length} className="h-24 text-center">
                                    No results.
                                </TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </div>

            <DataTablePagination
                table={table}
                pageSizeOptions={tableConfig.row_opt}
                className="mt-auto"
            />

            {mode === 'server' && (
                <div className="text-xs text-muted-foreground">
                    Showing {fetchedData?.recordsFiltered || 0} of {fetchedData?.recordsTotal || 0} total records
                    {isLoadingFilters && ' • Loading filters...'}
                </div>
            )}

            {mode === 'client' && isLoadingFilters && (
                <div className="text-xs text-muted-foreground">
                    Loading filters...
                </div>
            )}

            <ConfigurableTableCreateEditDialog
                open={createEditDialogOpen}
                onOpenChange={setCreateEditDialogOpen}
                tableConfig={tableConfig}
                rowData={currentEditRow}
                onSave={() => refetch()}
            />

            <ConfigurableTableDeleteDialog
                open={deleteDialogOpen}
                onOpenChange={setDeleteDialogOpen}
                tableConfig={tableConfig}
                rowData={currentDeleteRow || {}}
                onDelete={() => refetch()}
            />
        </div>
    )
}
