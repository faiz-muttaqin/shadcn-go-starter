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
  const [hasLocalToken, setHasLocalToken] = useState(false)

  /**
   * Sync user from backend. When called from background verification you can
   * suppress the loading flag so the UI doesn't flash a spinner.
   */
  const syncUser = async (opts?: { suppressLoading?: boolean }) => {
    const suppress = Boolean(opts?.suppressLoading)
    try {
      if (!suppress) setIsLoading(true)
      const userData = await syncUserData()

      if (userData) {
        setUser(userData)
        setIsSynced(true)
      }
      return userData ?? null
    } catch (error) {
      if (import.meta.env.DEV) {
        console.error('Failed to sync user:', error)
      }
      setIsSynced(false)
      return null
    } finally {
      if (!suppress) setIsLoading(false)
    }
  }

  useEffect(() => {
    // Check localStorage for optimistic token and optional cached user
    let stored: string | null = null
    let storedUserJson: string | null = null
    try {
      if (typeof window !== 'undefined') {
        stored = localStorage.getItem('firebase_id_token')
        storedUserJson = localStorage.getItem('firebase_user')
      }
    } catch {
      // ignore localStorage errors
    }

    const hadLocalToken = Boolean(stored)
    setHasLocalToken(hadLocalToken)
    // If we have an optimistic token, avoid showing a loading spinner and
    // hydrate the user from cache so UI can render immediately. We'll
    // still verify with Firebase in the background and reconcile.
    if (hadLocalToken) {
      setIsLoading(false)
      if (storedUserJson) {
        try {
          const parsed = JSON.parse(storedUserJson) as User
          setUser(parsed)
        } catch {
          // ignore parse errors
        }
      }
    } else {
      // No optimistic token — we are in an actual loading state until
      // Firebase responds.
      setIsLoading(true)
    }

    // Subscribe to Firebase auth state
    const unsub = onAuthStateChanged(firebaseAuth, async (fbUser) => {
      try {
        // Only flip loading to true when we didn't optimistically assume a
        // token on startup. If hadLocalToken is true we keep isLoading false
        // to avoid UI flashing.
        if (!hadLocalToken) setIsLoading(true)

        if (fbUser) {
          // Get fresh token and store it
          try {
            const token = await fbUser.getIdToken()
            if (token) {
              try {
                localStorage.setItem('firebase_id_token', token)
              } catch {
                // ignore storage errors
              }
              setHasLocalToken(true)
            }
          } catch {
            // ignore token set failure
          }

          // Sync user from backend but suppress loading when we bootstrapped
          // from a cached token so UI remains responsive. Persist the
          // returned snapshot to localStorage so next reload can hydrate.
          const synced = await syncUser({ suppressLoading: hadLocalToken })
          if (synced) {
            try {
              localStorage.setItem('firebase_user', JSON.stringify(synced))
            } catch {
              // ignore
            }
          }
        } else {
          // No firebase user: if we had an optimistic token, clear it
          if (hadLocalToken) {
            try {
              localStorage.removeItem('firebase_id_token')
              localStorage.removeItem('firebase_user')
            } catch {
              // ignore
            }
            setHasLocalToken(false)
          }
          setUser(null)
          setIsSynced(false)
        }
      } finally {
        if (!hadLocalToken) setIsLoading(false)
      }
    })

    return () => unsub()
  }, [])

  const signOut = async () => {
    try {
      await firebaseAuth.signOut()
    } catch {
      // ignore
    }
    try {
      localStorage.removeItem('firebase_id_token')
    } catch {
      // ignore
    }
    setUser(null)
    setHasLocalToken(false)
    setIsSynced(false)
  }

  const setAccessToken = (token: string) => {
    try {
      localStorage.setItem('firebase_id_token', token)
    } catch {
      // ignore
    }
    setHasLocalToken(Boolean(token))
  }

  const setUserInContext = (u: User | null) => {
    setUser(u)
    try {
      if (u) {
        localStorage.setItem('firebase_user', JSON.stringify(u))
      } else {
        localStorage.removeItem('firebase_user')
      }
    } catch {
      // ignore
    }
  }

  const clearAuth = () => {
    try {
      localStorage.removeItem('firebase_id_token')
      localStorage.removeItem('firebase_user')
    } catch {
      // ignore
    }
    setUser(null)
    setHasLocalToken(false)
    setIsSynced(false)
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        isLoading,
        isSynced,
        syncUser,
        hasLocalToken,
        signOut,
        setUser: setUserInContext,
        setAccessToken,
        clearAuth,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}
