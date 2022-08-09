<template>
  <div class="grid grid-cols-4 gap-8 text-center mt-3 mx-3">
    <div 
      class="bg-white rounded-[8px] py-[20px] px-[40px] duration-500 hover:scale-110"
      v-for="(item, index) in [1, 2, 3, 4, 5]"
      :key="index"
    >
      <img
        :src="`src/assets/images/subgraph-img.png`"
        class="w-full rounded-[8px]"
      />
      <div class="text-[14px] text-[#BFC6D4] mt-[6px]">0x5ca1—72b3f9</div>
      <div class="text-[#222222] mt-[6px] font-bold">Premia-Mainnet{{ item }}</div>
      <div class="text-[14px] text-[#BFC6D4] mt-[6px]">Signal: 209000 GRT</div>
      <div class="text-[#63A0FA] text-[14px] mt-[20px] border border-[#63A0FA] rounded-[4px] h-[30px]">{{ t('applications.see.stop') }}</div>
    </div>
  </div>
  <div class="text-center my-[40px]">
    <Button class="!h-[60px] w-[200px]" size="large" type="primary">
      <svg :class="[btnLoading===false?'hidden' : 'inline']" class="animate-spin align-text-bottom -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      {{ t('common.moreText') }}
    </Button>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Button } from 'ant-design-vue';

  const { t } = useI18n();
  const btnLoading = ref(true);
</script>
