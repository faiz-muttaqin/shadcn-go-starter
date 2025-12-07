import { apiClient } from '@/lib/api/client'
import type { User } from '@/types/auth'

/**
 * Fetch user data and table settings from backend after Clerk auth
 */
export async function syncUserData(): Promise<User | null> {
  try {
    // The new ApiClient returns an ApiResponse<T> shape directly.
    const response = await apiClient.get<User>('/auth/login')

    if (response.success) {
      // Table settings are now handled by the ApiClient/storage layer; just
      // return the synced user object here.
      return response.data ?? null
    }

    return null
  } catch (error) {
    if (import.meta.env.DEV) {
      console.error('Failed to sync user data:', error)
    }
    return null
  }
}

/**
 * Save table settings to localStorage
 * Merges with existing settings
 */
// Note: table settings are stored and managed by the ApiClient/storage
// module (see `src/lib/api/storage.ts`). Old helpers were removed from here
// to avoid duplication of responsibility.
