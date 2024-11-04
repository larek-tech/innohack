import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { AppSidebar } from '@/components/app-sidebar'
import { DashBoardPage } from '@/pages/dashboardPage'

export const Route = createFileRoute('/dash')({
    component: RouteComponent,
})

function RouteComponent() {
    return <AppSidebar>
        <DashBoardPage />
    </AppSidebar>
}
