import { defineConfig, loadEnv, UserConfigExport } from 'vite'
import react from '@vitejs/plugin-react'
import path from "path"
import { TanStackRouterVite } from "@tanstack/router-vite-plugin"

// Define the type for the configuration parameters
interface ViteConfigParams {
  mode: string;
}

// https://vite.dev/config/
export default ({ mode }: ViteConfigParams): UserConfigExport => {
  // Load app-level env vars to node-level env vars.
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return defineConfig({
    plugins: [react(), TanStackRouterVite()],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    }
  })
}