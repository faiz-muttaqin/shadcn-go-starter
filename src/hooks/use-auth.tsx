import { createContext, useContext, useEffect, useState, type ReactNode } from 'react'
import { useAuth as useClerkAuth, useUser } from '@clerk/clerk-react'
import { syncUserData } from '@/lib/auth-service'
import type { User } from '@/types/auth'

interface AuthContextType {
  user: User | null
  isLoading: boolean
  isSynced: boolean
  syncUser: () => Promise<void>
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
  const { isSignedIn, isLoaded } = useClerkAuth()
  const { user: clerkUser } = useUser()
  const [user, setUser] = useState<User | null>(null)
  const [isSynced, setIsSynced] = useState(false)
  const [isLoading, setIsLoading] = useState(true)

  const syncUser = async () => {
    try {
      setIsLoading(true)
      const userData = await syncUserData()

      if (userData) {
        setUser(userData)
        setIsSynced(true)
      }
    } catch (error) {
      if (import.meta.env.DEV) {
        // eslint-disable-next-line no-console
        console.error('Failed to sync user:', error)
      }
      setIsSynced(false)
    } finally {
      setIsLoading(false)
    }
  }

  // Auto sync when user signs in
  useEffect(() => {
    if (isLoaded && isSignedIn && clerkUser && !isSynced) {
      syncUser()
    } else if (isLoaded && !isSignedIn) {
      setUser(null)
      setIsSynced(false)
      setIsLoading(false)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isLoaded, isSignedIn, clerkUser])

  return (
    <AuthContext.Provider value={{ user, isLoading, isSynced, syncUser }}>
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth() {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error('useAuth must be used within AuthProvider')
  }
  return context
}
