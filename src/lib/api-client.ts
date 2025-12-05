import axios, { type AxiosError, type InternalAxiosRequestConfig } from 'axios'
import { toast } from 'sonner'

// Create axios instance with default config
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: 30000, // 30 seconds
  headers: {
    'Content-Type': 'application/json',
  },
})

// Store for getting Clerk token (will be set after Clerk loads)
let getClerkToken: (() => Promise<string | null>) | null = null

/**
 * Initialize API client with Clerk auth
 * Call this after ClerkProvider is mounted
 */
export function initializeApiClient(getToken: () => Promise<string | null>) {
  getClerkToken = getToken
}

// Request interceptor - Add Clerk token to all requests
apiClient.interceptors.request.use(
  async (config: InternalAxiosRequestConfig) => {
    try {
      // Get Clerk token if available
      if (getClerkToken) {
        const token = await getClerkToken()
        if (token) {
          config.headers.Authorization = `Bearer ${token}`
        }
      }
    } catch (error) {
      // Silent fail - continue without token
      if (import.meta.env.DEV) {
        // eslint-disable-next-line no-console
        console.error('Failed to get auth token:', error)
      }
    }

    return config
  },
  (error: AxiosError) => {
    return Promise.reject(error)
  }
)

// Response interceptor - Handle errors globally
apiClient.interceptors.response.use(
  (response) => {
    return response
  },
  async (error: AxiosError<{ error?: string; message?: string }>) => {
    const status = error.response?.status
    const errorMessage =
      error.response?.data?.error ||
      error.response?.data?.message ||
      'An error occurred'

    // Handle specific error codes
    switch (status) {
      case 400:
        toast.error('Bad Request', {
          description: errorMessage,
        })
        break

      case 401:
        toast.error('Session Expired', {
          description: 'Please sign in again to continue',
        })
        // Redirect to sign-in will be handled by QueryCache in main.tsx
        break

      case 403:
        toast.error('Access Denied', {
          description: "You don't have permission to access this resource",
        })
        break

      case 404:
        toast.error('Not Found', {
          description: 'The requested resource was not found',
        })
        break

      case 422:
        toast.error('Validation Error', {
          description: errorMessage,
        })
        break

      case 429:
        toast.error('Too Many Requests', {
          description: 'Please slow down and try again later',
        })
        break

      case 500:
        toast.error('Server Error', {
          description: 'Something went wrong on our end',
        })
        break

      case 503:
        toast.error('Service Unavailable', {
          description: 'The service is temporarily unavailable',
        })
        break

      default:
        if (error.code === 'ECONNABORTED') {
          toast.error('Request Timeout', {
            description: 'The request took too long to complete',
          })
        } else if (error.code === 'ERR_NETWORK') {
          toast.error('Network Error', {
            description: 'Please check your internet connection',
          })
        } else if (status && status >= 500) {
          toast.error('Server Error', {
            description: errorMessage,
          })
        }
    }

    return Promise.reject(error)
  }
)

export default apiClient
