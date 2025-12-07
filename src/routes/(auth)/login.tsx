import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { SignIn } from '@/features/auth/sign-in/sign-in'
import {auth} from '@/lib/firebase'
import { useEffect } from 'react'

export const Route = createFileRoute('/(auth)/login')({
  component: SignInPage,
  validateSearch: (search: Record<string, unknown>) => {
    return {
      redirect: (search.redirect as string) || '/dashboard',
    }
  },
})

function SignInPage() {
  // const { isSignedIn } = useAuth()
  const navigate = useNavigate()
  const { redirect } = Route.useSearch()
  // Redirect after successful sign in
  useEffect(() => {
    if (auth.currentUser) {
      // navigate({ to: redirect })
    }
  }, [navigate, redirect])
  
  return <SignIn />;
}

