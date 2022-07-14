<template>
  <PageWrapper>
    <Card class="!mb-4">
      <Descriptions :title="t('applications.see.appInfo')" bordered>
        <template #extra v-if="!isAppDeployed">
          <router-link :to="`/applications/${applicationId}/deploy`">Deploy</router-link>
        </template>
        <DescriptionsItem :label="t('applications.index.nameText')">
          {{ appInfo.name }}
        </DescriptionsItem>
        <DescriptionsItem :label="t('applications.index.addTimeText')">
          {{ formatToDateTime(appInfo.createdAt) }}
        </DescriptionsItem>
        <DescriptionsItem :label="t('applications.index.statusText')">
          {{ statusOptions.find((option) => option.value === appInfo.status)?.label }}
        </DescriptionsItem>
        <DescriptionsItem :label="t('applications.index.desText')">
          {{ appInfo.describe }}
        </DescriptionsItem>
      </Descriptions>
    </Card>
    <DeployInfo :deployInfo="deployInfo" v-if="isAppDeployed" />
    <div class="mt-4 text-right">
      <Button type="primary" @click="onClose">{{ t('common.closeText') }}</Button>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { onMounted, reactive, computed } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { useOptionsContent } from '.././useApp';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import DeployInfo from './components/DeployInfo.vue';
  import { QueryApplicationById } from '/@wails/go/app/Application';
  import { GetDeployInfo } from '/@wails/go/app/Deploy';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Descriptions, DescriptionsItem, Button } from 'ant-design-vue';

  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const { statusOptions } = useOptionsContent();
  const router = useRouter();
  const { params } = useRoute();
  const applicationId = Number(params.id);

  const appInfo = reactive({
    name: '',
    createdAt: '',
    status: 0,
    describe: '',
  });
  const deployInfo = reactive<{
    initialization: Recordable;
    staking: Recordable;
    deployment: Recordable;
  }>({ initialization: {}, staking: {}, deployment: {} });

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
    const { data } = await GetDeployInfo(applicationId);
    data && Object.assign(deployInfo, data);
  };

  onMounted(async () => {
    getAppInfo();
    getDeployInfo();
  });

  async function onClose() {
    router.push('/applications/index');
  }
</script>
