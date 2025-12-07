import { createFileRoute } from '@tanstack/react-router'
import { Chats } from '@/features/chats'

export const Route = createFileRoute('/dashboard/_authenticated/chats/')({
  component: Chats,
})
