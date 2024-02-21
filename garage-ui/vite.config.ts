import { defineConfig } from "vite";
import { splitVendorChunkPlugin } from "vite";
import vuetify from "vite-plugin-vuetify";
import vue from "@vitejs/plugin-vue";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  base: "/",
  plugins: [vue(), vuetify(), splitVendorChunkPlugin()],
  server: {
    host: "0.0.0.0",
    port: 3000,
    strictPort: true,
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  build: {
    // outDir: path.resolve(__dirname, "./src/assets/ui"),
    cssCodeSplit: true,
  },
});
