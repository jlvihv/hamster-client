<template>
  <PageWrapper>
    <Card>
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
              :default-value="wsUrl"
              v-model:value="formData.wsUrl"
              :placeholder="t('routes.settings.pleaseInputWsUrl')"
            />
          </FormItem>
          <FormItem :label="t('routes.settings.address')" name="address" v-if="address">
            <span class="text-[#666666]">{{ address }}</span>
            <Button class="ml-[50%]" @click="handleUnbind">{{
              t('routes.settings.unbind')
            }}</Button>
          </FormItem>
          <FormItem :label="t('routes.settings.accountBalance')" name="balance" v-if="address">
            <span class="text-[#666666]">{{ balance }}</span>
          </FormItem>
          <FormItem>
            <Button type="primary" class="w-[50%] ml-[25%] mt-16" @click="handleSave">{{
              t('routes.settings.save')
            }}</Button>
          </FormItem>
        </Form>
      </div>
    </Card>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import '@polkadot/api-augment';
  import { reactive, computed, ref, onMounted, watch } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useSettingStore } from '/@/store/modules/setting';
  import { createPolkadotApi, formatBalance } from '/@/utils/polkadotUtil';
  import { Button, Form, FormItem, Input, Card } from 'ant-design-vue';

  const { t } = useI18n();
  const settingStore = useSettingStore();

  const address = computed(() => settingStore.walletInfo?.address);
  const wsUrl = computed(() => settingStore.config?.wsUrl);
  const balance = ref('');

  watch([address, wsUrl], async ([addressVal, wsUrlVal]) => {
    if (!addressVal || !wsUrlVal) return;
    const api = await createPolkadotApi(wsUrlVal);
    const { data: balanceData } = await api.query.system.account(address);
    console.log(balanceData);
    balance.value = formatBalance(balanceData.free);
  });

  const formRef = ref();
  const formData = reactive<{
    wsUrl?: string;
  }>({});
  const formRules = computed(() => ({
    wsUrl: [{ message: t('routes.settings.pleaseInputWsUrl'), trigger: 'change', required: true }],
  }));

  async function handleUnbind() {
    try {
      settingStore.deleteWalletAction();
    } catch (err) {
      console.log('error', err);
    }
  }

  async function handleSave() {
    await formRef.value?.validate();
    if (!formData.wsUrl) return;

    try {
      settingStore.saveWsUrlAction(formData.wsUrl);
    } catch (err) {
      console.log('error', err);
    }
  }

  onMounted(() => {
    settingStore.getConfigAction();
  });
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
