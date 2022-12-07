import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: "127.0.0.1",
    port: process.env.SERVICE_PORT ?? 3000,
    watch: {
      usePolling: true,
    }
  }
})
