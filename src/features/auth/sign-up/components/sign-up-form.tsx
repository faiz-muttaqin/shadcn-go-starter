import { useState } from 'react'
import { z } from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { FaGithub } from "react-icons/fa";
import { FcGoogle } from "react-icons/fc";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { PasswordInput } from '@/components/PasswordInput'
import { useNavigate, Link } from '@tanstack/react-router'
import { toast } from 'sonner'
import { Loader2 } from 'lucide-react'
import { useAuth } from '@/hooks/use-auth'
// Firebase
import { auth as firebaseAuth } from '@/lib/firebase'
import {
  createUserWithEmailAndPassword,
  signInWithPopup,
  GoogleAuthProvider,
  GithubAuthProvider,
  type User as FirebaseUser,
} from 'firebase/auth'

const formSchema = z
  .object({
    email: z.email({
      error: (iss) =>
        iss.input === '' ? 'Please enter your email' : undefined,
    }),
    password: z
      .string()
      .min(1, 'Please enter your password')
      .min(7, 'Password must be at least 7 characters long'),
    confirmPassword: z.string().min(1, 'Please confirm your password'),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords don't match.",
    path: ['confirmPassword'],
  })

export function SignUpForm({
  className,
  ...props
}: React.HTMLAttributes<HTMLFormElement>) {
  const [isLoading, setIsLoading] = useState(false)
  const navigate = useNavigate()
  const { setUser, setAccessToken } = useAuth()

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      password: '',
      confirmPassword: '',
    },
  })

  

  async function finishSignIn(user: FirebaseUser) {
    try {
      const token = await user.getIdToken()
      const userPayload = {
        accountNo: user.uid,
        email: user.email ?? '',
        role: ['user'],
        // placeholder expiry (can be updated by server-side sync)
        exp: 0,
      }

      setUser(userPayload as unknown as import('@/types/auth').User)
      setAccessToken(token)
      navigate({ to: '/', replace: true })
    } catch (err) {
      if (import.meta.env.DEV) console.error('finishSignIn error', err)
      throw err
    }
  }

  async function onSubmit(data: z.infer<typeof formSchema>) {
    setIsLoading(true)

    const p = createUserWithEmailAndPassword(
      firebaseAuth,
      data.email,
      data.password
    )
      .then(async (cred) => {
        await finishSignIn(cred.user)
        return `Welcome, ${cred.user.email ?? data.email}!`
      })
      .catch((err) => {
        throw err
      })
      .finally(() => setIsLoading(false))

    toast.promise(p, {
      loading: 'Creating account...',
      success: (msg) => msg as string,
      error: (err) => (err as Error).message || 'Error creating account',
    })
  }

  async function handleSocialSignup(providerName: 'google' | 'github') {
    setIsLoading(true)
    const provider =
      providerName === 'google' ? new GoogleAuthProvider() : new GithubAuthProvider()

    const p = signInWithPopup(firebaseAuth, provider)
      .then(async (cred) => {
        await finishSignIn(cred.user)
        return `Welcome, ${cred.user.email ?? 'user'}!`
      })
      .catch((err) => {
        throw err
      })
      .finally(() => setIsLoading(false))

    toast.promise(p, {
      loading: `Signing in with ${providerName}...`,
      success: (msg) => msg as string,
      error: (err) => (err as Error).message || `Error signing in with ${providerName}`,
    })
  }

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className={cn('grid gap-4 w-full', className)}
        {...props}
      >
        {/* Social-first CTA */}
        <div className='space-y-2'>
          <div className='text-center'>
            <h2 className='text-lg font-semibold'>Create your account</h2>
            <p className='text-sm text-muted-foreground'>
              Quick start — sign up with Google or GitHub
            </p>
          </div>

          <div className='grid gap-2'>
            <Button
              variant='ghost'
              size='lg'
              className='w-full justify-center space-x-2 border'
              type='button'
              disabled={isLoading}
              onClick={() => handleSocialSignup('google')}
              aria-label='Sign up with Google (recommended)'
            >
              <FcGoogle className='h-5 w-5' />
              <span className='font-medium'>Continue with Google</span>
            </Button>

            <Button
              variant='outline'
              size='lg'
              className='w-full justify-center space-x-2'
              type='button'
              disabled={isLoading}
              onClick={() => handleSocialSignup('github')}
              aria-label='Sign up with GitHub'
            >
              <FaGithub className='h-5 w-5' />
              <span className='font-medium'>Continue with GitHub</span>
            </Button>
          </div>
        </div>

        {/* Divider */}
        <div className='relative my-3'>
          <div className='absolute inset-0 flex items-center'>
            <span className='w-full border-t' />
          </div>
          <div className='relative flex justify-center text-xs uppercase'>
            <span className='bg-background text-muted-foreground px-2'>
              Or sign up with email
            </span>
          </div>
        </div>

        {/* Email / Password fields */}
        <FormField
          control={form.control}
          name='email'
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder='name@example.com' {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name='password'
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <PasswordInput placeholder='********' {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name='confirmPassword'
          render={({ field }) => (
            <FormItem>
              <FormLabel>Confirm Password</FormLabel>
              <FormControl>
                <PasswordInput placeholder='********' {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button className='mt-2 w-full' type='submit' disabled={isLoading}>
          {isLoading ? <Loader2 className='animate-spin' /> : 'Create Account'}
        </Button>

        <p className='text-sm text-center mt-2'>
          Already have an account?{' '}
          <Link to='/sign-in' search={() => ({ redirect: '/' })} className='font-medium text-primary hover:underline'>
            Sign in instead
          </Link>
        </p>
      </form>
    </Form>
  )
}
