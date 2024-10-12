import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";
import { vite as vidstack } from "vidstack/plugins";
import { visualizer } from "rollup-plugin-visualizer";

const isDev = process.env.NODE_ENV === "development";
const mode = isDev ? "development" : "production";
console.log(process.env.NODE_ENV);
console.log(mode);

// https://vitejs.dev/config/
export default defineConfig({
  base: "./",
  mode: mode,
  server: {
    port: 3000,
    proxy: {
      "/api/v1": "http://localhost:8080",
    },
  },
  plugins: [
    vue({
      template: {
        compilerOptions: {
          isCustomElement: (tag) => tag.startsWith("media-"),
        },
      },
    }),
    vidstack(),
    visualizer({ open: false }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },

  build: {
    emptyOutDir: true,
    cssTarget: "chrome61",
    sourcemap: isDev,
  },
});
