<template>
  <PageWrapper>
    <Card>
      <Descriptions :title="t('applications.see.appInfo')" bordered>
        <template #extra v-if="true">
          <router-link :to="`/applications/${applicationId}/deploy`">Deploy</router-link>
        </template>
        <DescriptionsItem :label="t('applications.index.nameText')"
          >Cloud Database</DescriptionsItem
        >
        <DescriptionsItem :label="t('applications.index.addTimeText')"
          >Cloud Database</DescriptionsItem
        >
        <DescriptionsItem :label="t('applications.index.statusText')"
          >Cloud Database</DescriptionsItem
        >
        <DescriptionsItem :label="t('applications.index.desText')">Cloud Database</DescriptionsItem>
      </Descriptions>
    </Card>
    <div class="h-4"></div>
    <deployInfo />
    <div class="mt-4 text-right">
      <Button type="primary" @click="onClose">{{ t('common.closeText') }}</Button>
    </div>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { PageWrapper } from '/@/components/Page';
  import { useRoute, useRouter } from 'vue-router';
  import deployInfo from './components/deployInfo.vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Descriptions, DescriptionsItem, Button } from 'ant-design-vue';

  const { t } = useI18n();
  const router = useRouter();
  const { params } = useRoute();
  const { id: applicationId } = params;

  async function onClose() {
    router.push('/applications/index');
  }
</script>
