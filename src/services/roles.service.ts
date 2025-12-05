import apiClient from '@/lib/api-client'
import type { UserRole, ApiResponse } from '@/types/api'

export const rolesService = {
  /**
   * Get available user roles
   */
  async getRoles(): Promise<ApiResponse<UserRole[]>> {
    const { data } = await apiClient.get<ApiResponse<UserRole[]>>('/users/roles')
    return data
  },
}
