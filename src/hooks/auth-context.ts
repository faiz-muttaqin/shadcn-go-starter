import { createContext } from 'react'
import type { User } from '@/types/auth'

export interface AuthContextType {
  user: User | null
  isLoading: boolean
  isSynced: boolean
  syncUser: () => Promise<void>
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)
