import { Header } from '@/components/Header';
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/editor/theme')({
  component: RouteComponent,
})

function RouteComponent() {
  
  return (
    <div className="relative isolate flex h-svh flex-col overflow-hidden">
      <Header />
      <main className="isolate flex flex-1 flex-col overflow-hidden">
        <Outlet />
      </main>
    </div>
  );
}
