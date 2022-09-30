<template>
  <Header :showBack="true" class="px-4 pt-4" />
  <PageWrapper class="w-[1044px] mx-auto">
    <div class="flex justify-center">
      <img :src="selectBlockchain" class="h-6 mt-2" />
      <span class="ml-2 text-3xl font-bold"> {{ t('select.selectBlockchain') }} </span>
    </div>
    <div class="flex mt-8">
      <router-link to="/applications/new?blockchain=thegraph" class="block-container">
        <SvgIcon size="32" name="select-graph" />
        <span>{{ t('select.graph') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=ethereum" class="block-container">
        <SvgIcon size="32" name="select-ethereum" />
        <span>{{ t('select.ethereum') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=bsc" class="block-container">
        <SvgIcon size="32" name="select-bsc" />
        <span>{{ t('select.bsc') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=polygon" class="block-container">
        <SvgIcon size="32" name="select-polygon" />
        <span>{{ t('select.polygon') }}</span>
      </router-link>
    </div>
    <div class="flex mt-6">
      <router-link to="/applications/new?blockchain=avalanche" class="block-container">
        <SvgIcon size="32" name="select-avalanche" />
        <span>{{ t('select.avalanche') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=optimism" class="block-container">
        <SvgIcon size="32" name="select-optimism" />
        <span>{{ t('select.optimism') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=zksync" class="block-container">
        <SvgIcon size="32" name="select-zksync" />
        <span>{{ t('select.zkSync') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=starkware" class="block-container">
        <SvgIcon size="32" name="select-starkware" />
        <span>{{ t('select.starkWare') }}</span>
      </router-link>
    </div>
    <div class="flex mt-6">
      <router-link to="/applications/new?blockchain=near" class="block-container">
        <SvgIcon size="32" name="select-near" />
        <span>{{ t('select.near') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=aptos" class="block-container">
        <SvgIcon size="32" name="select-aptos" />
        <span>{{ t('select.aptos') }}</span>
      </router-link>
      <router-link to="/applications/new?blockchain=sui" class="block-container">
        <SvgIcon size="32" name="select-sui" />
        <span>{{ t('select.sui') }}</span>
      </router-link>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import Header from '../index/components/Header.vue';
  import { SvgIcon } from '/@/components/Icon';
  import selectBlockchain from '/@/assets/images/select-blockchain.png';

  const { t } = useI18n();
</script>

<style lang="less" scoped>
  .block-container {
    @apply h-24 w-52 bg-[#fff] rounded-lg flex pl-8 ml-6;

    & > svg {
      @apply self-center cursor-pointer;
    }

    & > span {
      @apply self-center text-base font-bold ml-3 text-black;
    }
  }
</style>
