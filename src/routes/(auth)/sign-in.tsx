import { createFileRoute } from '@tanstack/react-router'
import { SignIn } from '@/features/auth/sign-in/sign-in'

export const Route = createFileRoute('/(auth)/sign-in')({
  component: SignInPage,
  validateSearch: (search: Record<string, unknown>) => {
    return {
      redirect: (search.redirect as string) || '/',
    }
  },
})

function SignInPage() {

  return <SignIn />;
}

