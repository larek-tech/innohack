import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import BaseChatPage from '@/pages/baseChatPage'
import { RequireAuth } from '@/auth/RequireAuth'
import ChatInterface from '@/pages/chatPage'

export const Route = createFileRoute('/chat')({
  component: RouteComponent,
})

function RouteComponent() {
  return <RequireAuth>
    <ChatInterface />
  </RequireAuth>
}
