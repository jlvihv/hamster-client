<template>
  <div class="relative h-full" v-if="isGuideVisible">
    <Form
      class="home-bg h-full label-center text-center"
      :style="{ backgroundImage: `url(${getImageURL('home-bg.png')})` }"
      ref="formRef"
      :model="formData"
      :rules="formRules"
    >
      <transition-group name="next" appear mode="in-out">
        <div v-if="stepVal === 0">
          <div class="title-text">
            {{ t('home.setNode') }}
          </div>
          <div class="my-[40px] !text-left">
            <FormItem name="wsUrl">
              <Select
                class="rounded-[8px] bg-[#F8F7FA] w-[560px]"
                :options="nodeOptions"
                v-model:value="formData.wsUrl"
                :placeholder="t('home.wsUrlPlaceholder')"
              />
            </FormItem>
          </div>
          <SvgButton @click="stepAction.setWsUrl" iconClass="text-primary" size="56" icon="next" />
        </div>
        <div v-else-if="stepVal === 1">
          <div
            class="absolute left-[20px] top-[20px] cursor-pointer"
            v-if="hasBackButton"
            @click="goBack"
          >
            <SvgIcon name="left" color="#858B92" size="20" />
          </div>
          <WalletImporter @submit="stepAction.next" />
        </div>
        <div v-else-if="stepVal === 2">
          <div class="label-center">
            <img :src="doneImage" class="w-[200px]" />
          </div>
          <div class="title-text my-[40px]">{{ t('home.complete') }}</div>
          <SvgButton
            @click="stepAction.gotoApplicationsPage"
            iconClass="text-primary"
            size="56"
            icon="next"
          />
        </div>
      </transition-group>
    </Form>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, computed, reactive, ref } from 'vue';
  import { useRouter, useRoute } from 'vue-router';
  import { useSettingStore } from '/@/store/modules/setting';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useAssets } from '/@/hooks/web/useAssets';
  import { SvgIcon } from '/@/components/Icon';
  import { SvgButton } from '/@/components/SvgButton';
  import { createRule } from '/@/utils/formUtil';
  import WalletImporter from './components/WalletImporter.vue';
  import doneImage from '/@/assets/images/suc.png';
  import { Form, FormItem, Select } from 'ant-design-vue';

  const { t } = useI18n();
  const { getImageURL } = useAssets();
  const router = useRouter();
  const route = useRoute();
  const settingStore = useSettingStore();

  const nodeOptions = reactive([
    { label: '183.66.65.207:49944', value: 'ws://183.66.65.207:49944' },
  ]);
  const stepVal = ref(0);
  const hasBackButton = ref(false);

  // Form
  const formRef = ref();
  const formData = reactive({
    wsUrl: settingStore.config?.wsUrl,
  });
  const formRules = computed(() => ({
    wsUrl: [createRule(t('home.wsUrlPlaceholder'))],
  }));

  // Loading settingStore and check if should show guide
  const isGuideVisible = ref(false);
  onMounted(async () => {
    // fetch settings from API
    try {
      await settingStore.getWalletInfoAction();
    } catch {
      console.log('Failed to load wallet and config.');
    }

    if (route.query.step) {
      isGuideVisible.value = true;

      stepVal.value = parseInt(route.query.step);
      hasBackButton.value = true;
    } else {
      if (settingStore.walletInfo?.addressJson) {
        // Redirect to applidation index if wallet binded
        router.push('/applications/index');
      } else {
        isGuideVisible.value = true;
      }
    }
  });

  const stepAction = {
    next() {
      stepVal.value++;
    },
    async setWsUrl(callback) {
      await formRef.value?.validate('wsUrl');

      settingStore.saveWsUrlAction(formData.wsUrl);
      stepAction.next();
      callback();
    },
    gotoApplicationsPage() {
      // Add from=guide when setup account first time
      const query = route.query.step ? { from: 'guide' } : {};
      router.push({
        path: '/applications/index',
        query,
      });
    },
  };

  const goBack = () => {
    router.go(-1);
  };
</script>

<style lang="less" scoped>
  .home-bg {
    background: no-repeat top;
    background-size: 100% 100%;
  }

  .title-text {
    @apply text-[#222222] text-[40px] font-bold;
  }

  .label-center {
    @apply flex items-center justify-center;
  }

  :deep(.ant-select-selector) {
    @apply !rounded-[8px] !h-[42px];

    .ant-select-selection-placeholder {
      line-height: 42px;
    }
  }

  :deep(.ant-select-selection-item) {
    @apply !leading-[40px];
  }

  .next-enter-active {
    animation: show 0.5s linear;
  }

  .next-leave-active {
    animation: hide 0.4s linear reverse;
  }

  .next-enter,
  .next-leave-to {
    @apply opacity-0;
  }
  @keyframes show {
    0% {
      @apply opacity-0;
      transform: translateX(100%);
    }

    100% {
      @apply opacity-100;
      transform: translateX(0%);
    }
  }
  @keyframes hide {
    0% {
      @apply opacity-0;
      transform: translateX(-100%);
    }

    100% {
      @apply opacity-100;
      transform: translateX(0%);
    }
  }
</style>
