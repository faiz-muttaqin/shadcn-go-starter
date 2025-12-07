import z from 'zod'
// import { Users } from '@/features/users'
import { roles } from '@/features/users/data/data'
import {
  createFileRoute,
} from '@tanstack/react-router'
import { useAuth } from '@/hooks/use-auth'
import { Loader2 } from 'lucide-react'
import { Main } from '@/components/layout/main'
import { UsersDialogs } from '@/features/users/components/users-dialogs'
import { UsersPrimaryButtons } from '@/features/users/components/users-primary-buttons'
import { UsersProvider } from '@/features/users/components/users-provider'
import { UsersTable } from '@/features/users/components/users-table'
import { users } from '@/features/users/data/users'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { ConfigurableTable } from '@/components/ConfigurableTable'

const usersSearchSchema = z.object({
  page: z.number().optional().catch(1),
  pageSize: z.number().optional().catch(10),
  // Facet filters
  status: z
    .array(
      z.union([
        z.literal('active'),
        z.literal('inactive'),
        z.literal('invited'),
        z.literal('suspended'),
      ])
    )
    .optional()
    .catch([]),
  role: z
    .array(z.enum(roles.map((r) => r.value as (typeof roles)[number]['value'])))
    .optional()
    .catch([]),
  // Per-column text filter (example for username)
  username: z.string().optional().catch(''),
})

export const Route = createFileRoute('/dashboard/_authenticated/users/')({
  validateSearch: usersSearchSchema,
  component: UserManagement,
})

function UserManagement() {
  const search = Route.useSearch()
  const navigate = Route.useNavigate()
  // console.log('Users Route users : ', users)
  const { isLoading } = useAuth()

  if (isLoading) {
    return (
      <div className='flex h-svh items-center justify-center'>
        <Loader2 className='size-8 animate-spin' />
      </div>
    )
  }

  return (
    <>
      <UsersProvider>

          <Main fluid>
            <div className='mb-2 flex flex-wrap items-center justify-between space-y-2'>
              <div>
                <h2 className='text-2xl font-bold tracking-tight'>User & Roles</h2>
                <div className='flex gap-1'>
                  <p className='text-muted-foreground'>
                    Manage your users and their roles here.
                  </p>
                </div>
              </div>
              <UsersPrimaryButtons />
            </div>
            <Tabs
              orientation='vertical'
              defaultValue='user'
              className='space-y-4'
            >
              <div className='w-full overflow-x-auto mb-0'>
                <TabsList>
                  <TabsTrigger value='user'>User</TabsTrigger>
                  <TabsTrigger value='role'>Role</TabsTrigger>
                  <TabsTrigger value='navbar'>Navigation Bar</TabsTrigger>
                </TabsList>
              </div>
              <TabsContent value='user' className='space-y-4'>
                <div className='-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-y-0 lg:space-x-12'>
                  <UsersTable data={users} navigate={navigate} search={search} />
                </div>
              </TabsContent>
              <TabsContent value='role' className='space-y-4'>
                <ConfigurableTable tableName="users" search={search} mode="client" />
              </TabsContent>
              <TabsContent value='navbar' className='space-y-4'>

              </TabsContent>
            </Tabs>

          </Main>
          <UsersDialogs />
        </UsersProvider>
    </>
  )
}
