import Editor from '@/features/editor-theme/editor'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/editor/theme/')({
    component: RouteComponent,
})

function RouteComponent() {
    return <Editor themePromise={Promise.resolve(null)} />
}
