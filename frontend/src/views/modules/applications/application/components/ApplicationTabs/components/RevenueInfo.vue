<template>
  <div class="revenue-bg rounded-[20px] p-[20px] relative">
    <div
      @click="receiveBenefits"
      class="benefit-bg cursor-pointer absolute right-0 top-10 text-white py-[16px] pl-[16px] rounded-l-[100px]"
    >
      <label class="text-[18px] font-bold">{{ t('applications.see.receiveBenefits') }}</label>
      <SvgIcon class="text-white rounded-[50%] mx-[5px]" size="16" name="right" />
    </div>
    <div class="flex">
      <div class="bg-[#D8D8D8] rounded-[50%] h-[120px] w-[120px]"></div>
      <div class="text-[#F4F4F4] ml-[20px]">
        <div class="text-[26px]">0xd5f6—1a6b79</div>
        <div>0xd5f6e31199220a0d5334cad2b6ecd70c8f1a6b79</div>
        <div class="relative text-white"
          ><label class="text-[26px] font-bold">88999999</label
          ><label class="absolute top-0">GRT</label></div
        >
      </div>
    </div>
    <div class="grid grid-cols-3 text-white text-center mt-[20px]">
      <div class="right-line">
        <div class="text-[22px] font-bold">800000000GRT</div>
        <div>{{ t('applications.see.income') }}</div>
      </div>
      <div class="right-line">
        <div class="text-[22px] font-bold">800000000GRT</div>
        <div>{{ t('applications.see.stakAmount') }}</div>
      </div>
      <div>
        <div class="text-[22px] font-bold">800000000GRT</div>
        <div>{{ t('applications.see.unstakAmount') }}</div>
      </div>
    </div>
    <div class="my-[40px] text-center">
      <label
        class="label-btn"
        @click="
          drawerVisible = true;
          stakeVisible = true;
        "
        >{{ t('applications.see.stake') }}</label
      >
      <label
        class="label-btn ml-[30px]"
        @click="
          drawerVisible = true;
          unstakeVisible = true;
        "
        >{{ t('applications.see.unstake') }}</label
      >
      <label
        class="label-btn ml-[30px]"
        @click="
          drawerVisible = true;
          withdrawVisible = true;
        "
        >{{ t('applications.see.withdraw') }}</label
      >
    </div>
  </div>
  <Drawer
    v-model:visible="drawerVisible"
    :closable="false"
    placement="right"
    wrapClassName="drawer-revenue-info"
    @close="onDrawerClose"
  >
    <StakeDrawer v-if="stakeVisible" />
    <UnstakeDrawer v-if="unstakeVisible" />
    <WithdrawDrawer v-if="withdrawVisible" />
  </Drawer>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import StakeDrawer from './StakeDrawer.vue';
  import UnstakeDrawer from './UnstakeDrawer.vue';
  import WithdrawDrawer from './WithdrawDrawer.vue';
  import { Drawer } from 'ant-design-vue';

  const { t } = useI18n();
  const emits = defineEmits(['modalConfirm']);

  const drawerVisible = ref(false);
  const stakeVisible = ref(false);
  const unstakeVisible = ref(false);
  const withdrawVisible = ref(false);

  async function onDrawerClose() {
    stakeVisible.value = false;
    unstakeVisible.value = false;
    withdrawVisible.value = false;
  }

  const receiveBenefits = () => {
    emits('modalConfirm');
  };
</script>
<style lang="less" scoped>
  :deep(.ant-drawer-body) {
    padding: 0px;
  }

  .drawer-revenue-info {
    :global(.ant-drawer) {
      display: flex;
      align-items: center;
    }

    :global(.ant-drawer-right .ant-drawer-content-wrapper) {
      height: auto;
    }

    :global(.ant-drawer-content) {
      border-radius: 0 8px 8px 0;
      min-height: 500px;
    }

    :global(.ant-drawer-body) {
      padding: 0px;
    }
  }

  .revenue-bg {
    background: url('src/assets/images/application-bg.png') no-repeat top;
    background-size: 100% 100%;
  }

  .label-btn {
    @apply bg-[#376AED] px-[30px] py-[10px] rounded-[20px] text-[#E2EAFC] cursor-pointer;
  }

  .benefit-bg {
    background: linear-gradient(122deg, #ff893e 0%, #e76a93 100%);
  }
</style>
