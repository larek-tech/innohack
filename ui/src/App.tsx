import { createRouter, RouterProvider } from "@tanstack/react-router"
import { routeTree } from "./routeTree.gen";
import "./index.css"
import { AuthProvider } from "./auth/AuthProvider";
import { Toaster } from "./components/ui/toaster";
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
      <RouterProvider router={router} />
    </AuthProvider>
  </>

}

export default App
