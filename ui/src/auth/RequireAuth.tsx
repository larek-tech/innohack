import { SidebarProvider } from '@/components/ui/sidebar';
import { useAuth } from '.';
import { useNavigate } from '@tanstack/react-router'
import { AppSidebar } from '@/components/app-sidebar';


export function RequireAuth({ children }: { children: JSX.Element }) {
    const auth = useAuth();
    const navigate = useNavigate()
    if (!auth.user) {
        // Navigate to login if the user is not authenticated
        navigate({
            to: "/login"
        });
        return null; // Return null so nothing renders while navigating
    }

    return <SidebarProvider>
        { children }
    </SidebarProvider>;
}
