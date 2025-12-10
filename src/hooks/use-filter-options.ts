import { useQuery } from '@tanstack/react-query'
import {apiClient} from '@/lib/api/client'
import { type IconName } from 'lucide-react/dynamic'

export type FilterOption = {
  label: string
  value: string
  icon?: IconName
}

export type FilterOptionsResponse = {
  data: string
  message: string
  options: FilterOption[]
  success: boolean
  title: string
}

export function useFilterOptions(selectionUrl: string) {
  return useQuery({
    queryKey: ['filter-options', selectionUrl],
    queryFn: async () => {
      const response = await apiClient.get<FilterOptionsResponse>(selectionUrl)
      return response.data
    },
    enabled: !!selectionUrl,
    staleTime: 5 * 60 * 1000, // Cache for 5 minutes
  })
}
