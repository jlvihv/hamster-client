<template>
  <PageWrapper>
    <div class="rounded-[10px] bg-[#FFFFFF]">
      <div class="text-[16px] font-bold p-[24px]">{{ t('applications.deploy.appDeploy') }}</div>
      <div class="bg-[#F5F5F5]">
        <Steps type="navigation" labelPlacement="vertical" :current="currentStep">
          <Step
            v-for="item in [1, 2, 3]"
            :key="item"
            :title="t(`applications.deploy.stepTitle${item}`)"
          />
        </Steps>
      </div>
      <div class="p-[24px]">
        <div v-if="currentStep === 0">
          <Form layout="vertical" v-model:model="formInit">
            <FormItem :label="t('applications.deploy.leaseTerm')" name="lease_term">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formInit.lease_term"
                :placeholder="t('common.inputText') + t('applications.deploy.leaseTerm')"
              />
            </FormItem>
            <FormItem :label="t('applications.deploy.publicKey')" name="public_key">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formInit.public_key"
                :placeholder="t('common.inputText') + t('applications.deploy.publicKey')"
              />
            </FormItem>
            <FormItem :label="t('applications.deploy.importAccount')" name="import_account">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formInit.import_account"
                :placeholder="t('common.inputText') + t('applications.deploy.importAccount')"
              />
            </FormItem>
            <FormItem class="text-right">
              <Button @click="tabAction.onCancel">{{ t('common.cancelText') }}</Button>
              <Button class="ml-4" type="primary" @click="tabAction.onNext(1)">{{
                t('common.nextText')
              }}</Button>
            </FormItem>
          </Form>
        </div>
        <div v-if="currentStep === 1">
          <Form layout="vertical" v-model:model="formStak">
            <FormItem :label="t('applications.deploy.selectNetWork')" name="network">
              <Select
                :allowClear="true"
                class="input-width"
                @change="stakAction.changeNetWork"
                v-model:value="formStak.network"
                :options="networkOptions"
                :placeholder="t('applications.deploy.selectNetWorkInfo')"
              />
            </FormItem>
            <FormItem :label="t('applications.deploy.addressStak')" name="address_stak">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formStak.address_stak"
                disabled
              />
              <Button class="ml-4" type="primary" @click="visibleAddress = true">{{
                t('applications.deploy.addressStakBtn')
              }}</Button>
            </FormItem>
            <FormItem :label="t('applications.deploy.addressAgent')" name="address_agent">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formStak.address_agent"
                disabled
              />
              <Button class="ml-4" type="primary" @click="visibleAmount = true">{{
                t('applications.deploy.addressAgentBtn')
              }}</Button>
            </FormItem>
            <FormItem :label="t('applications.deploy.pledgeAmount')" name="pledge_amount">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formStak.pledge_amount"
                disabled
              />
              <Button class="ml-4" type="primary" @click="visibleStack = true">{{
                t('applications.deploy.pledgeAmountBtn')
              }}</Button>
            </FormItem>
            <FormItem class="text-right">
              <Button @click="tabAction.onLast(0)">{{ t('common.cancelText') }}</Button>
              <Button class="ml-4" type="primary" @click="tabAction.onNext(2)">{{
                t('common.nextText')
              }}</Button>
            </FormItem>
          </Form>
        </div>
        <div v-if="currentStep === 2">
          <Form layout="vertical" v-model:model="formDeploy">
            <FormItem :label="t('applications.deploy.nodeMainnet')" name="mainnet">
              <Input :allowClear="true" class="input-width" v-model:value="formDeploy.mainnet" />
            </FormItem>
            <FormItem :label="t('applications.deploy.ethereumUrl')" name="ethereum_url">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formDeploy.ethereum_url"
              />
            </FormItem>
            <FormItem :label="t('applications.deploy.ethereumNetWork')" name="ethereum_network">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formDeploy.ethereum_network"
              />
            </FormItem>
            <FormItem :label="t('applications.deploy.indexerAddress')" name="indexer_address">
              <Input
                :allowClear="true"
                class="input-width"
                v-model:value="formDeploy.indexer_address"
              />
            </FormItem>
            <FormItem class="text-right">
              <Button @click="tabAction.onLast(1)">{{ t('common.cancelText') }}</Button>
              <Button class="ml-4" type="primary" @click="tabAction.onNext(3)">{{
                t('common.nextText')
              }}</Button>
            </FormItem>
          </Form>
        </div>
      </div>
    </div>
    <Modal
      v-model:visible="visibleStack"
      title="Tips"
      :okText="t('common.confirmText')"
      @ok="stakAction.generateStack"
    >
      <div class="text-center">{{ t('applications.deploy.pledgeAmountTips') }}</div>
    </Modal>
    <Modal
      v-model:visible="visibleAddress"
      title="Tips"
      :okText="t('common.confirmText')"
      @ok="stakAction.generateStack"
    >
      <div class="text-center">{{ t('applications.deploy.addressStakTips') }}</div>
    </Modal>
    <Modal
      v-model:visible="visibleAmount"
      title="Tips"
      :okText="t('common.confirmText')"
      @ok="stakAction.onPledgeAmount"
    >
      <Form v-model:model="formStak">
        <FormItem :label="t('applications.deploy.pledgeAmount')" name="pledge_amount">
          <Input :allowClear="true" v-model:value="formStak.pledge_amount" />
        </FormItem>
      </Form>
    </Modal>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useRoute, useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Steps, Step, Form, FormItem, Button, Input, Select, Modal } from 'ant-design-vue';

  const { t } = useI18n();
  const router = useRouter();
  const { params } = useRoute();
  const { id: applicationId } = params;
  console.log('applicationId:', applicationId);

  const currentStep = ref(0);
  const visibleAmount = ref(false);
  const visibleStack = ref(false);
  const visibleAddress = ref(false);
  const networkOptions = reactive([{ label: 'http://193.65.66.207:9500', value: 1 }]);
  const formInit = reactive({
    lease_term: '',
    public_key: '',
    import_account: '',
  });
  const formStak = reactive({
    network: undefined,
    address_stak: '',
    address_agent: '',
    pledge_amount: '',
  });
  const formDeploy = reactive({
    mainnet: '',
    ethereum_url: '',
    ethereum_network: '',
    indexer_address: '',
  });

  const stakAction = {
    async changeNetWork() {},
    async generateStack() {},
    async showPledgeAmount() {
      visibleAmount.value = true;
    },
    async onPledgeAmount() {},
    async stackAmount() {},
  };

  const tabAction = {
    async onCancel() {
      router.push('/applications/' + applicationId);
    },
    async onLast(lastStep) {
      currentStep.value = lastStep;
    },
    async onNext(nextStep) {
      if (nextStep === 3) {
        router.push('/applications/' + applicationId);
      } else {
        currentStep.value = nextStep;
      }
    },
  };
</script>
<style lang="less" scoped>
  :deep(.input-width) {
    width: 70%;
  }
</style>
