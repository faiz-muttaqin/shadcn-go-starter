import { type QueryClient } from '@tanstack/react-query'
import { createRootRouteWithContext, Outlet } from '@tanstack/react-router'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'
import { Toaster } from '@/components/ui/sonner'
import { NavigationProgress } from '@/components/NavigationProgress'
import { GeneralError } from '@/features/errors/general-error'
import { NotFoundError } from '@/features/errors/not-found-error'
import { AuthProvider } from '@/hooks/use-auth'
import { initializeApiClient } from '@/lib/api-client'
import { auth as firebaseAuth } from '@/lib/firebase'
import { useEffect } from 'react'

// Component to initialize API client with Firebase token getter
function ApiClientInitializer() {
  useEffect(() => {
    // Provide a function that returns current Firebase ID token or null
    const getFirebaseToken = async () => {
      try {
        const user = firebaseAuth.currentUser
        if (user) return await user.getIdToken()
        return null
      } catch (err) {
        if (import.meta.env.DEV) {
          console.error('Failed to get Firebase token', err)
        }
        return null
      }
    }

    initializeApiClient(getFirebaseToken)
  }, [])

  return null
}

// Root component wrapper
function RootComponent() {
  return (
    <>
      <NavigationProgress />
      <ApiClientInitializer />
      <AuthProvider>
        <Outlet />
      </AuthProvider>
      <Toaster duration={5000} />
      {import.meta.env.MODE === 'development' && (
        <>
          <ReactQueryDevtools buttonPosition='bottom-left' />
          <TanStackRouterDevtools position='bottom-right' />
        </>
      )}
    </>
  )
}

export const Route = createRootRouteWithContext<{
  queryClient: QueryClient
}>()({
  component: RootComponent,
  notFoundComponent: NotFoundError,
  errorComponent: GeneralError,
})
