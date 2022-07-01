import type { Menu } from '/@/router/types';
import { asyncRoutes } from '/@/router/routes';
import { transformRouteToMenu } from '/@/router/helper/menuHelper';

export function getMenus(): Menu[] {
  const menuList = transformRouteToMenu(asyncRoutes, true);

  menuList.sort((a, b) => {
    return (a.meta?.orderNo || 0) - (b.meta?.orderNo || 0);
  });

  return menuList;
}
