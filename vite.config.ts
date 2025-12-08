import path from 'path'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'
import { tanstackRouter } from '@tanstack/router-plugin/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    tanstackRouter({
      target: 'react',
      autoCodeSplitting: true,
    }),
    react(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    hmr: {
      timeout: 30000, // Increase from default
    },
  },
  optimizeDeps: {
    // Explicitly list packages to pre-bundle. Using a glob string here causes
    // unpredictable pre-bundling behavior and may slow down or fail the dev server.
    include: [
      'react',
      'react-dom',
      // 'react/jsx-runtime',
      // 'react/jsx-dev-runtime',
      // 'react-dom/client',
      // 'axios',
      // 'sonner',
      // '@tanstack/react-query',
      // '@tanstack/react-router',
      // // Radix primitives used by the project - list them explicitly
      // '@radix-ui/react-tooltip',
      // '@radix-ui/react-dialog',
      // '@radix-ui/react-popover',
      // '@radix-ui/react-dropdown-menu',
      // '@radix-ui/react-select',
      // '@radix-ui/react-radio-group',
      // '@radix-ui/react-avatar',
      // '@radix-ui/react-checkbox',
      // '@radix-ui/react-switch',
      // '@radix-ui/react-tabs',
      // '@radix-ui/react-scroll-area',
      // '@radix-ui/react-accordion',
      // '@radix-ui/react-alert-dialog',
      // '@radix-ui/react-label',
      // '@radix-ui/react-separator',
      // '@radix-ui/react-slot',
    ],
    // force: true,
  },
})
