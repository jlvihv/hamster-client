<template>
  <div class="mx-[10px] mb-[100px]">
    <Tabs v-model:activeKey="activeKey" @change="tabChange">
      <TabPane key="1" :tab="t('applications.see.revenueInfo')">
        <RevenueInfo :deployInfo="deployInfo" :application="application" ref="revenueInfoRef" />
      </TabPane>
      <TabPane key="2" :tab="t('applications.see.subgraph')">
        <Subgraph :application="application" />
      </TabPane>
      <TabPane key="3" :tab="t('applications.see.serviceDetails')">
        <ServiceDetails :deployInfo="deployInfo" :application="application" />
      </TabPane>
    </Tabs>
  </div>
</template>
<script lang="ts" setup>
  import { ref, reactive, onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import RevenueInfo from './components/RevenueInfo.vue';
  import Subgraph from './components/Subgraph.vue';
  import ServiceDetails from './components/ServiceDetails.vue';
  import { QueryApplicationById } from '/@wails/go/app/Application';
  import { GetDeployInfo } from '/@wails/go/app/Deploy';
  import { Tabs, TabPane } from 'ant-design-vue';

  const props = defineProps({
    applicationId: Number,
  });

  const { t } = useI18n();

  const activeKey = ref('1');
  const revenueInfoRef = ref();
  const application = reactive({});
  const deployInfo = ref<{
    initialization: Recordable;
    staking: Recordable;
    deployment: Recordable;
  }>({
    initialization: {},
    staking: {},
    deployment: {},
  });
  const tabChange = (key) => {
    if (key == '1') {
      revenueInfoRef.value.initData();
    }
  };
  onMounted(() => {
    QueryApplicationById(props.applicationId).then((app) => {
      console.log(app);
      Object.assign(application, app);
    });
    GetDeployInfo(props.applicationId).then((info) => {
      console.log(info);
      Object.assign(deployInfo.value, info);
    });
  });
</script>
