import 'virtual:svg-icons-register';
import './styles/index.less';
import { createApp } from 'vue';
import App from './App.vue';
import { router, setupRouter } from '/@/router';
import { setupRouterGuard } from '/@/router/guard';
import { setupStore } from '/@/store';
import { setupI18n } from '/@/locales/setupI18n';
import { registerGlobComp } from '/@/components/registerGlobComp';

async function bootstrap() {
  const app = createApp(App);

  // Configure store
  setupStore(app);

  // Register global components
  registerGlobComp(app);

  // I18n configuration
  setupI18n(app);

  // Configure routing
  setupRouter(app);

  // router-guard
  setupRouterGuard(router);

  // Configure wails api mocking,
  // will run when in development env and VITE_USE_MOCK is true
  if (process.env.NODE_ENV !== 'production') {
    if (import.meta.env.VITE_USE_MOCK === 'true') {
      import('../mock');
    }
  }

  // Mount app
  app.mount('#app');
}

bootstrap();
