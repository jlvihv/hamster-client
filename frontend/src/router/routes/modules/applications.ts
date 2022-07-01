import type { AppRouteModule } from '/@/router/types';
import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const applications: AppRouteModule = {
  path: '/applications',
  name: 'Applications',
  component: LAYOUT,
  redirect: '/applications/index',
  meta: {
    hideChildrenInMenu: true,
    icon: 'simple-icons:about-dot-me',
    title: t('routes.applications.applications'),
    orderNo: 100000,
  },
  children: [
    {
      path: 'index',
      name: 'ApplicationsIndex',
      component: () => import('/@/views/modules/applications/index/index.vue'),
      meta: {
        title: t('routes.applications.applications'),
        icon: 'simple-icons:about-dot-me',
        hideMenu: true,
      },
    },
    {
      path: 'new',
      name: 'AddApplication',
      component: () => import('/@/views/modules/applications/application-add/index.vue'),
      meta: {
        title: t('routes.applications.addApplication'),
        icon: 'simple-icons:about-dot-me',
        hideMenu: true,
      },
    },
    {
      path: ':id/edit',
      name: 'EditApplication',
      component: () => import('/@/views/modules/applications/application-edit/index.vue'),
      meta: {
        title: t('routes.applications.editApplication'),
        icon: 'simple-icons:about-dot-me',
        hideMenu: true,
      },
    },
    {
      path: ':id',
      name: 'ShowApplication',
      component: () => import('/@/views/modules/applications/application/index.vue'),
      meta: {
        title: t('routes.applications.showApplication'),
        icon: 'simple-icons:about-dot-me',
        hideMenu: true,
      },
    },
  ],
};

export default applications;
