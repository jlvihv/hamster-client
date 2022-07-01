import { watch, unref } from 'vue';
import { useI18n } from '/@/hooks/web/useI18n';
import { useTitle as usePageTitle } from '@vueuse/core';
import { useRouter } from 'vue-router';

import { REDIRECT_NAME } from '/@/router/constant';

/**
 * Listening to page changes and dynamically changing site titles
 */
export function useTitle() {
  const title = import.meta.env.VITE_APP_TITLE;
  const { t, locale } = useI18n();
  const { currentRoute } = useRouter();

  const pageTitle = usePageTitle();

  watch(
    [() => currentRoute.value.path, () => locale.value],
    () => {
      const route = unref(currentRoute);

      if (route.name === REDIRECT_NAME) {
        return;
      }

      const tTitle = t(route?.meta?.title as string);
      pageTitle.value = tTitle ? ` ${tTitle} - ${title} ` : `${title}`;
    },
    { immediate: true },
  );
}
