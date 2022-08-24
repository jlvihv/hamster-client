<template>
  <PageWrapper>
    <Header :showBack="true" />
    <template v-if="!isLoading">
      <ApplicationTabs :applicationId="applicationId" v-if="isAppDeployed" />
      <Deployment :applicationId="applicationId" v-else />
    </template>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { onMounted, reactive, ref, computed } from 'vue';
  import { useRoute } from 'vue-router';
  import { DictCodeEnum } from '/@/enums/dictCodeEnum';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import Deployment from './components/Deployment.vue';
  import ApplicationTabs from './components/ApplicationTabs';
  import Header from '../index/components/Header.vue';
  import { QueryApplicationById } from '/@wails/go/app/Application';
  import { useI18n } from '/@/hooks/web/useI18n';

  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const { params } = useRoute();
  const applicationId = Number(params.id);

  const appInfo = reactive({});
  const isLoading = ref(true);

  const isAppDeployed = computed(() =>
    DictCodeEnum.ApplicationDeployStatus_Running.is(appInfo.status),
  );

  const getAppInfo = async () => {
    isLoading.value = true;

    try {
      const result = await QueryApplicationById(applicationId);
      Object.assign(appInfo, result);
    } catch (error: any) {
      createErrorModal({
        title: t('common.errorTip'),
        content: t('common.operateFailText'),
      });
    } finally {
      isLoading.value = false;
    }
  };

  onMounted(async () => {
    getAppInfo();
  });
</script>
