import { createFileRoute } from '@tanstack/react-router'
import { SettingsNotifications } from '@/features/settings/notifications'

export const Route = createFileRoute('/dashboard/_authenticated/settings/notifications')({
  component: SettingsNotifications,
})
