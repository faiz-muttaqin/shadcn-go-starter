import { createFileRoute } from '@tanstack/react-router'
import { SettingsProfile } from '@/features/settings/profile'

export const Route = createFileRoute('/dashboard/_authenticated/settings/')({
  component: SettingsProfile,
})
