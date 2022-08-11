<template>
  <div class="grid grid-cols-4 gap-8 text-center mt-3 mx-3">
    <div
      class="bg-white rounded-[8px] py-[20px] px-[40px] duration-500 hover:scale-110"
      v-for="item in subgraphs"
      :key="item.id"
    >
      <img :src="item.image" class="w-full rounded-[8px]" />
      <div class="text-[14px] text-[#BFC6D4] mt-[6px]">{{ shortenAddress(item.owner.id) }}</div>
      <div class="text-[#222222] mt-[6px] font-bold truncate">{{ item.displayName }}</div>
      <div class="text-[14px] text-[#BFC6D4] mt-[6px]">
        Signal: {{ formatfromWei(item.currentSignalledTokens) }} GRT
      </div>
      <div
        class="text-[#63A0FA] text-[14px] mt-[20px] border border-[#63A0FA] rounded-[4px] h-[30px]"
      >
        {{ t('applications.see.stop') }}
      </div>
    </div>
  </div>
  <div class="text-center my-[40px]" v-if="isTouchedEnd">
    <Button
      class="!h-[60px] w-[200px]"
      size="large"
      type="primary"
      @click="loadSubgraphs"
      :loading="isSubgraphsLoading"
    >
      {{ t('common.moreText') }}
    </Button>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useLoadMore } from '/@/hooks/web/useLoadMore';
  import { fetchSubgraphs } from '/@/utils/thegraphUtil/subgraph';
  import { shortenAddress } from '/@/utils/thegraphUtil';
  import { formatfromWei } from '/@/utils/web3Util';
  import { Button } from 'ant-design-vue';

  const { t } = useI18n();
  const {
    isTouchedEnd,
    items: subgraphs,
    isLoading: isSubgraphsLoading,
    loadMore: loadSubgraphs,
  } = useLoadMore(fetchSubgraphs, {
    responseHandler: (data) => data.data.subgraphs,
    perPage: 20,
  });

  onMounted(loadSubgraphs);
</script>
