import { createFileRoute } from '@tanstack/react-router'
import { SettingsAccount } from '@/features/settings/account'

export const Route = createFileRoute('/dashboard/_authenticated/settings/account')({
  component: SettingsAccount,
})
