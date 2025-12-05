import apiClient from '@/lib/api-client'
import type {
  ApiResponse,
  PaginatedResponse,
  User,
  CreateUserInput,
  UpdateUserInput,
} from '@/types/api'

export const usersService = {
  /**
   * Get all users with optional pagination and filters
   */
  async getUsers(params?: {
    page?: number
    pageSize?: number
    status?: string[]
    role?: string[]
    username?: string
  }): Promise<PaginatedResponse<User>> {
    const { data } = await apiClient.get<PaginatedResponse<User>>('/users', {
      params,
    })
    return data
  },

  /**
   * Get a single user by ID
   */
  async getUser(userId: string): Promise<ApiResponse<User>> {
    const { data } = await apiClient.get<ApiResponse<User>>(`/users/${userId}`)
    return data
  },

  /**
   * Create a new user
   */
  async createUser(input: CreateUserInput): Promise<ApiResponse<User>> {
    const { data } = await apiClient.post<ApiResponse<User>>('/users', input)
    return data
  },

  /**
   * Update an existing user
   */
  async updateUser(
    userId: string,
    input: UpdateUserInput
  ): Promise<ApiResponse<User>> {
    const { data } = await apiClient.put<ApiResponse<User>>(
      `/users/${userId}`,
      input
    )
    return data
  },

  /**
   * Delete a user
   */
  async deleteUser(userId: string): Promise<ApiResponse> {
    const { data } = await apiClient.delete<ApiResponse>(`/users/${userId}`)
    return data
  },

  /**
   * Bulk delete users
   */
  async deleteUsers(userIds: string[]): Promise<ApiResponse> {
    const { data } = await apiClient.post<ApiResponse>('/users/bulk-delete', {
      userIds,
    })
    return data
  },

  /**
   * Update user role
   */
  async updateUserRole(
    userId: string,
    roleId: number
  ): Promise<ApiResponse<User>> {
    const { data } = await apiClient.patch<ApiResponse<User>>(
      `/users/${userId}/role`,
      { roleId }
    )
    return data
  },
}
