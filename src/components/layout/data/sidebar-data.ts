import {
  Construction,
  LayoutDashboard,
  Monitor,
  Bug,
  ListTodo,
  FileX,
  HelpCircle,
  Lock,
  Bell,
  Package,
  Palette,
  ServerOff,
  Settings,
  Wrench,
  UserCog,
  UserX,
  Users,
  MessagesSquare,
  ShieldCheck,
  AudioWaveform,
  Command,
  GalleryVerticalEnd,
} from 'lucide-react'
import { type SidebarData } from '../types'

export const sidebarData: SidebarData = {
  user: {
    name: 'satnaing',
    email: 'satnaingdev@gmail.com',
    avatar: '/avatars/shadcn.jpg',
  },
  teams: [
    {
      name: 'Shadcn Admin Go Starter',
      logo: Command,
      plan: 'Vite + ShadcnUI',
    },
    {
      name: 'Acme Inc',
      logo: GalleryVerticalEnd,
      plan: 'Enterprise',
    },
    {
      name: 'Acme Corp.',
      logo: AudioWaveform,
      plan: 'Startup',
    },
  ],
  navGroups: [
    {
      title: 'General',
      items: [
        {
          title: 'Dashboard',
          url: '/dashboard/',
          icon: LayoutDashboard,
        },
        {
          title: 'Tasks',
          url: '/dashboard/tasks',
          icon: ListTodo,
        },
        {
          title: 'Apps',
          url: '/dashboard/apps',
          icon: Package,
        },
        {
          title: 'Chats',
          url: '/dashboard/chats',
          badge: '3',
          icon: MessagesSquare,
        },
        {
          title: 'Users',
          url: '/dashboard/users',
          icon: Users,
        },
      ],
    },
    {
      title: 'Pages',
      items: [
        {
          title: 'Auth',
          icon: ShieldCheck,
          items: [
            {
              title: 'Sign In',
              url: '/dashboard/sign-in',
            },
            {
              title: 'Sign In (2 Col)',
              url: '/dashboard/sign-in-2',
            },
            {
              title: 'Sign Up',
              url: '/dashboard/sign-up',
            },
            {
              title: 'Forgot Password',
              url: '/dashboard/forgot-password',
            },
            {
              title: 'OTP',
              url: '/dashboard/otp',
            },
          ],
        },
        {
          title: 'Errors',
          icon: Bug,
          items: [
            {
              title: 'Unauthorized',
              url: '/dashboard/errors/unauthorized',
              icon: Lock,
            },
            {
              title: 'Forbidden',
              url: '/dashboard/errors/forbidden',
              icon: UserX,
            },
            {
              title: 'Not Found',
              url: '/dashboard/errors/not-found',
              icon: FileX,
            },
            {
              title: 'Internal Server Error',
              url: '/dashboard/errors/internal-server-error',
              icon: ServerOff,
            },
            {
              title: 'Maintenance Error',
              url: '/dashboard/errors/maintenance-error',
              icon: Construction,
            },
          ],
        },
      ],
    },
    {
      title: 'Other',
      items: [
        {
          title: 'Settings',
          icon: Settings,
          items: [
            {
              title: 'Profile',
              url: '/dashboard/settings',
              icon: UserCog,
            },
            {
              title: 'Account',
              url: '/dashboard/settings/account',
              icon: Wrench,
            },
            {
              title: 'Appearance',
              url: '/dashboard/settings/appearance',
              icon: Palette,
            },
            {
              title: 'Notifications',
              url: '/dashboard/settings/notifications',
              icon: Bell,
            },
            {
              title: 'Display',
              url: '/dashboard/settings/display',
              icon: Monitor,
            },
          ],
        },
        {
          title: 'Help Center',
          url: '/dashboard/help-center',
          icon: HelpCircle,
        },
      ],
    },
  ],
}
