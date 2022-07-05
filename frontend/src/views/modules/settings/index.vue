<template>
  <PageWrapper>
    <div class="p-3 bg-white">
      <div class="text-color-[#141212] text-xl font-bold mb-8 ml-5 mt-3">{{
        t('routes.settings.settings')
      }}</div>
      <Form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        :label-col="{ style: { width: '150px' } }"
      >
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
          <Button type="primary" class="w-[50%] ml-[25%] mt-16" @click="handleWsurl">{{
            t('routes.settings.save')
          }}</Button>
        </FormItem>
      </Form>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { reactive, computed, ref } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SettingWsUrl } from '/@wails/go/app/Setting';
  import { Button, Form, FormItem, Input } from 'ant-design-vue';

  const { t } = useI18n();
  const formRef = ref();

  const formData = reactive<{
    wsurl?: string;
    address?: string;
    balance?: string;
  }>({});

  const formRules = computed(() => ({
    wsurl: [{ message: t('routes.settings.pleaseInputWsUrl'), trigger: 'change', required: true }],
  }));

  async function handleWsurl() {
    await formRef.value?.validate();
    if (!formData.wsurl) return;

    try {
      const data = await SettingWsUrl(formData.wsurl);
      console.log(data);
    } catch (err) {
      console.log('error', err);
    }
  }
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
