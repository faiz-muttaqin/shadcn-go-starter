import { type QueryClient } from '@tanstack/react-query'
import { createRootRouteWithContext, Outlet } from '@tanstack/react-router'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'
import { Toaster } from '@/components/ui/sonner'
import { NavigationProgress } from '@/components/NavigationProgress'
import { GeneralError } from '@/features/errors/general-error'
import { NotFoundError } from '@/features/errors/not-found-error'
import { ClerkProvider, useAuth as useClerkAuth } from '@clerk/clerk-react'
import { AuthProvider } from '@/hooks/use-auth'
import { initializeApiClient } from '@/lib/api-client'
import { useEffect } from 'react'

// Import your Publishable Key
const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY

// Component to initialize API client with Clerk token
function ApiClientInitializer() {
  const { getToken } = useClerkAuth()

  useEffect(() => {
    // Initialize API client with Clerk token getter
    initializeApiClient(getToken)
  }, [getToken])

  return null
}

// Root component wrapper
function RootComponent() {
  if (!PUBLISHABLE_KEY) {
    throw new Error('Missing Clerk Publishable Key')
  }

  return (
    <>
      <NavigationProgress />
      <ClerkProvider
        publishableKey={PUBLISHABLE_KEY}
        afterSignOutUrl='/sign-in'
        signInUrl='/sign-in'
        signUpUrl='/sign-up'
        signInFallbackRedirectUrl='/'
        signUpFallbackRedirectUrl='/'
      >
        <ApiClientInitializer />
        <AuthProvider>
          <Outlet />
        </AuthProvider>
      </ClerkProvider>
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
