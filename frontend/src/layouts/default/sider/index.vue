<template>
  <div :class="prefixCls">
    <div :class="`${prefixCls}-logo`">
      <div class="text-white text-2xl">{{ appTitle }}</div>
    </div>
    <LayoutMenu />
  </div>
</template>

<script lang="ts" setup>
  import LayoutMenu from '../menu/index.vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { appTitle } from '/@/settings/siteSetting';

  const { prefixCls } = useDesign('layout-sider');
</script>

<style lang="less">
  @prefix-cls: ~'@{namespace}-layout-sider';

  .@{prefix-cls} {
    &-logo {
      margin-top: 20px;
      margin-bottom: 50px;
      display: flex;
      justify-content: center;
    }
  }
</style>
