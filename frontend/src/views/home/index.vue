<template>
  <div class="home-bg h-screen label-center text-center">
    <transition-group name="next" appear mode="in-out">
      <div v-if="stepVal === 0">
        <div class="title-text">
          {{ t('home.setNode') }}
        </div>
        <div class="my-[40px] !text-left">
          <Select class="rounded-[8px] bg-[#F8F7FA] w-[560px]" :options="nodeOptions" />
        </div>
        <SvgIcon @click="nextStep(1)" class="text-primary cursor-pointer" size="56" name="next" />
      </div>
      <div v-else-if="stepVal === 1">
        <div class="title-text">{{ t('home.importWallet') }}</div>
        <div class="label-center my-[40px]">
          <img src="src/assets/images/wallet.png" class="cursor-pointer w-[200px]" />
        </div>
        <SvgIcon
          @click="nextStep(2)"
          class="text-primary cursor-pointer"
          size="56"
          name="importWallet"
        />
      </div>
      <div v-else-if="stepVal === 2">
        <ImportWalletModal />
        <SvgIcon @click="nextStep(3)" class="text-primary cursor-pointer" size="56" name="next" />
      </div>
      <div v-else-if="stepVal === 3">
        <div class="label-center">
          <img src="src/assets/images/suc.png" class="cursor-pointer w-[200px]" />
        </div>
        <div class="title-text my-[40px]">{{ t('home.complete') }}</div>
        <SvgIcon @click="complete" class="text-primary cursor-pointer" size="56" name="next" />
      </div>
    </transition-group>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, reactive, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useSettingStore } from '/@/store/modules/setting';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import ImportWalletModal from './components/ImportWalletModal.vue';
  import { Select } from 'ant-design-vue';

  const { t } = useI18n();
  const router = useRouter();
  const settingStore = useSettingStore();

  const nodeOptions = reactive([
    { label: '183.66.65.207:49944', value: '183.66.65.207:49944moon' },
  ]);
  const stepVal = ref(0);

  onMounted(() => {
    // Redirect to applidation index if wallet binded
    if (settingStore.walletInfo) {
      router.push('/applications/index');
    }
  });

  async function nextStep(step) {
    stepVal.value = step;
  }
  async function complete() {
    router.push('/applications/index');
  }
</script>
<style lang="less" scoped>
  .home-bg {
    background: url('src/assets/images/home-bg.png') no-repeat top;
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
  }

  :deep(.ant-select-selection-item) {
    @apply !leading-[40px];
  }

  .next-enter-active {
    animation: show 0.5s linear;
  }

  .next-leave-active {
    animation: hide 0.5s linear reverse;
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
