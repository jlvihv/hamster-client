<template>
  <div class="text-center text-[30px] font-bold mb-[40px]">
    {{ t('applications.deploy.stepTitle3') }}
  </div>
  <div class="bg-white rounded-[20px] mx-[40px] mb-[100px] py-[50px] px-[90px]">
    <Timeline>
      <TimelineItem class="ant-timeline-suc" label="Complete">
        <template #dot>
          <div class="ant-timeline-item-icon !bg-[#63A0FA]">
            <SvgIcon class="text-white rounded-[50%]" size="20" name="yes" />
          </div>
        </template>
        <div class="ant-timeline-item-title">{{ t('applications.see.staking') }}</div>
      </TimelineItem>
      <TimelineItem class="ant-timeline-run" label="Runing">
        <template #dot>
          <div class="ant-timeline-item-icon !bg-[#63A0FA]">
            <SvgIcon class="text-white rounded-[50%]" size="20" name="yes" />
          </div>
        </template>
        <svg class="svg" width="200" height="200">
          <circle cx="100" cy="100" r="80" />
        </svg>
        <div class="ant-timeline-item-title">{{ t('applications.see.waiting') }}</div>
      </TimelineItem>
      <TimelineItem class="ant-timeline-fail" label="Runing">
        <template #dot>
          <div class="!bg-[#E70000] ant-timeline-item-icon">
            <SvgIcon class="text-white rounded-[50%]" size="20" name="no" />
          </div>
        </template>
        <div class="ant-timeline-item-title">{{ t('applications.see.servicePull') }}</div>
        <div class="ant-timeline-fail-realod">
          <SvgIcon class="text-[#E70000]" size="40" name="reload" />
        </div>
      </TimelineItem>
      <TimelineItem label="Runing">
        <template #dot>
          <div class="ant-timeline-item-icon">
            <div class="deployment-status-not-start"></div>
          </div>
        </template>
        <div class="ant-timeline-item-title">{{ t('applications.see.serviceDeploy') }}</div>
      </TimelineItem>
    </Timeline>
    <div class="text-center">
      <router-link to="/applications">
        <Button size="large" class="w-32 mt-6 ml-4" type="primary">
          {{ t('common.doneText') }}
        </Button>
      </router-link>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { Timeline, TimelineItem, Button } from 'ant-design-vue';

  const { t } = useI18n();
</script>

<style lang="less" scoped>
  .deployment-status-not-start {
    @apply w-full h-full rounded-full bg-white border-2 border-[#63A0FA];
  }

  :deep(.ant-timeline) {
    .svg {
      margin: 0;
    }
  }
</style>
