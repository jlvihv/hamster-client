<template>
  <Layout :class="prefixCls">
    <LayoutSider><SideBar /></LayoutSider>
    <Layout :class="`${prefixCls}-main`">
      <LayoutHeader :class="`${prefixCls}-main-header`">
        <Header />
      </LayoutHeader>
      <LayoutContent><Content /></LayoutContent>
      <LayoutFooter><Footer /></LayoutFooter>
    </Layout>
  </Layout>
</template>

<script lang="ts" setup>
  import Header from './header/index.vue';
  import SideBar from './sider/index.vue';
  import Content from './content/index.vue';
  import Footer from './footer/index.vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { Layout, LayoutSider, LayoutHeader, LayoutContent, LayoutFooter } from 'ant-design-vue';

  const { prefixCls } = useDesign('layout');
</script>

<style lang="less">
  @prefix-cls: ~'@{namespace}-layout';

  .@{prefix-cls} {
    min-height: 100%;

    &-main {
      background-color: #f7f3ee;

      &-header {
        background-color: @white;
        padding: 0 1rem;
      }
    }
  }
</style>
