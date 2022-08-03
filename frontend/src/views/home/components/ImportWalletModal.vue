<template>
  <div class="text-[#222222] text-[40px] font-bold">{{ t('home.importWallet') }}</div>
  <div class="mt-[40px]">
    <SvgIcon
      size="16"
      color="#F77522"
      name="link"
    />5EKQCEm834AMHZ2CvdWeQUP5QrzBJzYa5PQrzBJzYa5PnhByhitgK9RS5o
    <SvgIcon size="16" color="#BBC0C5" name="close" />
  </div>
  <div class="my-[20px] font-bold !text-left">{{ t('home.name') }}：Client0729</div>
  <div class="!text-left">5EKQCEm834AMHZ2CvdWeQUP5QrzBJzYa5PQrzBJzYa5PnhByhitgK9RS5o</div>
  <div class="!text-left my-[20px]">
    <InputPassword :placeholder="t('home.pwd')" />
  </div>
</template>

<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { InputPassword } from 'ant-design-vue';

  const { t } = useI18n();
</script>
<style lang="less" scoped>
  :deep(.ant-input-affix-wrapper) {
    @apply !rounded-[8px] !h-[42px];
  }
</style>
