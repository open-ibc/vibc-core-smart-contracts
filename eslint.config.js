import globals from 'globals';
import pluginJs from '@eslint/js';
// import tseslint from "typescript-eslint";
import tseslint, { parser } from 'typescript-eslint'

export default [
  {files: ['**/src/*.{ts}']},
  {ignores:['**/lib/**', '**/node_modules/**', '**/dist/**']},
  {languageOptions: { globals: globals.node }},
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  {
    rules:{
      '@typescript-eslint/quotes': ['error', 'single', { avoidEscape: true }],
      '@/quotes': ['error', 'single', { avoidEscape: true }],
      '@/jsx-quotes': ['error', 'prefer-double'],
      '@/eol-last': ['error', 'always'],
      '@/no-unused-vars': ['warn'],
      '@typescript-eslint/no-unused-vars': ['warn'],
      '@/no-undef': ['off'],
      '@/no-case-declarations': ['off'],
    }
  }
];
