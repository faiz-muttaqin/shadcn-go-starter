import React from 'react'
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select'
import { Cross2Icon, PlusIcon } from '@radix-ui/react-icons'
import { type Table } from '@tanstack/react-table'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { DataTableFacetedFilter } from './faceted-filter'
import { DataTableViewOptions } from './view-options'

type DataTableToolbarProps<TData> = {
  table: Table<TData>
  searchPlaceholder?: string
  searchKey?: string | string[] // ✅ Accept string or array of strings
  onCreate?: () => void
  filters?: {
    data: string
    title: string
    options: {
      label: string
      value: string
      icon?: React.ComponentType<{ className?: string }> | string
    }[]
  }[]
}

export function DataTableToolbar<TData>({
  table,
  searchPlaceholder = 'Filter...',
  searchKey,
  filters = [],
  onCreate,
}: DataTableToolbarProps<TData>) {
  const isFiltered =
    table.getState().columnFilters.length > 0 || table.getState().globalFilter

  // --- Dynamic handling for array of search keys ---
  const [activeSearchKey, setActiveSearchKey] = React.useState<string>(
    Array.isArray(searchKey)
      ? searchKey[0] ?? '*'
      : searchKey ?? '*'
  )

  const handleSearchChange = (value: string) => {
    if (activeSearchKey === '*' || activeSearchKey === '') {
      table.setGlobalFilter(value)
    } else {
      table.getColumn(activeSearchKey)?.setFilterValue(value)
    }
  }

  const currentValue =
    activeSearchKey === '*' || activeSearchKey === ''
      ? (table.getState().globalFilter as string) ?? ''
      : ((table.getColumn(activeSearchKey)?.getFilterValue() as string) ?? '')

  return (
    <div className='flex items-center justify-between'>
      <div className='flex flex-1 flex-col-reverse items-start gap-y-2 sm:flex-row sm:items-center sm:space-x-2'>
        {/* --- Search Key Dropdown if Array --- */}
        {Array.isArray(searchKey) && (
          <Select
            value={activeSearchKey}
            onValueChange={(v) => {
              setActiveSearchKey(v)
              // Reset filters when changing key
              table.setGlobalFilter('')
              table.resetColumnFilters()
            }}
          >
            <SelectTrigger className='h-8 w-[120px]'>
              <SelectValue placeholder='Select Field' />
            </SelectTrigger>
            <SelectContent>
              {/* Add special global option */}
              <SelectItem value='*'>All Columns</SelectItem>
              {searchKey.map((key) => (
                <SelectItem key={key} value={key}>
                  {key}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        )}

        {/* --- Search Input --- */}
        <Input
          placeholder={searchPlaceholder}
          value={currentValue}
          onChange={(event) => handleSearchChange(event.target.value)}
          className='h-8 w-[150px] lg:w-[250px]'
        />
        <div className='flex gap-x-2'>
          {filters.map((filter) => {
            const column = table.getColumn(filter.data)
            if (!column) return null
            return (
              <DataTableFacetedFilter
                key={filter.data}
                column={column}
                title={filter.title}
                options={filter.options}
              />
            )
          })}
        </div>
        {isFiltered && (
          <Button
            variant='ghost'
            onClick={() => {
              table.resetColumnFilters()
              table.setGlobalFilter('')
            }}
            className='h-8 px-2 lg:px-3'
          >
            Reset
            <Cross2Icon className='ms-2 h-4 w-4' />
          </Button>
        )}
      </div>
      <DataTableViewOptions table={table} />
      {onCreate && (
        <Button variant='default'
          onClick={() => {
            onCreate?.()
          }}
          className='ml-2 h-8 w-8 px-2 lg:px-3'>
          <PlusIcon />
        </Button>

      )}
    </div>
  )
}
