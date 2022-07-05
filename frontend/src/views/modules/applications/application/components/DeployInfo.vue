<template>
  <Card>
    <Descriptions :title="t('applications.see.appInit')" bordered>
      <DescriptionsItem :label="t('applications.deploy.leaseTerm')"
        >Cloud Database</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.publicKey')"
        >Cloud Database</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.importAccount')"
        >Cloud Database</DescriptionsItem
      >
    </Descriptions>
  </Card>
  <div class="h-4"></div>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.pledgeInfo')" bordered>
      <DescriptionsItem :label="t('applications.deploy.addressNetWork')"
        >Cloud Database</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.addressStak')"
        >Cloud Database</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.addressAgent')"
        >Cloud Database</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.pledgeAmount')"
        >Cloud Database</DescriptionsItem
      >
    </Descriptions>
  </Card>
  <div class="h-4"></div>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.deployInfo')" bordered>
      <DescriptionsItem :label="t('applications.deploy.nodeUrl')"
        >"mainnet:https://eth-mainnet.alchemyapi.io/v2/wHl5-FZKD68DlhCuUiOXgYh4Z0q0L5fh",
        //graph_node 配置的以太坊节点</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.ethereumUrl')">
        "https://rinkeby.infura.io/v3/af7a79eb36f64e609b5dda130cd62946", //
        index-service,index-agent配置的以太坊节点</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.ethereumNetWork')"
        >"rinkeby", //以太坊网络</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.indexerAddress')"
        >"0x9438BbE4E7AF1ec6b13f75ECd1f53391506A12DF", //索引人地址</DescriptionsItem
      >
      <DescriptionsItem :label="t('applications.deploy.mnemonic')"
        >"please output text solve glare exit divert boil nerve eagle attack turkey"
        //索引人助记词</DescriptionsItem
      >
    </Descriptions>
  </Card>
</template>
<script lang="ts" setup>
  import { useRoute } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Descriptions, DescriptionsItem } from 'ant-design-vue';

  const { t } = useI18n();
  const { params } = useRoute();
  const { id: applicationId } = params;
  console.log(applicationId);
</script>
