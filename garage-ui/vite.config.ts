import { defineConfig } from "vite";
import { splitVendorChunkPlugin } from "vite";
import vuetify from "vite-plugin-vuetify";
import vue from "@vitejs/plugin-vue";
import path from "path";
import { visualizer } from "rollup-plugin-visualizer";

const buildReportplugin =
  process.env.npm_lifecycle_event === "build:report"
    ? visualizer({
        open: true,
        brotliSize: true,
        gzipSize: true,
        filename: "dist/report.html",
      })
    : null;

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vuetify(), splitVendorChunkPlugin(), buildReportplugin],
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
