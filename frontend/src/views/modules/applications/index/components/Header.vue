<template>
  <div class="flex justify-between mb-[20px] relative">
    <div>
      <SvgIcon
        v-if="showBack"
        @click="onBack"
        class="cursor-pointer"
        color="#858B92"
        size="20"
        name="left"
      />
    </div>
    <div>
      <Popover placement="bottomRight" arrow-point-at-center v-model:visible="showPeople" trigger="click">
        <template #content>
          <div class="w-[300px]">
            <div class="title-div">
              <SvgIcon size="16" color="#63A0FA" name="address" />
              {{ t('applications.index.address') }}
            </div>
            <div class="text-[12px] mb-[20px] mx-[20px] whitespace-normal break-all">
              {{ settingStore.walletInfo?.address }}
            </div>
            <div class="title-div">
              <SvgIcon size="16" color="#63A0FA" name="balance" />
              {{ t('applications.index.balance') }}
            </div>
            <div class="text-[12px] mb-[20px] mx-[20px]">
              <template v-if="balance.loading">
                <LoadingOutlined />
              </template>
              <template v-else>
                {{ balance.value }}
              </template>
            </div>
            <div class="button-div">
              <router-link :to="{ path: '/home', query: { step: 1 } }">
                <Button type="primary">{{ t('applications.index.changeWallet') }}</Button>
              </router-link>
            </div>
          </div>
        </template>
        <SvgIcon
          class="cursor-pointer border-none"
          color="#858B92"
          size="20"
          name="people"
        />
      </Popover>
      <Popover placement="bottomRight" arrow-point-at-center v-model:visible="showSetting" trigger="click">
        <template #content>
          <Form ref="formRef" :model="formData" :rules="formRules" class="w-[300px]">
            <div class="title-div">
              <SvgIcon size="16" color="#63A0FA" name="ws" />
              {{ t('applications.index.wsUrl') }}
            </div>
            <div>
              <Select
                class="w-full"
                :allowClear="true"
                :placeholder="t('applications.index.wsUrlPlaceholder')"
                :options="urlOptions"
                v-model:value="formData.wsUrl"
              />
            </div>
            <div class="button-div">
              <Button type="primary" @click="handleWsUrlSave">{{ t('common.saveText') }}</Button>
            </div>
          </Form>
        </template>
        <SvgIcon
          class="cursor-pointer ml-[20px]"
          color="#858B92"
          size="20"
          name="setting"
        />
      </Popover>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { reactive, ref, computed, watchEffect } from 'vue';
  import { SvgIcon } from '/@/components/Icon';
  import { useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useSettingStore } from '/@/store/modules/setting';
  import { LoadingOutlined } from '@ant-design/icons-vue';
  import { createPolkadotApi, formatBalance } from '/@/utils/polkadotUtil';
  import { createRule } from '/@/utils/formUtil';
  import { Button, Select, Popover } from 'ant-design-vue';

  const { t } = useI18n();
  const router = useRouter();
  const settingStore = useSettingStore();

  defineProps({
    showBack: Boolean,
  });

  const showSetting = ref(false);
  const showPeople = ref(false);

  // balance
  const balance = reactive({ loading: false, value: '' });

  // Fetching balance once wsUrl and address changed
  watchEffect(async (onInvalidate) => {
    const address = settingStore.walletInfo?.address;
    const wsUrl = settingStore.config?.wsUrl;

    if (!address || !wsUrl) return;
    balance.loading = true;

    const api = await createPolkadotApi(wsUrl);
    if (api.isConnected) {
      const { data: balanceData } = await api.query.system.account(address);

      balance.value = formatBalance(balanceData.free);
      balance.loading = false;

      onInvalidate(() => api?.disconnect());
    }
  });

  // Form
  const urlOptions = reactive([
    { label: '183.66.65.207:49944', value: 'wss://183.66.65.207:49944' },
  ]);

  const formRef = ref();
  const formData = reactive({
    wsUrl: settingStore.config?.wsUrl,
  });
  const formRules = computed(() => ({
    wsUrl: [createRule(t('home.wsUrlPlaceholder'))],
  }));

  const handleWsUrlSave = () => {
    settingStore.saveWsUrlAction(formData.wsUrl);
    showSetting.value = false;
  };

  const onBack = async () => {
    router.push({ path: '/applications/index' });
  };
</script>
<style lang="less" scoped>
  :deep(.humster-svg-icon){
    outline: none;
  }
  .title-div {
    @apply font-bold;
    margin-bottom: 10px;
  }

  .button-div {
    @apply text-center;
    margin-top: 30px;
    margin-bottom: 30px;
  }

  :deep(.ant-select-single:not(.ant-select-customize-input) .ant-select-selector) {
    height: auto !important;
  }

  :deep(.ant-select-single .ant-select-selector .ant-select-selection-item) {
    line-height: 20px !important;
    padding-top: 5px;
    padding-bottom: 5px;
    white-space: normal;
    word-break: break-all;
  }
</style>
