<template>
  <div class="text-center text-[30px] font-bold mb-[40px]">
    {{ t('applications.deploy.stepTitle3') }}
  </div>
  <div class="bg-white rounded-[20px] mx-[40px] mb-[100px] py-[50px] px-[90px]">
    <div v-if="queueInfo.length">
      <Timeline>
        <TimelineItem
          :class="getTimelineItemClass(item.status)"
          :label="DictCodeEnum.ApplicationQueueStatus.getOptionLabel(item.status || 0)"
          v-for="item in queueInfo"
          :key="item.name"
        >
          <template #dot>
            <div
              class="ant-timeline-item-icon !bg-[#63A0FA]"
              v-if="DictCodeEnum.ApplicationQueueStatus_Succeeded.is(item.status)"
            >
              <SvgIcon class="text-white rounded-[50%]" size="20" name="yes" />
            </div>

            <div
              class="ant-timeline-item-icon !bg-[#63A0FA]"
              v-else-if="DictCodeEnum.ApplicationQueueStatus_Running.is(item.status)"
            >
              <SvgIcon class="text-white rounded-[50%]" size="20" name="yes" />
            </div>

            <div
              class="!bg-[#E70000] ant-timeline-item-icon"
              v-else-if="DictCodeEnum.ApplicationQueueStatus_Failed.is(item.status)"
            >
              <SvgIcon class="text-white rounded-[50%]" size="20" name="no" />
            </div>

            <div class="ant-timeline-item-icon" v-else>
              <div class="deployment-status-not-start"></div>
            </div>
          </template>

          <svg
            class="svg"
            width="200"
            height="200"
            v-if="DictCodeEnum.ApplicationQueueStatus_Running.is(item.status)"
          >
            <circle cx="100" cy="100" r="80" />
          </svg>

          <div
            class="ant-timeline-fail-realod cursor-pointer"
            @click="handleQueueFailed"
            v-if="DictCodeEnum.ApplicationQueueStatus_Failed.is(item.status)"
          >
            <SvgIcon class="text-[#E70000]" size="40" name="reload" />
          </div>
          <div class="ant-timeline-item-title">{{ item.name }}</div>
        </TimelineItem>
      </Timeline>
      <div class="text-center">
        <router-link to="/applications">
          <Button size="large" class="w-32 mt-6 ml-4" type="primary" shape="round">
            {{ t('common.doneText') }}
          </Button>
        </router-link>
        <Button
          size="large"
          class="w-32 mt-6 ml-4"
          type="danger"
          @click="handlerDelete"
          shape="round"
        >
          {{ t('common.delText') }}
        </Button>
      </div>
    </div>
    <div class="text-center text-xl font-bold p-6" v-else>{{ t('common.loadingText') }}</div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, watchEffect, onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { GetQueueInfo } from '/@wails/go/app/Queue';
  import { DictCodeEnum } from '/@/enums/dictCodeEnum';
  import { Timeline, TimelineItem, Button } from 'ant-design-vue';
  import { RefreshGraphDeployJob, DeleteApplication } from '/@wails/go/app/Application';
  import { useRouter } from 'vue-router';
  import { useMessage } from '/@/hooks/web/useMessage';

  const props = defineProps({
    applicationId: Number,
  });

  const { t } = useI18n();
  const router = useRouter();
  const { createConfirm } = useMessage();

  const queueInfo = ref([]);
  const fetchQueueInfo = async () => {
    const { info } = await GetQueueInfo(props.applicationId);
    queueInfo.value = info;
  };

  const interval = 5000;
  watchEffect((onInvalidate) => {
    const timer = setInterval(fetchQueueInfo, interval);
    onInvalidate(() => clearInterval(timer));
  });

  const getTimelineItemClass = (status?: number) => {
    return {
      [DictCodeEnum.ApplicationQueueStatus_Succeeded.value]: 'ant-timeline-suc',
      [DictCodeEnum.ApplicationQueueStatus_Running.value]: 'ant-timeline-run',
      [DictCodeEnum.ApplicationQueueStatus_Failed.value]: 'ant-timeline-fail',
    }[status];
  };

  const handleQueueFailed = async () => {
    const data = await RefreshGraphDeployJob(props.applicationId);
    console.info(data);
  };

  const handlerDelete = async () => {
    createConfirm({
      title: 'Confirm',
      content: 'Are you sure to delete this service ?',
      onOk: async () => {
        const data = await DeleteApplication(props.applicationId);
        if (data) {
          await router.push({
            path: '/applications',
          });
        }
      },
    });
  };

  // Load info when entering page
  onMounted(fetchQueueInfo);
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
