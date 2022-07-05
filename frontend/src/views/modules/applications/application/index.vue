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
    <Card :title="t('applications.see.deployInfo')">
      <pre>
nodeEthereumUrl: "mainnet:https://eth-mainnet.alchemyapi.io/v2/wHl5-FZKD68DlhCuUiOXgYh4Z0q0L5fh",   //graph_node 配置的以太坊节点
  ethereumUrl: "https://rinkeby.infura.io/v3/af7a79eb36f64e609b5dda130cd62946",     // index-service,index-agent配置的以太坊节点
  ethereumNetwork: "rinkeby",  //以太坊网络
  indexerAddress: "0x9438BbE4E7AF1ec6b13f75ECd1f53391506A12DF",   //索引人地址
  mnemonic:"please output text solve glare exit divert boil nerve eagle attack turkey"  //索引人助记词
}</pre
      >
    </Card>
    <div class="mt-4 text-right">
      <Button type="primary" @click="onClose">{{ t('common.closeText') }}</Button>
      <!-- 
        <Button>{{ t('common.cancelText') }}</Button>
      <Button class="ml-4" type="primary" >{{ t('common.deployText') }}</Button> -->
    </div>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { PageWrapper } from '/@/components/Page';
  import { useRoute, useRouter } from 'vue-router';
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
