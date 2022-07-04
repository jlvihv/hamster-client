<template>
  <PageWrapper>
    <div class="p-3 bg-white">
      <div class="text-color-[#141212] text-xl font-bold mb-8 ml-5 mt-3">{{
        t('routes.settings.settings')
      }}</div>
      <Form :model="formData" :label-col="{ style: { width: '150px' } }">
        <FormItem :label="t('routes.settings.pleaseInputWsUrl')" name="wsurl">
          <Input
            v-model:value="formData.wsurl"
            :placeholder="t('routes.settings.pleaseInputWsUrl')"
          />
        </FormItem>
        <FormItem :label="t('routes.settings.address')" name="address">
          <span class="text-[#666666]">{{ formData.address }}</span>
          <Button class="ml-[50%]">{{ t('routes.settings.unbind') }}</Button>
        </FormItem>
        <FormItem :label="t('routes.settings.accountBalance')" name="balance">
          <span class="text-[#666666]">{{ formData.balance }}</span>
        </FormItem>
        <FormItem>
          <Button type="primary" class="w-[50%] ml-[25%] mt-16">{{
            t('routes.settings.save')
          }}</Button>
        </FormItem>
      </Form>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { reactive } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Button, Form, FormItem, Input } from 'ant-design-vue';

  const { t } = useI18n();

  const formData = reactive<{
    wsurl?: string;
    address: string;
    balance: string;
  }>({ address: 'ws://183.66.65.207:49944', balance: 'ws://183.66.65.207:49944' });
</script>

<style lang="less" scoped>
  .ant-input {
    @apply w-1/2 rounded text-[#666666];
  }

  .ant-btn {
    @apply rounded;
  }

  .ant-btn-primary {
    @apply h-10;
  }

  :deep(.ant-form-item-label > label) {
    @apply text-[#666666];
  }
</style>
