import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    host: "frontend.local",
    open: true,
    allowedHosts: [
      "frontend.local",
      "localhost",
      "127.0.0.1",
      "api.frontend.local",
    ],
  },
});
