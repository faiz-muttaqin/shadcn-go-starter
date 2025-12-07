import { useState } from 'react'
import { z } from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Link, useNavigate } from '@tanstack/react-router'
import { Loader2, LogIn } from 'lucide-react'
import { toast } from 'sonner'
import { FaGithub } from "react-icons/fa";
import { FcGoogle } from "react-icons/fc";
import { useAuth } from '@/hooks/use-auth'
import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
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
// Firebase
import { auth as firebaseAuth } from '@/lib/firebase'
import {
  signInWithEmailAndPassword,
  signInWithPopup,
  GoogleAuthProvider,
  GithubAuthProvider,
} from 'firebase/auth'

const formSchema = z.object({
  email: z.email({
    error: (iss) => (iss.input === '' ? 'Please enter your email' : undefined),
  }),
  password: z
    .string()
    .min(1, 'Please enter your password')
    .min(7, 'Password must be at least 7 characters long'),
})

interface UserAuthFormProps extends React.HTMLAttributes<HTMLFormElement> {
  redirectTo?: string
}

export function UserAuthForm({
  className,
  redirectTo,
  ...props
}: UserAuthFormProps) {
  const [isLoading, setIsLoading] = useState(false)
  const navigate = useNavigate()
  const { setUser, setAccessToken } = useAuth()

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  async function onSubmit(data: z.infer<typeof formSchema>) {
    setIsLoading(true)

    const p = signInWithEmailAndPassword(
      firebaseAuth,
      data.email,
      data.password
    )
      .then(async (cred) => {
        const token = await cred.user.getIdToken()

        const userPayload = {
          accountNo: cred.user.uid,
          email: cred.user.email ?? data.email,
          role: ['user'],
          exp: Date.now() + 24 * 60 * 60 * 1000,
        }

  // store a minimal snapshot immediately; provider will reconcile with backend
  setUser(userPayload as unknown as import('@/types/auth').User)
        setAccessToken(token)

        const targetPath = redirectTo || '/dashboard'
        navigate({ to: targetPath, replace: true })

        return `Welcome back, ${cred.user.email ?? data.email}!`
      })
      .catch((err) => {
        // rethrow to let toast.promise show error
        throw err
      })
      .finally(() => setIsLoading(false))

    toast.promise(p, {
      loading: 'Signing in...',
      success: (msg) => msg as string,
      error: (err) => (err as Error).message || 'Error signing in',
    })
  }

  async function handleOAuthLogin(providerName: 'google' | 'github') {
    setIsLoading(true)

    const provider =
      providerName === 'google' ? new GoogleAuthProvider() : new GithubAuthProvider()

    const p = signInWithPopup(firebaseAuth, provider)
      .then(async (cred) => {
        const token = await cred.user.getIdToken()
        const userPayload = {
          accountNo: cred.user.uid,
          // Ensure email is a string to satisfy AuthUser type
          email: cred.user.email ?? '',
          role: ['user'],
          exp: Date.now() + 24 * 60 * 60 * 1000,
        }

        setUser(userPayload as unknown as import('@/types/auth').User)
        setAccessToken(token)

        const targetPath = redirectTo || '/'
        navigate({ to: targetPath, replace: true })

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
        {/* Primary call-to-action: social sign-ins (Google emphasized) */}
        <div className="space-y-2">
          <div className="text-center">
            <h2 className="text-lg font-semibold">Sign in</h2>
            <p className="text-sm text-muted-foreground">
              Quick sign in — recommended for the best experience
            </p>
          </div>

          <div className="grid gap-2">
            <Button
              variant="ghost"
              size="lg"
              className="w-full justify-center space-x-2 border"
              type="button"
              disabled={isLoading}
              onClick={() => handleOAuthLogin('google')}
              aria-label="Sign in with Google (recommended)"
            >
              <FcGoogle className="h-5 w-5" />
              <span className="font-medium">Continue with Google</span>
            </Button>

            <Button
              variant="outline"
              size="lg"
              className="w-full justify-center space-x-2"
              type="button"
              disabled={isLoading}
              onClick={() => handleOAuthLogin('github')}
              aria-label="Sign in with GitHub"
            >
              <FaGithub className="h-5 w-5" />
              <span className="font-medium">Continue with GitHub</span>
            </Button>
          </div>
        </div>

        {/* Divider */}
        <div className="relative my-3">
          <div className="absolute inset-0 flex items-center">
            <span className="w-full border-t" />
          </div>
          <div className="relative flex justify-center text-xs uppercase">
            <span className="bg-background text-muted-foreground px-2">
              Or use your email
            </span>
          </div>
        </div>

        {/* Email / Password form (secondary) */}
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="name@example.com" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem className="relative">
              <FormLabel>Password</FormLabel>
              <FormControl>
                <PasswordInput placeholder="********" {...field} />
              </FormControl>
              <FormMessage />
              <Link
                to="/forgot-password"
                className="text-muted-foreground absolute end-0 -top-0.5 text-sm font-medium hover:opacity-75"
              >
                Forgot password?
              </Link>
            </FormItem>
          )}
        />

        <Button className="mt-2 w-full" type="submit" disabled={isLoading}>
          {isLoading ? <Loader2 className="animate-spin" /> : <LogIn />}
          <span className="ml-2">Sign in with email</span>
        </Button>
      </form>
    </Form>
  )
}
