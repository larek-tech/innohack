import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import BaseChatPage from '@/pages/baseChatPage'
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
      <ChatInterface sessionId={{ sessionId: "123" }} />
    </div>
  </RequireAuth>
}
