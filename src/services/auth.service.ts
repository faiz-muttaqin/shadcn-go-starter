import apiClient from '@/lib/api-client'
import type { LoginInput, LoginResponse, ApiResponse } from '@/types/api'

export const authService = {
  /**
   * Login with email and password
   */
  async login(input: LoginInput): Promise<LoginResponse> {
    const { data } = await apiClient.post<LoginResponse>('/auth/login', input)
    return data
  },

  /**
   * Logout current user
   */
  async logout(): Promise<ApiResponse> {
    const { data } = await apiClient.post<ApiResponse>('/auth/logout')
    return data
  },

  /**
   * Refresh access token
   */
  async refreshToken(): Promise<ApiResponse<{ accessToken: string }>> {
    const { data } = await apiClient.post<ApiResponse<{ accessToken: string }>>(
      '/auth/refresh'
    )
    return data
  },
}
