<template>
  <div class="mx-[40px] mb-[100px]">
    <Tabs v-model:activeKey="activeKey">
      <TabPane key="1" :tab="t('applications.see.revenueInfo')">
        <RevenueInfo @modalConfirm="modalConfirm" />
      </TabPane>
      <TabPane key="2" :tab="t('applications.see.subgraph')">
        <Subgraph />
      </TabPane>
      <TabPane key="3" :tab="t('applications.see.serviceDetails')">
        <ServiceDetails @modalConfirm="modalConfirm" />
      </TabPane>
    </Tabs>
  </div>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import RevenueInfo from './components/RevenueInfo.vue';
  import Subgraph from './components/Subgraph.vue';
  import ServiceDetails from './components/ServiceDetails.vue';
  import { Tabs, TabPane, Modal } from 'ant-design-vue';

  const { t } = useI18n();

  const activeKey = ref('1');
  
  const modalConfirm = () => {
    Modal.confirm({
      title: t('applications.see.receiveBenefitsInfo'),
      icon: '',
      okText: t('common.okText'),
      cancelText: t('common.cancelText'),
      onOk() {
        console.log('OK');
      },
      onCancel() {
        console.log('Cancel');
      },
    });
  };
</script>
