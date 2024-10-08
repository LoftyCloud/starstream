import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
const path = require('path');

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    sourcemap: true  // 确保开启了源映射
  },
  plugins: [vue()],
  server:{
    port: 3000
  },
  resolve: {
    alias: { '@': path.resolve(__dirname, './src'),},  // 配置 @ 别名
  },
})
