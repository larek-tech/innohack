import { createFileRoute } from '@tanstack/react-router'
import { LoginForm } from '@/view/loginForm'


export const Route = createFileRoute('/login')({
    component: RouteComponent,
})

function RouteComponent() {
    return <LoginForm />

}
