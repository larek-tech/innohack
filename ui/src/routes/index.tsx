import { Button } from '@/components/ui/button'
import { Landing } from '@/pages/landing'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
    component: () => (
        <Landing />
    ),
})
