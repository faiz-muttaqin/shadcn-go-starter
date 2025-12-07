import { createFileRoute } from '@tanstack/react-router'
import { SettingsDisplay } from '@/features/settings/display'

export const Route = createFileRoute('/dashboard/_authenticated/settings/display')({
  component: SettingsDisplay,
})
