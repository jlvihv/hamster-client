<template>
  <Form>
    <div class="stake-box p-[20px]">
      <div class="flex">
        <div class="bg-[#D8D8D8] rounded-[50%] w-[60px] h-[60px]"></div>
        <div class="ml-[16px]">
          <div>{{ t('applications.see.stake') }}</div>
          <div class="text-[20px] font-bold">0xd5f6—1a6b79</div>
        </div>
      </div>
      <div class="grid grid-cols-2 text-center mt-[20px]">
        <div class="right-line">
          <div>
            <SvgIcon color="#63A0FA" size="16" name="grt" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.yStake') }}</label>
          </div>
          <div>
            <label class="text-[18px] font-bold">0</label>
            <label class="text-[12px]">{{ t('applications.see.grt') }}</label>
          </div>
        </div>
        <div>
          <div>
            <SvgIcon color="#63A0FA" size="16" name="balance" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.wBalance') }}</label>
          </div>
          <div>
            <label class="text-[18px] font-bold">0</label>
            <label class="text-[12px]">{{ t('applications.see.grt') }}</label>
          </div>
        </div>
      </div>
    </div>
    <div class="px-[20px] pb-[60px]">
      <div class="font-bold my-[10px]">{{ t('applications.reward.stakeAmount') }}</div>
      <div
        class="border border-[#043CC1] rounded-[8px] h-[60px] px-[10px] flex items-center justify-between"
      >
        <div>10000000000</div>
    <div class="font-bold my-[10px]">{{ t('applications.reward.stakeAmount') }}</div>
    <Input class="border !border-[#043CC1] rounded-[8px] h-[60px] px-[10px]" value="10000000">
      <template #suffix>
        <div>
          <label class="text-[#7B8082] mr-[10px]">{{ t('applications.see.grt') }}</label>
          <label class="bg-[#63A0FA] px-[20px] py-[8px] rounded-[4px] text-white">
            {{ t('applications.see.max') }}
          </label>
        </div>
      </div>
      <div class="flex justify-between mt-[10px]">
        <div>{{ t('applications.see.minAmount') }}</div>
        <div>
          <label class="text-[18px] font-bold mr-[3px]">10000000000</label
          >{{ t('applications.see.grt') }}
        </div>
      </template>
    </Input>
    <div class="flex justify-between mt-[10px]">
      <div>{{ t('applications.see.minAmount') }}</div>
      <div>
        <label class="text-[18px] font-bold mr-[3px]">10000000000</label
        >{{ t('applications.see.grt') }}
      </div>
      <div class="flex justify-between mt-[10px]">
        <div>{{ t('applications.see.thawPeriod') }}</div>
        <div>
          <label class="text-[18px] font-bold mr-[3px]">28</label>{{ t('applications.see.days') }}
        </div>
      </div>
      <div class="flex my-[10px]">
        <div class="seq-div !bg-[#043CC1]">1</div>{{ t('applications.see.gtrStak') }}
      </div>
      <Button type="primary" size="large">{{ t('applications.see.grtAccess') }}</Button>
      <div class="flex my-[10px]">
        <div class="seq-div">2</div
        ><label class="text-[#7B8082]">{{ t('applications.see.gtrStak') }}</label>
      </div>
      <Button type="primary" size="large" disabled>{{ t('applications.see.stake') }}</Button>
    </div>
  </div>
    <div class="flex my-[10px]">
      <div class="seq-div !bg-[#043CC1]">1</div>{{ t('applications.see.gtrStak') }}
    </div>
    <Button type="primary" size="large">{{ t('applications.see.grtAccess') }}</Button>
    <div class="flex my-[10px]">
      <div class="seq-div">2</div
      ><label class="text-[#7B8082]">{{ t('applications.see.gtrStak') }}</label>
    </div>
    <Button type="primary" size="large" disabled>{{ t('applications.see.stake') }}</Button>
  </Form>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { Button, Input, Form } from 'ant-design-vue';

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

  .seq-div {
    @apply text-white text-center;
    background: rgba(4, 60, 193, 0.5);
    border-radius: 50%;
    height: 20px;
    width: 20px;
    font-size: 12px;
    margin-right: 5px;
  }

  :deep(.ant-btn-primary) {
    @apply w-full;
    border-radius: 24px;
  }

  :deep(.ant-btn[disabled]) {
    @apply text-white border-none;
    background: rgba(0, 0, 0, 0.1);
  }
</style>
