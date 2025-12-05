import type { KnipConfig } from 'knip';

const config: KnipConfig = {
  ignore: [
    'src/components/ui/**',
    'src/routeTree.gen.ts',
    // ignore backend and non-frontend folders so knip focuses on frontend src
    'backend/**',
    'tmp/**',
    'log/**',
    'scripts/**',
    'dist/**',
    'bin/**',
    'public/**',
  ],
  ignoreDependencies: [
    'tailwindcss',
    'tw-animate-css',
    // add backend-only deps here if any
  ],
  // If knip supports ignoring binaries, add them here (sleep, air, go, cz)
  ignoreBinaries: ['sleep', 'air', 'go', 'cz'],
};

export default config;