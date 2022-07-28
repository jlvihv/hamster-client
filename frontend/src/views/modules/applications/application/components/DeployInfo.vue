<template>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.appInit')" bordered>
      <DescriptionsItem :label="t('applications.deploy.leaseTerm')">
        {{ deployInfo.initialization.leaseTerm }}
      </DescriptionsItem>
      <!--      <DescriptionsItem :label="t('applications.deploy.publicKey')">-->
      <!--        {{ deployInfo.initialization.publicKey }}-->
      <!--      </DescriptionsItem>-->
      <DescriptionsItem :label="t('applications.deploy.importAccount')">
        {{ deployInfo.initialization.accountMnemonic }}
      </DescriptionsItem>
    </Descriptions>
  </Card>
  <div class="h-4"></div>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.pledgeInfo')" bordered>
      <DescriptionsItem :label="t('applications.deploy.addressNetWork')">
        {{ deployInfo.staking.networkUrl }}
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.deploy.addressStak')">
        {{ deployInfo.staking.address }}
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.deploy.addressAgent')">
        {{ deployInfo.staking.agentAddress }}
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.deploy.pledgeAmount')">
        {{ deployInfo.staking.pledgeAmount }}
      </DescriptionsItem>
    </Descriptions>
  </Card>
  <div class="h-4"></div>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.deployInfo')" bordered>
      <DescriptionsItem :label="t('applications.deploy.nodeMainnet')">
        {{ deployInfo.deployment.nodeEthereumUrl }}
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.deploy.ethereumUrl')">
        {{ deployInfo.deployment.ethereumUrl }}
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.deploy.ethereumNetwork')">
        {{ deployInfo.deployment.ethereumNetwork }}
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.deploy.indexerAddress')">
        {{ deployInfo.deployment.indexerAddress }}
      </DescriptionsItem>
    </Descriptions>
  </Card>
</template>

<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Descriptions, DescriptionsItem } from 'ant-design-vue';

  // defines
  defineProps({
    deployInfo: Object as PropType<Recordable>,
  });

  const { t } = useI18n();
</script>

<style lang="less" scoped>
  :deep(.ant-descriptions) {
    .ant-descriptions-item-label {
      width: 30%;
    }
  }
</style>
