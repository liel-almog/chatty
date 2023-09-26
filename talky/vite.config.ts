import react from "@vitejs/plugin-react-swc";
import { setDefaultResultOrder } from "dns";
import hash from "hash-it";
import * as path from "path";
import { defineConfig } from "vite";
import Inspect from "vite-plugin-inspect";
import svgLoader from "vite-svg-loader";
import tsonfigpathes from "vite-tsconfig-paths";

setDefaultResultOrder("verbatim");
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), Inspect(), tsonfigpathes(), svgLoader()],
  css: {
    modules: {
      generateScopedName: (name, filePath) => {
        const matches = path.basename(filePath).match(/^([a-z-]+)(.module)?.s?css/);
        if (!matches) {
          throw new Error("Could not match filename");
        }

        const hashValue = hash(filePath);
        const baseFilename = matches[1];
        return `${baseFilename}-${name}-${hashValue}`;
      },
      localsConvention: "camelCaseOnly",
    },
  },
  server: {
    port: 3000,
    open: false,
    strictPort: true,
    host: true,
  },
});
