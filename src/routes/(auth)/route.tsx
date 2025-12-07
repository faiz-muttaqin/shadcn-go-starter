import { createFileRoute, Outlet} from '@tanstack/react-router'

export const Route = createFileRoute('/(auth)')({
  component: AuthGuardLayout,

})

function AuthGuardLayout() {
  return <Outlet />
}