<template>
  <div class="left-line div-box div-card">
    <div class="div-title">{{ t('applications.see.serviceInfo') }}</div>
    <Row :gutter="[10, 10]" class="text-center">
      <Col :span="5" class="font-bold">{{ t('applications.index.nameText') }}</Col>
      <Col :span="5" class="font-bold">{{ t('applications.see.createTime') }}</Col>
      <Col :span="4" class="font-bold">{{ t('applications.index.statusText') }}</Col>
      <Col :span="5" class="font-bold">{{ t('applications.new.leaseTerm') }}</Col>
      <Col :span="5" class="font-bold">{{ t('applications.see.nodeType') }}</Col>
      <Col :span="5">{{ application.name }}</Col>
      <Col :span="5">{{ formatToDateTime(application.createdAt) }}</Col>
      <Col :span="4">
        {{ DictCodeEnum.ApplicationDeployStatus.getOptionLabel(application.status) }}
      </Col>
      <Col :span="5">{{ deployInfo.initialization?.leaseTerm }}</Col>
      <Col :span="5">{{
        pluginConfigs.find((x) => x.value === application.selectNodeType).label ||
        application.selectNodeType ||
        application.plugin
      }}</Col>
    </Row>
  </div>
  <div class="left-line line-color div-card">
    <div class="div-title">{{ t('applications.see.pledgeInfo') }}</div>
    <Row :gutter="[10, 10]">
      <Col :span="12" class="font-bold">{{ t('applications.see.network') }}</Col>
      <Col :span="12" class="font-bold">{{ t('applications.deploy.addressStak') }}</Col>
      <Col :span="12">{{ deployInfo.staking?.networkUrl }}</Col>
      <Col :span="12">{{ deployInfo.staking?.address }}</Col>
      <Col :span="24" class="font-bold">{{ t('applications.see.addressProxy') }}</Col>
      <Col :span="24">{{ deployInfo.staking?.agentAddress }}</Col>
    </Row>
  </div>
  <div class="left-line line-color2 div-card">
    <div class="div-title">{{ t('applications.see.deployInfo') }}</div>
    <Row :gutter="[10, 10]">
      <Col :span="12" class="font-bold">{{ t('applications.see.nodeEthereumUrl') }}</Col>
      <Col :span="12" class="font-bold">{{ t('applications.see.indexerAddress') }}</Col>
      <Col :span="12">
        {{ deployInfo.deployment?.nodeEthereumUrl }}
      </Col>
      <Col :span="12">{{ deployInfo.deployment?.indexerAddress }}</Col>
      <Col :span="12" class="font-bold">{{ t('applications.see.ethereumUrl') }}</Col>
      <Col :span="12" class="font-bold">{{ t('applications.see.ethereumNetwork') }}</Col>
      <Col :span="12">{{ deployInfo.deployment?.ethereumUrl }}</Col>
      <Col :span="12">{{ deployInfo.deployment?.ethereumNetwork }}</Col>
      <Col :span="24" class="font-bold">{{ t('applications.deploy.mnemonic') }}</Col>
      <Col :span="24">
        {{ deployInfo.initialization?.accountMnemonic }}
      </Col>
    </Row>
  </div>
  <div class="text-center my-[40px]">
    <Button
      class="!h-[60px] w-[200px]"
      size="large"
      type="primary"
      shape="round"
      danger
      :loading="isLoading"
      @click="onDelete"
    >
      {{ t('common.delText') }}
    </Button>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import { pluginConfigs } from '/@/utils/thegraphUtil';
  import { DictCodeEnum } from '/@/enums/dictCodeEnum';
  import { DeleteApplication } from '/@wails/go/app/Application';
  import { Button, Row, Col, Modal } from 'ant-design-vue';

  const props = defineProps({
    application: Object as PropType<Recordable>,
    deployInfo: Object as PropType<Recordable>,
  });

  const { t } = useI18n();
  const router = useRouter();

  const isLoading = ref(false);

  const onDelete = () => {
    Modal.confirm({
      title: t('common.confirmText'),
      // icon: '',
      okText: t('common.okText'),
      cancelText: t('common.cancelText'),
      onOk() {
        DeleteApplication(props.application.id).then(() => {
          router.push('/applications');
        });
      },
    });
  };
</script>

<style lang="less" scoped>
  .div-box {
    box-shadow: 0px 6px 20px 0px rgb(31 31 35 / 10%);
  }

  .div-card {
    @apply bg-white;
    border-radius: 20px;
    margin-top: 30px;
    padding: 20px 20px 30px 20px;
  }

  .div-title {
    @apply font-bold;
    font-size: 18px;
    color: #072772;
    margin-bottom: 20px;
  }

  .line-color::before {
    border-color: #e24066;
  }

  .line-color2::before {
    border-color: #009eff;
  }
</style>
