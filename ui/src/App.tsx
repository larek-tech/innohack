import { createRouter, RouterProvider } from "@tanstack/react-router"
import { routeTree } from "./routeTree.gen";
import "./index.css"
import { AuthProvider } from "./auth/AuthProvider";
import { Toaster } from "./components/ui/toaster";
import { SidebarProvider } from "./components/ui/sidebar";
const router = createRouter({ routeTree });


declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

function App() {

  return <>
    <Toaster />
    <AuthProvider>
      <SidebarProvider>
        <RouterProvider router={router} />
      </SidebarProvider>
    </AuthProvider>
  </>

}

export default App
