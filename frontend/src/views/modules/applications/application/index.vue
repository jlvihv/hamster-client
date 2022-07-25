<template>
  <PageWrapper>
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
        <DescriptionsItem :label="t('applications.index.desText')">
          {{ appInfo.describe }}
        </DescriptionsItem>
      </Descriptions>
    </Card>
    <Reward :deployInfo="deployInfo" v-if="isAppDeployed" />
    <div class="h-4"></div>
    <DeployInfo :deployInfo="deployInfo" v-if="isAppDeployed" />
    <div class="mt-4 text-right">
      <Button type="primary" @click="onClose">{{ t('common.closeText') }}</Button>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { onMounted, ref, reactive, computed } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { DictCodeEnum } from '/@/enums/dictCodeEnum';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import DeployInfo from './components/DeployInfo.vue';
  import { QueryApplicationById } from '/@wails/go/app/Application';
  import { GetDeployInfo } from '/@wails/go/app/Deploy';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Descriptions, DescriptionsItem, Button } from 'ant-design-vue';
  import Reward from '/@/views/modules/applications/application/components/Reward.vue';

  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const router = useRouter();
  const { params } = useRoute();
  const applicationId = Number(params.id);

  const appInfo = reactive({
    name: '',
    createdAt: '',
    status: 0,
    describe: '',
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
