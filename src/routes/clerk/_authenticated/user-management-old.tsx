import { toast } from 'sonner'
import { useEffect, useState } from 'react'

import {
  createFileRoute,
  Link,
  useNavigate,
  useRouter,
} from '@tanstack/react-router'
import { SignedIn, useAuth, UserButton } from '@clerk/clerk-react'
import { CheckCircle2, XCircle, ExternalLink, Loader2 } from 'lucide-react'
import { ClerkLogo } from '@/assets/clerk-logo'
import { Button } from '@/components/ui/button'
import { Header } from '@/components/layout/header'
import { Main } from '@/components/layout/main'
import { LearnMore } from '@/components/LearnMore'
import { Search } from '@/components/Search'
import { ThemeSwitch } from '@/components/ThemeSwitch'
import { UsersDialogs } from '@/features/users/components/users-dialogs'
import { UsersPrimaryButtons } from '@/features/users/components/users-primary-buttons'
import { UsersProvider } from '@/features/users/components/users-provider'
import { UsersTable } from '@/features/users/components/users-table'
import { users } from '@/features/users/data/users'
import apiClient from '@/lib/api-client'


export const Route = createFileRoute('/clerk/_authenticated/user-management-old')({
  component: UserManagement,
})

function UserManagement() {
  const search = Route.useSearch()
  const navigate = Route.useNavigate()

  const [opened, setOpened] = useState(true)
  const { isLoaded, isSignedIn } = useAuth()


  const { getToken, userId } = useAuth()
  const [isTestingAuth, setIsTestingAuth] = useState(false)

  const testAuthStatus = async () => {
    setIsTestingAuth(true)

    try {
      // 1. Check if user is signed in
      if (!isSignedIn) {
        toast.error('Not Authenticated', {
          description: 'You are not signed in. Please sign in first.',
        })
        setIsTestingAuth(false)
        return
      }

      // 2. Get Clerk token
      const token = await getToken()

      if (!token) {
        toast.error('No Token Found', {
          description: 'Could not retrieve authentication token.',
        })
        setIsTestingAuth(false)
        return
      }

      // 3. Test API call to backend (replace with your actual endpoint)
      const response = await apiClient.get('/auth/verify', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })

      // 4. Show success message with token info
      toast.success('Authentication Successful!', {
        description: (
          <div className="space-y-1 text-xs">
            <p>✅ Signed In: Yes</p>
            <p>✅ User ID: {userId}</p>
            <p>✅ Token: {token.substring(0, 20)}...</p>
            <p>✅ API Response: {response.status}</p>
          </div>
        ),
        duration: 5000,
      })
    } catch (error: unknown) {
      // Handle API errors
      const apiError = error as { response?: { status?: number; data?: { error?: string } }; message?: string }
      const status = apiError.response?.status
      const message = apiError.response?.data?.error || apiError.message

      if (status === 401) {
        toast.error('Token Invalid or Expired', {
          description: 'Your authentication token is not valid.',
        })
      } else if (status === 404) {
        toast.warning('Endpoint Not Found', {
          description: 'Backend auth verification endpoint not implemented yet.',
        })
      } else {
        toast.error('API Request Failed', {
          description: message,
        })
      }
    } finally {
      setIsTestingAuth(false)
    }
  }

  const showTokenInfo = async () => {
    if (!isSignedIn) {
      toast.info('Not Signed In', {
        description: 'Please sign in to view token information.',
      })
      return
    }

    const token = await getToken()

    // Show token in console for debugging
    // eslint-disable-next-line no-console
    console.log('🔐 Clerk Authentication Info:', {
      isSignedIn,
      userId,
      token,
      tokenLength: token?.length,
      tokenPreview: token?.substring(0, 50) + '...',
    })

    toast.success('Token Info', {
      description: (
        <div className="space-y-1 text-xs font-mono">
          <p>User ID: {userId}</p>
          <p>Token Length: {token?.length} chars</p>
          <p>Check console for full token</p>
        </div>
      ),
    })
  }

  if (!isLoaded) {
    return (
      <div className='flex h-svh items-center justify-center'>
        <Loader2 className='size-8 animate-spin' />
      </div>
    )
  }

  if (!isSignedIn) {
    return <Unauthorized />
  }

  return (
    <>
      <SignedIn>
        <UsersProvider>
          <Header fixed>
            <Search />
            <div className='ms-auto flex items-center space-x-4'>
              <ThemeSwitch />
              <UserButton />
            </div>
          </Header>

          <Main>
            <div className='mb-2 flex flex-wrap items-center justify-between space-y-2'>
              <div>
                <h2 className='text-2xl font-bold tracking-tight'>User List</h2>
                <Button
                  variant='outline'
                  size='sm'
                  onClick={showTokenInfo}
                  disabled={!isLoaded || !isSignedIn}
                >
                  {isSignedIn ? (
                    <>
                      <CheckCircle2 className='mr-2 h-4 w-4 text-green-500' />
                      Show Token
                    </>
                  ) : (
                    <>
                      <XCircle className='mr-2 h-4 w-4 text-red-500' />
                      Not Signed In
                    </>
                  )}
                </Button>
                <Button
                  variant='default'
                  size='sm'
                  onClick={testAuthStatus}
                  disabled={isTestingAuth || !isLoaded || !isSignedIn}
                >
                  {isTestingAuth ? (
                    <>
                      <Loader2 className='mr-2 h-4 w-4 animate-spin' />
                      Testing...
                    </>
                  ) : (
                    'Test Auth'
                  )}
                </Button>
                <div className='flex gap-1'>
                  <p className='text-muted-foreground'>
                    Manage your users and their roles here.
                  </p>
                  <LearnMore
                    open={opened}
                    onOpenChange={setOpened}
                    contentProps={{ side: 'right' }}
                  >
                    <p>
                      This is the same as{' '}
                      <Link
                        to='/users'
                        className='text-blue-500 underline decoration-dashed underline-offset-2'
                      >
                        '/users'
                      </Link>
                    </p>

                    <p className='mt-4'>
                      You can sign out or manage/delete your account via the
                      User Profile menu in the top-right corner of the page.
                      <ExternalLink className='inline-block size-4' />
                    </p>
                  </LearnMore>
                </div>
              </div>
              <UsersPrimaryButtons />
            </div>
            <div className='-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-y-0 lg:space-x-12'>
              <UsersTable data={users} navigate={navigate} search={search} />
            </div>
          </Main>

          <UsersDialogs />
        </UsersProvider>
      </SignedIn>
    </>
  )
}

const COUNTDOWN = 5 // Countdown second

function Unauthorized() {
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
    navigate({ to: '/clerk/sign-in' })
  }, [countdown, navigate])

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
                  to='/users'
                  className='text-blue-500 underline decoration-dashed underline-offset-2'
                >
                  '/users'
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
          <Button onClick={() => navigate({ to: '/clerk/sign-in' })}>
            <ClerkLogo className='invert' /> Sign in
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
