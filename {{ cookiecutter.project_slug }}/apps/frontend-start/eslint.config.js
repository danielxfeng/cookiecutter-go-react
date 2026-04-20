//  @ts-check

import { tanstackConfig } from '@tanstack/eslint-config';
import { defineConfig, globalIgnores } from 'eslint/config';

export default defineConfig([
  globalIgnores(['dist', 'src/components/ui/**', 'src/routeTree.gen.ts']),
  ...tanstackConfig,
]);
