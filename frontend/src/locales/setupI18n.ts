import type { App } from 'vue';
import type { I18n, I18nOptions } from 'vue-i18n';
import { createI18n } from 'vue-i18n';

// langs
import EnLang from './lang/en';

enum Locale {
  EN = 'en',
}

function createI18nOptions(): I18nOptions {
  return {
    legacy: false,
    locale: Locale.EN,
    fallbackLocale: Locale.EN,
    messages: {
      [Locale.EN]: EnLang.message,
    },
    availableLocales: [Locale.EN],
    sync: true, // Inherit locale from global scope
    silentTranslationWarn: true, // true - warning off
    missingWarn: false,
    silentFallbackWarn: true,
  };
}

// setup i18n instance with glob
export let i18n: ReturnType<typeof createI18n>;
export function setupI18n(app: App) {
  const options = createI18nOptions();
  i18n = createI18n(options) as I18n;
  app.use(i18n);
}
