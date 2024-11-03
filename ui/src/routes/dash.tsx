import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { DashBoardPage } from '@/pages/dashboardPage'
import { SidebarProvider } from '@/components/ui/sidebar'

export const Route = createFileRoute('/dash')({
    component: RouteComponent,
})

function RouteComponent() {
    console.log("dash")
    return <SidebarProvider>
        <DashBoardPage />
    </SidebarProvider>
}
