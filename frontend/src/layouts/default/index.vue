<template>
  <Layout :class="prefixCls">
    <LayoutContent><Content /></LayoutContent>
  </Layout>
</template>

<script lang="ts" setup>
  import Content from './content/index.vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { Layout, LayoutContent } from 'ant-design-vue';

  const { prefixCls } = useDesign('layout');
</script>

<style lang="less">
  @prefix-cls: ~'@{namespace}-layout';

  .@{prefix-cls} {
    min-height: 100%;

    &-main {
      background-color: linear-gradient(132deg, #f4f5fa 0%, #f8f9fc 100%, #f8f9fc 100%);

      &-header {
        background-color: @white;
        padding: 0 1rem;
      }

      &-footer {
        background-color: #f7f3ee;
      }
    }
  }
</style>
