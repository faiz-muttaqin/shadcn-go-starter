import { createFileRoute, Outlet, useNavigate } from '@tanstack/react-router'
import { useEffect } from 'react'
import { useAuth } from '@/hooks/use-auth'
import { z } from 'zod'

export const Route = createFileRoute('/(auth)')({
  component: AuthGuardLayout,
  validateSearch: z.object({
    redirect: z
      .string()
      .optional()
      .refine(v => !v || v.startsWith('/'))
      .default('/dashboard'),
  }),
})

function AuthGuardLayout() {
  const navigate = useNavigate()
  const { redirect } = Route.useSearch()
  const { user, hasLocalToken, isLoading } = useAuth()

  // Redirect to dashboard if we already have an authenticated session or an
  // optimistic token in localStorage. We also guard by isLoading so we don't
  // redirect while verification is in progress for non-optimistic flows.
  useEffect(() => {
    const target = (redirect as string) || '/dashboard'
    if (user) {
      navigate({ to: target, replace: true })
      return
    }

    // If there's an optimistic token stored, treat the user as logged in and
    // navigate immediately (provider hydrates cached user in background).
    if (hasLocalToken && !isLoading) {
      navigate({ to: target, replace: true })
    }
  }, [user, hasLocalToken, isLoading, navigate, redirect])

  return <Outlet />
}