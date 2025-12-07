import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/editor/theme/')({
    component: RouteComponent,
})

function RouteComponent() {
    return <div>Hello "/"!</div>
}
