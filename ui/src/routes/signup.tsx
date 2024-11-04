import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { SignUp } from '@/view/signupForm'

export const Route = createFileRoute('/signup')({
    component: RouteComponent,
})

function RouteComponent() {
    return <div className="w-full h-screen flex items-center justify-center">
        <SignUp />
    </div>
}
