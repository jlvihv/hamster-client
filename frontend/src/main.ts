import 'virtual:svg-icons-register';
import './styles/index.less';
import { createApp } from 'vue';
import App from './App.vue';
import { setupRouter } from '/@/router';
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

  // Mount app
  app.mount('#app');
}

bootstrap();
