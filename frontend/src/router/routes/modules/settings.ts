import type { AppRouteModule } from '/@/router/types';
import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const settings: AppRouteModule = {
  path: '/settings',
  name: 'Settings',
  component: LAYOUT,
  redirect: '/settings/index',
  meta: {
    hideChildrenInMenu: true,
    icon: 'ep:setting',
    title: t('routes.settings.settings'),
    orderNo: 200000,
  },
  children: [
    {
      path: 'index',
      name: 'SettingsIndex',
      component: () => import('/@/views/modules/settings/index.vue'),
      meta: {
        title: t('routes.settings.settings'),
        hideMenu: true,
      },
    },
  ],
};

export default settings;
