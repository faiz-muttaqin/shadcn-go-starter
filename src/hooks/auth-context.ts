import { createContext } from 'react'
import type { User } from '@/types/auth'

export interface AuthContextType {
  user: User | null
  isLoading: boolean
  isSynced: boolean
  syncUser: (opts?: { suppressLoading?: boolean }) => Promise<import('@/types/auth').User | null>
  /** True when there is a token stored in localStorage (optimistic auth) */
  hasLocalToken: boolean
  /** Sign the user out and clear local storage token */
  signOut: () => Promise<void>
  /** Set the current user (used after login) */
  setUser: (u: User | null) => void
  /** Set the current access token (used after login) */
  setAccessToken: (token: string) => void
  /** Clear auth state client-side (remove token, user) */
  clearAuth: () => void
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)
