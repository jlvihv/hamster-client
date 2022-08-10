<template>
  <div>
    <div class="stake-box p-[20px]">
      <div class="flex">
        <div class="bg-[#D8D8D8] rounded-[50%] w-[60px] h-[60px]"></div>
        <div class="ml-[16px]">
          <div>{{ t('applications.see.withdraw') }}</div>
          <div class="text-[20px] font-bold">0xd5f6—1a6b79</div>
        </div>
      </div>
      <div class="grid grid-cols-2 text-center mt-[20px]">
        <div class="right-line">
          <div>
            <SvgIcon color="#63A0FA" size="16" name="with" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.withdraw') }}</label>
          </div>
          <div>
            <label class="text-[18px] font-bold">0</label>
            <label class="text-[12px]">{{ t('applications.see.grt') }}</label>
          </div>
        </div>
        <div>
          <div>
            <SvgIcon color="#63A0FA" size="16" name="time" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.thawLeft') }}</label>
          </div>
          <div class="text-[18px] font-bold">1D 0H 19M</div>
        </div>
      </div>
    </div>
    <div class="px-[20px] pb-[60px]">
      <Button class="mt-[20px]" type="primary" size="large">
        {{ t('applications.see.withdraw') }}
      </Button>
    </div>
  </div>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { Button } from 'ant-design-vue';

  const { t } = useI18n();
</script>
<style lang="less" scoped>
  .stake-box {
    box-shadow: 0px 0px 5px 0px rgba(150, 150, 150, 0.2);
  }

  .right-line::after {
    border-color: rgba(216, 216, 216, 0.4);
    top: 30%;
  }

  :deep(.ant-btn-primary) {
    @apply w-full;
    border-radius: 24px;
  }
</style>
