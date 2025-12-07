import { type ReactNode, useEffect, useState } from 'react'
import { onAuthStateChanged } from 'firebase/auth'
import { auth as firebaseAuth } from '@/lib/firebase'
import { syncUserData } from '@/lib/auth-service'
import { AuthContext } from './auth-context'
import type { User } from '@/types/auth'

export function AuthProvider({ children }: { children: ReactNode }) {
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
        console.error('Failed to sync user:', error)
      }
      setIsSynced(false)
    } finally {
      setIsLoading(false)
    }
  }

  useEffect(() => {
    const unsub = onAuthStateChanged(firebaseAuth, async (fbUser) => {
      if (fbUser) {
        await syncUser()
      } else {
        setUser(null)
        setIsSynced(false)
        setIsLoading(false)
      }
    })

    return () => unsub()
  }, [])

  return (
    <AuthContext.Provider value={{ user, isLoading, isSynced, syncUser }}>
      {children}
    </AuthContext.Provider>
  )
}
