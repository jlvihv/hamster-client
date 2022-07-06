<template>
  <PageWrapper>
    <Card>
      <div class="text-color-[#141212] text-xl font-bold mb-8 ml-5 mt-3">
        {{ t('settings.index.settings') }}
      </div>
      <Form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        :label-col="{ style: { width: '150px' } }"
      >
        <FormItem :label="t('settings.index.pleaseInputWsUrl')" name="wsUrl">
          <Input
            :default-value="wsUrl"
            v-model:value="formData.wsUrl"
            :placeholder="t('settings.index.pleaseInputWsUrl')"
          />
        </FormItem>
        <FormItem :label="t('settings.index.address')" name="address" v-if="address">
          <span class="text-[#666666]">{{ address }}</span>
          <Button class="ml-[50%]" @click="handleUnbind">{{ t('settings.index.unbind') }}</Button>
        </FormItem>
        <FormItem
          :label="t('settings.index.accountBalance')"
          name="balance"
          v-if="wsUrl && address"
        >
          <LoadingOutlined v-if="balance.loading" />
          <span class="text-[#666666]" v-else>{{ balance.value }}</span>
        </FormItem>
        <FormItem>
          <Button type="primary" class="w-[50%] ml-[25%] mt-16" @click="handleSave">
            {{ t('settings.index.save') }}
          </Button>
        </FormItem>
      </Form>
    </Card>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import '@polkadot/api-augment';
  import { reactive, computed, ref, onMounted, watchEffect } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useSettingStore } from '/@/store/modules/setting';
  import { createPolkadotApi, formatBalance } from '/@/utils/polkadotUtil';
  import { LoadingOutlined } from '@ant-design/icons-vue';
  import { Button, Form, FormItem, Input, Card } from 'ant-design-vue';

  const { t } = useI18n();
  const settingStore = useSettingStore();

  const address = computed(() => settingStore.walletInfo?.address);
  const wsUrl = computed(() => settingStore.config?.wsUrl);
  const balance = reactive({ loading: false, value: '' });

  // Fetching balance once wsUrl and address changed
  watchEffect(async (onInvalidate) => {
    if (!address.value || !wsUrl.value) return;

    balance.loading = true;

    const api = await createPolkadotApi(wsUrl.value);
    const { data: balanceData } = await api.query.system.account(address.value);

    balance.value = formatBalance(balanceData.free);
    balance.loading = false;

    onInvalidate(() => api?.disconnect());
  });

  const formRef = ref();
  const formData = reactive<{ wsUrl?: string }>({});
  const formRules = computed(() => ({
    wsUrl: [{ message: t('settings.index.pleaseInputWsUrl'), trigger: 'change', required: true }],
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
