<template>
  <PageWrapper>
    <!-- Application Id: {{ $route.params.id }} -->
    <Card>
      <Descriptions :title="t('applications.see.appInfo')" bordered>
        <DescriptionsItem :label="t('applications.index.nameText')"
          >Cloud Database</DescriptionsItem
        >
        <DescriptionsItem :label="t('applications.index.abbText')">Cloud Database</DescriptionsItem>
        <DescriptionsItem :label="t('applications.index.addTimeText')"
          >Cloud Database</DescriptionsItem
        >
        <DescriptionsItem :label="t('applications.index.statusText')"
          >Cloud Database</DescriptionsItem
        >
        <DescriptionsItem :label="t('applications.index.desText')">Cloud Database</DescriptionsItem>
      </Descriptions>
    </Card>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useRoute } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Descriptions, DescriptionsItem } from 'ant-design-vue';

  const { t } = useI18n();
  const { params } = useRoute();
  const developerId = ref(params.id);
  console.log('developerId:', developerId.value);
</script>
