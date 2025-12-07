import { createFileRoute } from '@tanstack/react-router'
import { ComingSoon } from '@/components/ComingSoon'

export const Route = createFileRoute('/dashboard/_authenticated/help-center/')({
  component: ComingSoon,
})
