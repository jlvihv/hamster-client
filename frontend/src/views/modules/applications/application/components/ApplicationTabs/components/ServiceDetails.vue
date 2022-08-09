<template>
  <div class="left-line div-box div-card">
    <div class="div-title">{{ t('applications.see.serviceInfo') }}</div>
    <Row :gutter="[10,10]" class="text-center">
      <Col :span="5" class="font-bold">{{ t('applications.index.nameText') }}</Col>
      <Col :span="5" class="font-bold">{{ t('applications.see.createTime') }}</Col>
      <Col :span="4" class="font-bold">{{ t('applications.index.statusText') }}</Col>
      <Col :span="5" class="font-bold">{{ t('applications.new.leaseTerm') }}</Col>
      <Col :span="5" class="font-bold">{{ t('applications.see.nodeType') }}</Col>
      <Col :span="5">Thegraph</Col>
      <Col :span="5">2022 05 20 18:56:30</Col>
      <Col :span="4">Running</Col>
      <Col :span="5">128</Col>
      <Col :span="5">The Graph Rinkeby network</Col>
    </Row>
  </div>
  <div class="left-line line-color div-card">
    <div class="div-title">{{ t('applications.see.pledgeInfo') }}</div>
    <Row :gutter="[10,10]">
      <Col :span="12" class="font-bold">{{ t('applications.see.network') }}</Col>
      <Col :span="12" class="font-bold">{{ t('applications.deploy.addressStak') }}</Col>
      <Col :span="12">https://eth-mainnet.alchemyapi.io/v2/wHl5-FZKD68DlhCuUiOXgYh4Z0q0L5fh</Col>
      <Col :span="12">0xd5f6e31199220a0d5334cad2b6ecd70c8f1a6b79</Col>
      <Col :span="24" class="font-bold">{{ t('applications.see.addressProxy') }}</Col>
      <Col :span="24">0xd5f6e31199220a0d5334cad2b6ecd70c8f1a6b79</Col>
    </Row>
  </div>
  <div class="left-line line-color2 div-card">
    <div class="div-title">{{ t('applications.see.deployInfo') }}</div>
    <Row :gutter="[10,10]">
      <Col :span="12" class="font-bold">{{ t('applications.see.nodeEthereumUrl') }}</Col>
      <Col :span="12" class="font-bold">{{ t('applications.see.indexerAddress') }}</Col>
      <Col :span="12">mainnet:https://eth-mainnet.alchemyapi.io/v2/wHl5-FZKD68DlhCuUiOXgYh4Z0q0L5fh</Col>
      <Col :span="12">0x9438BbE4E7AF1ec6b13f75ECd1f53391506A12DF</Col>
      <Col :span="12" class="font-bold">{{ t('applications.see.ethereumUrl') }}</Col>
      <Col :span="12" class="font-bold">{{ t('applications.see.ethereumNetwork') }}</Col>
      <Col :span="12">https://rinkeby.infura.io/v3/af7a79eb36f64e609b5dda130cd62946</Col>
      <Col :span="12">Rinkeby</Col>
      <Col :span="24" class="font-bold">{{ t('applications.deploy.mnemonic') }}</Col>
      <Col :span="24">Please output text solve glare exit divert boil nerve eagle attack turkey  </Col>
    </Row>
  </div>
  <div class="text-center my-[40px]">
    <Button class="!h-[60px] w-[200px]" size="large" type="primary" danger @click="onDelete">
      <svg :class="[btnLoading===false?'hidden' : 'inline']" class="animate-spin align-text-bottom -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      {{ t('common.delText') }}
    </Button>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Button, Row, Col } from 'ant-design-vue';

  const { t } = useI18n();
  const emits = defineEmits(['modalConfirm']);
  const btnLoading = ref(false);

  const onDelete = () => {
    emits('modalConfirm');
  }
</script>
<style lang="less" scoped>
.div-box {
  box-shadow: 0px 6px 20px 0px rgb(31 31 35 / 10%)
}
.div-card{
  @apply bg-white;
  border-radius: 20px;
  margin-top: 30px;
  padding: 20px 20px 30px 20px;
}
.div-title{
  @apply font-bold;
  font-size: 18px;
  color: #072772;
  margin-bottom: 20px;
}
.line-color::before {
  border-color: #E24066;
}
.line-color2::before {
  border-color: #009EFF;
}
</style>