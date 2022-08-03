<template>
  <div :class="prefixCls">
    <slot name="headerContent"></slot>
    <div :class="`${prefixCls}-content`">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { useDesign } from '/@/hooks/web/useDesign';

  const { prefixCls } = useDesign('page-wrapper');
</script>

<style lang="less">
  @prefix-cls: ~'@{namespace}-page-wrapper';

  .@{prefix-cls} {
    &-content {
      margin: 1rem;
    }
  }
</style>
