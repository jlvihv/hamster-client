<template>
  <PageWrapper>
    <Card>
      <Descriptions :title="t('applications.see.appInfo')" bordered>
        <template #extra v-if="true">
          <router-link :to="`/applications/${applicationId}/deploy`">Deploy</router-link>
        </template>
        <DescriptionsItem :label="t('applications.index.nameText')">{{
          appInfo.name
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('applications.index.addTimeText')">{{
          appInfo.createdAt
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('applications.index.statusText')">{{
          appInfo.status
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('applications.index.desText')">{{
          appInfo.describe
        }}</DescriptionsItem>
      </Descriptions>
    </Card>
    <div class="h-4"></div>
    <DeployInfo />
    <div class="mt-4 text-right">
      <Button type="primary" @click="onClose">{{ t('common.closeText') }}</Button>
    </div>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { onMounted, reactive } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import DeployInfo from './components/DeployInfo.vue';
  import { QueryApplicationById } from '/@wails/go/app/Application';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Descriptions, DescriptionsItem, Button } from 'ant-design-vue';

  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const router = useRouter();
  const { params } = useRoute();
  const { id: applicationId } = params;
  const appInfo = reactive({
    name: '',
    createdAt: '',
    status: 0,
    describe: '',
  });

  async function getAppInfo() {
    try {
      const result = await QueryApplicationById(Number(applicationId));
      Object.assign(appInfo, result);
    } catch (error: any) {
      createErrorModal({
        title: t('common.errorTip'),
        content: t('common.operateFailText'),
      });
    }
  }

  onMounted(async () => {
    getAppInfo();
  });

  async function onClose() {
    router.push('/applications/index');
  }
</script>
