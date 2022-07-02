<template>
  <div v-if="menus.length > 0" :class="prefixCls">
    <Menu mode="inline" theme="dark" v-model:selectedKeys="selectedKeys" @click="handleMenuClick">
      <template v-for="item in menus.filter((x) => !x.hideMenu)">
        <SubMenu
          v-if="item.children && item.children.length > 0 && !item.hideChildrenInMenu"
          :key="item.name"
        >
          <template #icon v-if="item.icon">
            <Icon :icon="item.icon" />
          </template>
          <template #title>{{ getMenuTitle(item) }}</template>
          <MenuItem
            :key="subItem.name"
            v-for="subItem in item.children?.filter((x) => !x.hideMenu) || []"
          >
            <template #icon v-if="subItem.icon">
              <Icon :icon="subItem.icon" />
            </template>
            {{ getMenuTitle(subItem) }}
          </MenuItem>
        </SubMenu>
        <MenuItem :key="item.name" v-else>
          <template #icon v-if="item.icon">
            <Icon :icon="item.icon" />
          </template>
          {{ getMenuTitle(item) }}
        </MenuItem>
      </template>
    </Menu>
  </div>
</template>

<script lang="ts" setup>
  import type { Menu as MenuType } from '/@/router/types';
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { Icon } from '/@/components/Icon';
  import { getMenus } from '/@/router/menus';
  import { REDIRECT_NAME } from '/@/router/constant';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { listenerRouteChange } from '/@/utils/routeChange';
  import { Menu, SubMenu, MenuItem } from 'ant-design-vue';

  const { prefixCls } = useDesign('layout-menu');
  const { t } = useI18n();

  const menus = getMenus();
  const router = useRouter();

  // selectedKeys
  const selectedKeys = ref<string[]>([]);

  const getMenuTitle = (item: { title?: string; name: string }) =>
    item.title ? t(item.title) : item.name;

  // Set selectedKeys
  listenerRouteChange((route) => {
    if (route.name === REDIRECT_NAME) return;

    const { matched } = route;
    const currentSelectedKeys: any[] = [];

    // `return true` means break loop
    matched.some((item) => {
      if (item.meta.hideMenu) return true;
      if (item.name) currentSelectedKeys.push(item.name);
      if (item.meta.hideChildrenInMenu) return true;
    });

    selectedKeys.value = currentSelectedKeys;
  });

  const handleMenuClick = ({ keyPath }: { keyPath?: any[] }) => {
    if (!keyPath) return;

    let item: MenuType | undefined;
    let keyMenus = menus;

    keyPath.forEach((key) => {
      item = keyMenus.find((x) => x.name === key);
      keyMenus = item?.children || [];
    });

    if (item) {
      router.push(item.path);
    }
  };
</script>
