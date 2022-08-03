<template>
  <div :class="prefixCls">
    <div :class="`${prefixCls}-logo`">
      <img :src="logoImage" />
    </div>
    <LayoutMenu />
  </div>
</template>

<script lang="ts" setup>
  import LayoutMenu from '../menu/index.vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import logoImage from '/@/assets/images/logo.png';

  const { prefixCls } = useDesign('layout-sider');
</script>

<style lang="less">
  @prefix-cls: ~'@{namespace}-layout-sider';

  .@{prefix-cls} {
    &-logo {
      margin-top: 20px;
      padding: 20px;
      display: flex;
      justify-content: center;
    }
  }
</style>
