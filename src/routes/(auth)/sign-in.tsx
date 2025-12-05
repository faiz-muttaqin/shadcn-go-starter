import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { SignIn, useAuth } from '@clerk/clerk-react'
import { Skeleton } from '@/components/ui/skeleton'
import { clerkAppearance, clerkDarkAppearance } from '@/lib/clerk-appearance'
import { useTheme } from '@/context/theme-provider'
import { Logo } from '@/assets/logo'
import { useEffect } from 'react'

export const Route = createFileRoute('/(auth)/sign-in')({
  component: SignInPage,
  validateSearch: (search: Record<string, unknown>) => {
    return {
      redirect: (search.redirect as string) || '/',
    }
  },
})

function SignInPage() {
  const { resolvedTheme } = useTheme()
  const { isSignedIn } = useAuth()
  const navigate = useNavigate()
  const { redirect } = Route.useSearch()
  const appearance = resolvedTheme === 'dark' ? clerkDarkAppearance : clerkAppearance
  
  // Redirect after successful sign in
  useEffect(() => {
    if (isSignedIn) {
      navigate({ to: redirect })
    }
  }, [isSignedIn, navigate, redirect])
  
  return (
    <div className="space-y-2 ">
      {/* Custom Header */}
      <div className='lg:p-8'>
        <div className='mx-auto flex w-full flex-col justify-center space-y-2 py-8 sm:w-[480px] sm:p-8'>
          <div className='mb-4 flex items-center justify-center'>
            <Logo className='me-2' />
            <h1 className='text-xl font-medium'>Shadcn Admin Go Starter</h1>
          </div>
        </div>
        <div className='mx-auto flex w-full max-w-sm flex-col justify-center space-y-2'>
          <div className='flex flex-col space-y-2 text-start px-3'>
            <h2 className='text-lg font-semibold tracking-tight'>Sign in</h2>
            <p className='text-muted-foreground text-sm'>
              Enter your email and password below OR use Gmail <br />
              to log into your account
            </p>
          </div>
          {/* Clerk Form */}
          <SignIn
            initialValues={{ emailAddress: 'your_mail+shadcn_admin@gmail.com' }}
            fallback={<Skeleton className='h-[30rem] w-[25rem]' />}
            appearance={appearance}
          />
          <p className='text-muted-foreground px-8 text-center text-sm'>
            By clicking sign in, you agree to our{' '}
            <a
              href='/terms'
              className='hover:text-primary underline underline-offset-4'
            >
              Terms of Service
            </a>{' '}
            and{' '}
            <a
              href='/privacy'
              className='hover:text-primary underline underline-offset-4'
            >
              Privacy Policy
            </a>
            .
          </p>
        </div>
      </div>
    </div>
  )
}

