import {
  createFileRoute, Link,
  useNavigate,
  useRouter,
} from '@tanstack/react-router'
import { AuthenticatedLayout } from '@/components/layout/authenticated-layout'
import { useEffect, useState } from 'react'
import { Button } from '@/components/ui/button'
import { LearnMore } from '@/components/LearnMore'
import { Loader2 } from 'lucide-react'
import { useAuth } from '@/hooks/use-auth'

export const Route = createFileRoute('/dashboard/_authenticated')({
  component: RouteAuthGuard,
})

function RouteAuthGuard() {
  const { isLoading, user } = useAuth()
  const router = useRouter()

  if (isLoading) {
    return (
      <div className='flex h-svh items-center justify-center'>
        <Loader2 className='size-8 animate-spin' />
      </div>
    )
  }

  if (!user) {
    return <Unauthorized currentPath={router.state.location.pathname} />
  }

  return <AuthenticatedLayout />
}

const COUNTDOWN = 5 // Countdown second

function Unauthorized({ currentPath }: { currentPath: string }) {
  const navigate = useNavigate()
  const { history } = useRouter()

  const [opened, setOpened] = useState(true)
  const [cancelled, setCancelled] = useState(false)
  const [countdown, setCountdown] = useState(COUNTDOWN)

  // Set and run the countdown conditionally
  useEffect(() => {
    if (cancelled || opened) return
    const interval = setInterval(() => {
      setCountdown((prev) => (prev > 0 ? prev - 1 : 0))
    }, 1000)
    return () => clearInterval(interval)
  }, [cancelled, opened])

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
          You must be authenticated via Clerk{' '}
          <sup>
            <LearnMore open={opened} onOpenChange={setOpened}>
              <p>
                This is the same as{' '}
                <Link
                  to='/dashboard/users'
                  className='text-blue-500 underline decoration-dashed underline-offset-2'
                >
                  '/dashboard/users'
                </Link>
                .{' '}
              </p>
              <p>You must first sign in using Clerk to access this route. </p>

              <p className='mt-4'>
                After signing in, you'll be able to sign out or delete your
                account via the User Profile dropdown on this page.
              </p>
            </LearnMore>
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
          {!cancelled && !opened && (
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
