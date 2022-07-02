import 'virtual:svg-icons-register';
import './styles/index.less';
import { createApp } from 'vue';
import App from './App.vue';
import { router, setupRouter } from '/@/router';
import { setupRouterGuard } from '/@/router/guard';
import { setupI18n } from '/@/locales/setupI18n';
import { registerGlobComp } from '/@/components/registerGlobComp';

async function bootstrap() {
  const app = createApp(App);

  // Register global components
  registerGlobComp(app);

  // I18n configuration
  setupI18n(app);

  // Configure routing
  setupRouter(app);

  // router-guard
  setupRouterGuard(router);

  // Mount app
  app.mount('#app');
}

bootstrap();
