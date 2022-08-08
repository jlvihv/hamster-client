<template>
  <div class="relative h-full">
    <form
      class="home-bg h-full label-center text-center"
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
            <Select
              class="rounded-[8px] bg-[#F8F7FA] w-[560px]"
              :options="nodeOptions"
              v-model:value="formData.wsUrl"
            />
          </div>
          <SvgIcon
            @click="stepAction.setWsUrl"
            class="text-primary cursor-pointer"
            size="56"
            name="next"
          />
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
          <SvgIcon
            @click="stepAction.gotoApplicationsPage"
            class="text-primary cursor-pointer"
            size="56"
            name="next"
          />
        </div>
      </transition-group>
    </form>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, computed, reactive, ref } from 'vue';
  import { useRouter, useRoute } from 'vue-router';
  import { useSettingStore } from '/@/store/modules/setting';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { createRule } from '/@/utils/formUtil';
  import WalletImporter from './components/WalletImporter.vue';
  import doneImage from '/@/assets/images/suc.png';
  import { Select } from 'ant-design-vue';

  const { t } = useI18n();
  const router = useRouter();
  const route = useRoute();
  const settingStore = useSettingStore();

  const nodeOptions = reactive([
    { label: '183.66.65.207:49944', value: '183.66.65.207:49944moon' },
  ]);
  const stepVal = ref(0);
  const hasBackButton = ref(false);

  // Form
  const formRef = ref();
  const formData = reactive({});
  const formRules = computed(() => ({
    wsUrl: [createRule(t('home.wsUrlPlaceholder'))],
  }));

  onMounted(() => {
    if (route.query.step) {
      stepVal.value = parseInt(route.query.step);
      hasBackButton.value = true;
    } else {
      // Redirect to applidation index if wallet binded
      if (settingStore.walletInfo) {
        router.push('/applications/index');
      }
    }
  });

  const stepAction = {
    next() {
      stepVal.value++;
    },
    setWsUrl() {
      settingStore.saveWsUrlAction(formData.wsUrl);
      stepAction.next();
    },
    gotoApplicationsPage() {
      router.push('/applications/index');
    },
  };

  const goBack = () => {
    router.go(-1);
  };
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
