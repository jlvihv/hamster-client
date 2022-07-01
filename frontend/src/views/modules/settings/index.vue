<template>
  <PageWrapper>
    <div class="bg-white p-3">
      <div class="text-primary text-[30px]">Setting</div>
      <Button type="primary">YES {{ t('common.loadingText') }}</Button>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Button } from 'ant-design-vue';

  const { t } = useI18n();
</script>
