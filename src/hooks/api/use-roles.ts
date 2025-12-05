import { useQuery } from '@tanstack/react-query'
import { rolesService } from '@/services'

/**
 * Query key factory for roles
 */
export const roleKeys = {
  all: ['roles'] as const,
  list: () => [...roleKeys.all, 'list'] as const,
}

/**
 * Hook to fetch available user roles
 */
export function useRoles() {
  return useQuery({
    queryKey: roleKeys.list(),
    queryFn: async () => {
      const response = await rolesService.getRoles()
      return response.data || []
    },
    staleTime: 60 * 60 * 1000, // 1 hour - roles rarely change
    // Enable persistent cache
    meta: { persist: true },
  })
}
