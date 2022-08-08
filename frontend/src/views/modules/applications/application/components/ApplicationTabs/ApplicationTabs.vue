<template>
  <div class="mx-[40px] mb-[100px]">
    <Tabs v-model:activeKey="activeKey">
      <TabPane key="1" :tab="t('applications.see.revenueInfo')">
        <RevenueInfo />
      </TabPane>
      <TabPane key="2" :tab="t('applications.see.subgraph')">2 </TabPane>
      <TabPane key="3" :tab="t('applications.see.serviceDetails')">3 </TabPane>
    </Tabs>
  </div>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  // import { SvgIcon } from '/@/components/Icon';
  import RevenueInfo from './components/RevenueInfo.vue';
  import { Tabs, TabPane } from 'ant-design-vue';

  const { t } = useI18n();

  const activeKey = ref('1');
</script>
<style lang="less" scoped></style>
