<template>
  <PageWrapper>
    <div v-if="true">
      <Header :showBack="true"/>
      <div class="text-center text-[30px] font-bold mb-[40px]">{{ t("applications.deploy.stepTitle3") }}</div>
      <div class="bg-white rounded-[20px] mx-[40px] mb-[100px] py-[40px] px-[90px]">
        <Timeline>
          <TimelineItem class="ant-timeline-suc" label="Complete">
            <template #dot>
              <div class="ant-timeline-item-icon !bg-[#63A0FA]">
                <SvgIcon class="text-white rounded-[50%]" size="20" name="yes" />
              </div> 
            </template>
            <div class="ant-timeline-item-title">{{ t("applications.see.staking") }}</div>
          </TimelineItem>
          <TimelineItem class="ant-timeline-run" label="Runing">
            <template #dot>
              <div class="ant-timeline-item-icon !bg-[#63A0FA]">
                <SvgIcon class="text-white rounded-[50%]" size="20" name="yes" />
              </div>  
            </template>
            <svg class="svg" width=200 height=200>
                <circle cx=100 cy=100 r=80 />
            </svg>
            <div class="ant-timeline-item-title">{{ t("applications.see.waiting") }}</div>
          </TimelineItem>
          <TimelineItem class="ant-timeline-fail" label="Runing">
            <template #dot>
              <div class="!bg-[#E70000] ant-timeline-item-icon">
                <SvgIcon class="text-white rounded-[50%]" size="20" name="no" />
              </div>  
            </template>
            <div class="ant-timeline-item-title">{{ t("applications.see.servicePull") }}</div>
            <div class="ant-timeline-fail-realod">
              <SvgIcon class="text-[#E70000]" size="40" name="reload" />
            </div>
          </TimelineItem>
          <TimelineItem label="Runing">
            <template #dot>
              <div class="ant-timeline-item-icon"></div>  
            </template>
            <div class="ant-timeline-item-title">{{ t("applications.see.serviceDeploy") }}</div>
          </TimelineItem>
        </Timeline>
        <div class="text-center">
          <Button
            size="large"
            class="w-32 mt-6 ml-4 !rounded-[30px]"
            type="primary"
          >
            {{ t('common.doneText') }}
          </Button>
        </div>
      </div>
    </div>
    <div v-else>
      <Card class="!mb-4" :loading="!appInfo.id">
        <Descriptions :title="t('applications.see.appInfo')" bordered>
          <template #extra>
            <router-link v-if="isAppDeployed" :to="`/applications/${applicationId}/cli`"
              >Cli</router-link
            >
            <router-link v-if="!isAppDeployed" :to="`/applications/${applicationId}/deploy`"
              >Deploy</router-link
            >
          </template>
          <DescriptionsItem :label="t('applications.index.nameText')">
            {{ appInfo.name }}
          </DescriptionsItem>
          <DescriptionsItem :label="t('applications.index.addTimeText')">
            {{ formatToDateTime(appInfo.createdAt) }}
          </DescriptionsItem>
          <DescriptionsItem :label="DictCodeEnum.ApplicationDeployStatus.getLabel()">
            {{ DictCodeEnum.ApplicationDeployStatus.getOptionLabel(appInfo.status) }}
          </DescriptionsItem>
          <DescriptionsItem :label="t('applications.index.plugText')">
            {{ appInfo.plugin }}
          </DescriptionsItem>
        </Descriptions>
      </Card>
      <Reward :deployInfo="deployInfo" v-if="isAppDeployed" />
      <div class="h-4"></div>
      <DeployInfo :deployInfo="deployInfo" v-if="isAppDeployed" />
      <div class="mt-4 text-right">
        <Button type="primary" @click="onClose">{{ t('common.closeText') }}</Button>
      </div>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { onMounted, ref, reactive, computed } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { ClockCircleOutlined } from '@ant-design/icons-vue';
  import { DictCodeEnum } from '/@/enums/dictCodeEnum';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import DeployInfo from './components/DeployInfo.vue';
  import Header from '../index/components/Header.vue';
  import { SvgIcon } from '/@/components/Icon';
  import { QueryApplicationById } from '/@wails/go/app/Application';
  import { GetDeployInfo } from '/@wails/go/app/Deploy';
  import { useI18n } from '/@/hooks/web/useI18n';
  import Reward from '/@/views/modules/applications/application/components/Reward.vue';
  import { Card, Descriptions, DescriptionsItem, Button, Timeline, TimelineItem } from 'ant-design-vue';

  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const router = useRouter();
  const { params } = useRoute();
  const applicationId = Number(params.id);

  const appInfo = reactive({
    name: '',
    createdAt: '',
    status: 0,
    plugin: '',
  });
  const deployInfo = ref<{
    initialization: Recordable;
    staking: Recordable;
    deployment: Recordable;
  }>({
    initialization: {},
    staking: {},
    deployment: {},
  });

  // 0: Not deployed, 1: Deployed
  const isAppDeployed = computed(() => appInfo.status === 1);

  const getAppInfo = async () => {
    try {
      const result = await QueryApplicationById(applicationId);
      Object.assign(appInfo, result);
    } catch (error: any) {
      createErrorModal({
        title: t('common.errorTip'),
        content: t('common.operateFailText'),
      });
    }
  };

  // Get saved deployInfo from API
  const getDeployInfo = async () => {
    const data = await GetDeployInfo(applicationId);
    if (data) deployInfo.value = data;
  };

  onMounted(async () => {
    getAppInfo();
    getDeployInfo();
  });

  async function onClose() {
    router.push('/applications/index');
  }
</script>
