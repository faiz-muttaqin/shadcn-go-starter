import { createFileRoute, useNavigate, useRouter } from '@tanstack/react-router'
import { AuthenticatedLayout } from '@/components/layout/authenticated-layout'
import { useEffect, useState } from 'react'
import { Button } from '@/components/ui/button'
import { Loader2 } from 'lucide-react'
import { useAuth } from '@/hooks/use-auth'

export const Route = createFileRoute('/dashboard/_authenticated')({
  component: RouteAuthGuard,
})

function RouteAuthGuard() {
  const { isLoading, user, hasLocalToken } = useAuth()
  const router = useRouter()
  const navigate = useNavigate()

  const currentPath = router.state.location.pathname

  // If the user is neither fully authenticated nor has an optimistic token,
  // redirect to sign-in. If there's an optimistic token we allow the
  // AuthenticatedLayout to render while background verification completes.
  useEffect(() => {
    if (user) return
    if (!user && !hasLocalToken) {
      navigate({ to: '/sign-in', search: { redirect: currentPath }, replace: true })
    }
  }, [user, hasLocalToken, navigate, currentPath])

  // While verifying, prefer to show a spinner if no optimistic token is present
  if (isLoading && !hasLocalToken) {
    return (
      <div className='flex h-svh items-center justify-center'>
        <Loader2 className='size-8 animate-spin' />
      </div>
    )
  }

  // Allow access if authenticated or optimistic token exists
  if (user || hasLocalToken) {
    return <AuthenticatedLayout />
  }

  // Fallback: render unauthorized UI (this should rarely be reached because
  // we navigate away in the effect above when not authenticated)
  return <Unauthorized currentPath={currentPath} />
}

const COUNTDOWN = 5 // Countdown second

function Unauthorized({ currentPath }: { currentPath: string }) {
  const navigate = useNavigate()
  const { history } = useRouter()

  const [cancelled, setCancelled] = useState(false)
  const [countdown, setCountdown] = useState(COUNTDOWN)

  // Set and run the countdown conditionally
  useEffect(() => {
    if (cancelled ) return
    const interval = setInterval(() => {
      setCountdown((prev) => (prev > 0 ? prev - 1 : 0))
    }, 1000)
    return () => clearInterval(interval)
  }, [cancelled])

  // Navigate to sign-in page when countdown hits 0
  useEffect(() => {
    if (countdown > 0) return
    // Redirect to sign-in with return URL
    navigate({ 
      to: '/sign-in',
      search: { redirect: currentPath }
    })
  }, [countdown, navigate, currentPath])

  return (
    <div className='h-svh'>
      <div className='m-auto flex h-full w-full flex-col items-center justify-center gap-2'>
        <h1 className='text-[7rem] leading-tight font-bold'>401</h1>
        <span className='font-medium'>Unauthorized Access</span>
        <p className='text-muted-foreground text-center'>
          You must be authenticated
          <sup>
          </sup>
          <br />
          to access this resource.
        </p>
        <div className='mt-6 flex gap-4'>
          <Button variant='outline' onClick={() => history.go(-1)}>
            Go Back
          </Button>
          <Button onClick={() => navigate({ to: '/sign-in', search: { redirect: currentPath } })}>
            Sign in
          </Button>
        </div>
        <div className='mt-4 h-8 text-center'>
          {!cancelled && (
            <>
              <p>
                {countdown > 0
                  ? `Redirecting to Sign In page in ${countdown}s`
                  : `Redirecting...`}
              </p>
              <Button variant='link' onClick={() => setCancelled(true)}>
                Cancel Redirect
              </Button>
            </>
          )}
        </div>
      </div>
    </div>
  )
}
