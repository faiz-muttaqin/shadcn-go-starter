import { createFileRoute } from '@tanstack/react-router'
import { SignUp } from '@clerk/clerk-react'
import { Skeleton } from '@/components/ui/skeleton'
import { clerkAppearance, clerkDarkAppearance } from '@/lib/clerk-appearance'
import { useTheme } from '@/context/theme-provider'
import { Logo } from '@/assets/logo'

export const Route = createFileRoute('/(auth)/sign-up')({
  component: SignUpPage,
})

function SignUpPage() {
  const { resolvedTheme } = useTheme()
  const appearance = resolvedTheme === 'dark' ? clerkDarkAppearance : clerkAppearance

  return (
    <div className='container grid h-svh max-w-none items-center justify-center'>
      <div className='mx-auto flex w-full flex-col justify-center space-y-2 py-8 sm:w-[480px] sm:p-8'>
        <div className='mb-4 flex items-center justify-center'>
          <Logo className='me-2' />
          <h1 className='text-xl font-medium'>Shadcn Admin Go Starter</h1>
        </div>
        <div className='flex flex-col space-y-2 text-start px-3'>
          <h2 className='text-lg font-semibold tracking-tight'>Create an account</h2>
          <p className='text-muted-foreground text-sm'>
            Enter your email and password below to create your account
          </p>
        </div>
          
          <SignUp
            fallback={<Skeleton className='h-[30rem] w-[25rem]' />}
            appearance={appearance}
          />
          <p className='text-muted-foreground px-8 text-center text-sm'>
            By clicking sign up, you agree to our{' '}
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

  )
}

