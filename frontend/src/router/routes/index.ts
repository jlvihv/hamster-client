import type { AppRouteRecordRaw, AppRouteModule } from '/@/router/types';
import { PAGE_NOT_FOUND_ROUTE } from '/@/router/routes/basic';
import { PageEnum } from '/@/enums/pageEnum';

const modules = import.meta.globEager('./modules/**/*.ts');
const routeModuleList: AppRouteModule[] = [];

Object.keys(modules).forEach((key) => {
  const mod = modules[key].default || {};
  const modList = Array.isArray(mod) ? [...mod] : [mod];
  routeModuleList.push(...modList);
});

// moduleRoutes
export const asyncRoutes = routeModuleList;

// Root
export const RootRoute: AppRouteRecordRaw = {
  path: '/',
  name: 'Root',
  redirect: PageEnum.BASE_HOME,
  meta: {
    title: 'Root',
  },
};

export const HomeRoute: AppRouteRecordRaw = {
  path: '/home',
  name: 'Home',
  component: () => import('/@/views/home/index.vue'),
  meta: {
    title: 'Home',
  },
};

// Basic routing without permission
export const basicRoutes = [PAGE_NOT_FOUND_ROUTE, RootRoute, HomeRoute, ...routeModuleList];
