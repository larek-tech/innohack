import * as React from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { RequireAuth } from '@/auth/RequireAuth'
import ChatInterface from '@/pages/chatPage'
import { type } from 'arktype'
import Layout from '@/layout/layout'

const chatSearchSchema = type({
  chatId: 'number = 1',
})

export const Route = createFileRoute('/chat')({
  component: RouteComponent,
  validateSearch: chatSearchSchema,
})




function RouteComponent() {

  return (
    <RequireAuth>
      <ChatInterface />
    </RequireAuth>
  )
}
