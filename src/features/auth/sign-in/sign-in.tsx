import { Logo } from '@/assets/logo'
// import { cn } from '@/lib/utils'
// import dashboardDark from './assets/dashboard-dark.png'
// import dashboardLight from './assets/dashboard-light.png'

import {
    Card,
} from '@/components/ui/card'
import { UserAuthForm } from './components/user-auth-form'
import { Link } from '@tanstack/react-router'

export function SignIn() {
    return (
        <div className='relative container grid h-svh flex-col items-center justify-center lg:max-w-none lg:px-0'>
            <div className='lg:p-8'>
                <div className='mb-4 flex items-center justify-center'>
                    <Logo className='me-2' />
                    <h1 className='text-xl font-medium'>Shadcn Go Starter</h1>
                </div>
                <Card className='mx-auto flex w-full max-w-sm flex-col justify-center space-y-2 p-5'>
                    <UserAuthForm />
                    {/* Link to sign-up for users without an account */}
                    <p className='text-sm text-center'>
                        Don't have an account?{' '}
                        <Link
                            to='/sign-up'
                            className='font-medium text-primary hover:underline'
                        >
                            Sign up
                        </Link>
                    </p>
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
                </Card>
            </div>
        </div>
    )
}
