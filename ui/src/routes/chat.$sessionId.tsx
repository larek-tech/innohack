import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { RequireAuth } from '@/auth/RequireAuth'
import ChatInterface from '@/pages/chatPage'
import { type } from 'arktype'

const chatSearchSchema = type({
  chatId: 'number = 1',
})

export const Route = createFileRoute('/chat/$sessionId')({
  component: RouteComponent,
  validateSearch: chatSearchSchema,
})

function RouteComponent() {
  console.log("we are at the chat session id route")
  return (
    <RequireAuth>
      <ChatInterface />
    </RequireAuth>
  )
}
