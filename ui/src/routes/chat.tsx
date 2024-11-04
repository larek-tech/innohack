
import { createFileRoute } from '@tanstack/react-router'

import { RequireAuth } from '@/auth/RequireAuth'
import ChatInterface from '@/pages/chatPage'
import { AppSidebar } from '@/components/app-sidebar'

export const Route = createFileRoute('/chat')({
  component: RouteComponent,
})

function RouteComponent() {
  return <RequireAuth>
    <div className="flex h-screen w-full">
      <AppSidebar />
      <ChatInterface />
    </div>
  </RequireAuth>
}
