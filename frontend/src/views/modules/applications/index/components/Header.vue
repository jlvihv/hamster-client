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
      <SvgIcon
        class="cursor-pointer"
        color="#858B92"
        size="20"
        name="people"
        @click="
          showPeople = true;
          showSetting = false;
        "
      />
      <SvgIcon
        class="cursor-pointer ml-[20px]"
        color="#858B92"
        size="20"
        name="setting"
        @click="
          showSetting = true;
          showPeople = false;
        "
      />
    </div>
    <div v-if="showPeople" @mouseleave="showPeople = false" class="pop-div">
      <div class="top-div right-[50px]"></div>
      <div class="border-box">
        <div class="title-div">
          <SvgIcon size="16" color="#63A0FA" name="address" />
          {{ t('applications.index.address') }}
        </div>
        <div class="text-[12px] mb-[20px] whitespace-normal break-all">
          {{ settingStore.walletInfo?.address }}
        </div>
        <div class="title-div">
          <SvgIcon size="16" color="#63A0FA" name="balance" />
          {{ t('applications.index.balance') }}
        </div>
        <div class="text-[12px] mb-[20px]">
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
    </div>
    <div v-if="showSetting" @mouseleave="onMouseLeave($event)" class="pop-div">
      <div class="top-div right-[10px]"></div>
      <Form class="border-box" ref="formRef" :model="formData" :rules="formRules">
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
  import { Button, Select } from 'ant-design-vue';

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

  async function onMouseLeave(e) {
    var currTargetEl = e.relatedTarget || e.toElement;
    const targetClassName = currTargetEl.className
    if (targetClassName.indexOf('ant-select-dropdown') !== -1) return; 
    showSetting.value = false;
  }
</script>
<style lang="less" scoped>
  .pop-div {
    @apply absolute z-50 top-9;
    width: 250px;
    right: -10px;
  }

  .border-box {
    @apply bg-white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0px 0px 4px 0px rgba(31, 31, 35, 0.2);
  }

  .top-div {
    @apply absolute;
    top: -40px;
    width: 20px;
    height: 40px;
    border-bottom: 10px solid #fff;
    border-left: 10px solid transparent;
    border-right: 10px solid transparent;
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
