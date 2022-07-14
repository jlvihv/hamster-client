import type { UserConfig, ConfigEnv } from 'vite';
import { loadEnv } from 'vite';
import { resolve } from 'path';
import { NodeGlobalsPolyfillPlugin } from '@esbuild-plugins/node-globals-polyfill';
import { NodeModulesPolyfillPlugin } from '@esbuild-plugins/node-modules-polyfill';
import rollupNodePolyfillPlugin from 'rollup-plugin-polyfill-node';
import { wrapperEnv, nodePolyfillAlias } from './build/utils';
import { generateModifyVars } from './build/generate/generateModifyVars';
import { createVitePlugins } from './build/vitePlugins';

const root = () => process.cwd();
const pathResolve = (dir: string) => resolve(root(), '.', dir);

export default ({ command, mode }: ConfigEnv): UserConfig => {
  const isBuild = command === 'build';
  const env = loadEnv(mode, root());
  const viteEnv = wrapperEnv(env);
  const { VITE_PORT, VITE_DROP_CONSOLE, VITE_PUBLIC_PATH } = viteEnv;

  return {
    base: VITE_PUBLIC_PATH,
    root: root(),
    resolve: {
      alias: {
        '/@/': pathResolve('src') + '/',
        '/@wails/': pathResolve('wailsjs') + '/',
        'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js',
        ...nodePolyfillAlias(),
      },
    },
    server: {
      host: true,
      port: VITE_PORT,
    },
    esbuild: {
      pure: VITE_DROP_CONSOLE ? ['console.log', 'debugger'] : [],
    },
    build: {
      target: 'es2016',
      cssTarget: 'chrome80',
      outDir: 'dist',
      brotliSize: false,
      chunkSizeWarningLimit: 5000,
      // sourcemap: true,
      rollupOptions: {
        plugins: [rollupNodePolyfillPlugin()],
      },
    },
    define: {
      // setting vue-i18-next
      // Suppress warning
      __INTLIFY_PROD_DEVTOOLS__: false,
    },
    css: {
      preprocessorOptions: {
        less: {
          modifyVars: generateModifyVars(),
          javascriptEnabled: true,
        },
      },
    },

    // The vite plugin used by the project. The quantity is large, so it is separately extracted and managed
    plugins: createVitePlugins(viteEnv, isBuild),

    optimizeDeps: {
      include: ['ant-design-vue/es/locale/zh_CN', 'ant-design-vue/es/locale/en_US'],
      esbuildOptions: {
        // Node.js global to browser globalThis
        define: {
          global: 'globalThis',
        },
        // Enable esbuild polyfill plugins
        plugins: [
          NodeGlobalsPolyfillPlugin({ buffer: true, process: true }),
          NodeModulesPolyfillPlugin(),
        ],
      },
    },
  };
};
