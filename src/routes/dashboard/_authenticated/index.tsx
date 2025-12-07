import { createFileRoute } from '@tanstack/react-router'
import { Dashboard } from '@/features/dashboard'

export const Route = createFileRoute('/dashboard/_authenticated/')({
  component: Dashboard,
})
