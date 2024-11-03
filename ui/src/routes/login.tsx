import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { LoginForm } from '@/view/loginForm'


export const Route = createFileRoute('/login')({
    component: RouteComponent,
})

function RouteComponent() {
    return <div className="w-full h-screen flex items-center justify-center">
        <LoginForm />
    </div>
}