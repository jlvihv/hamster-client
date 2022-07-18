<template>
  <Form layout="vertical" :model="formData" :rules="formRules" ref="formRef">
    <FormItem :label="t('applications.deploy.selectNetWork')" name="networkUrl">
      <Select
        :allowClear="true"
        class="input-width"
        v-model:value="formData.networkUrl"
        :options="networkOptions"
        :placeholder="t('applications.deploy.selectNetWorkInfo')"
        @change="handleEndpointChange"
      />
    </FormItem>
    <FormItem :label="t('applications.deploy.addressStak')" name="address">
      <Input :allowClear="true" class="input-width" v-model:value="formData.address" disabled />
      <Button
        class="ml-4"
        type="primary"
        @click="generateStakingProxyContract"
        v-if="!formData.agentAddress"
      >
        {{ t('applications.deploy.addressStakBtn') }}
      </Button>
    </FormItem>
    <FormItem :label="t('applications.deploy.addressAgent')" name="agentAddress">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value="formData.agentAddress"
        disabled
      />
      <Modal
        v-model:visible="pledgeAmountModalVisible"
        title="Tips"
        :okText="t('common.confirmText')"
        @ok="approveStakingProxyContract"
        :okButtonProps="{ disabled: !pledgeAmountInModal }"
      >
        <Form layout="vertical">
          <FormItem :label="t('applications.deploy.pledgeAmount')" name="pledgeAmount">
            <InputNumber v-model:value="pledgeAmountInModal" :min="1" />
          </FormItem>
        </Form>
      </Modal>
      <Button
        class="ml-4"
        type="primary"
        @click="pledgeAmountModalVisible = true"
        v-if="formData.pledgeAmount == null"
      >
        {{ t('applications.deploy.addressAgentBtn') }}
      </Button>
    </FormItem>
    <FormItem :label="t('applications.deploy.pledgeAmount')" name="pledgeAmount">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value="formData.pledgeAmount"
        disabled
      />
      <Button class="ml-4" type="primary" @click="stakeAmount">
        {{ t('applications.deploy.pledgeAmountBtn') }}
      </Button>
    </FormItem>
    <FormItem class="text-center">
      <Button class="w-32 mt-6 ml-4" type="primary" @click="handleSubmit">
        {{ t('common.nextText') }}
      </Button>
    </FormItem>
  </Form>
</template>

<script lang="ts" setup>
  import { ref, toRaw, toRefs, computed } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { createRule } from '/@/utils/formUtil';
  import {
    createWeb3Api,
    web3Configs,
    web3Abi,
    getProviderAddress,
    buildContract,
    runContractMethod,
  } from '/@/utils/web3Util';
  import { SaveDeployInfo } from '/@wails/go/app/Deploy';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { Form, FormItem, Button, Input, InputNumber, Select, Modal } from 'ant-design-vue';

  // defines
  const props = defineProps({
    applicationId: Number,
    deployInfo: Object as PropType<Recordable>,
  });
  const { applicationId, deployInfo } = toRefs(props);

  const emits = defineEmits(['update:deployInfo', 'submited']);

  const { t } = useI18n();
  const { createConfirm } = useMessage();

  const formRef = ref();
  const formData: {
    networkUrl?: string;
    address?: string;
    agentAddress: string;
    pledgeAmount?: number;
  } = deployInfo.value.staking;

  const formRules = computed(() => ({
    networkUrl: [createRule(t('applications.deploy.selectNetwork'))],
    address: [createRule(t('applications.deploy.selectNetworkAbove'))],
    agentAddress: [createRule(t('applications.deploy.generateStakingProxyContract'))],
    pledgeAmount: [createRule(t('applications.deploy.approveStakingProxyContract'))],
  }));

  // networkOptions,
  // address is up on endpoint field
  const networkOptions = web3Configs.map(({ endpoint }) => ({ label: endpoint, value: endpoint }));
  const handleEndpointChange = (newEndpoint) => {
    if (newEndpoint) {
      const config = web3Configs.find((x) => x.endpoint === newEndpoint);
      formData.address = config?.stakeContractAddress;
    } else {
      formData.address = undefined;
    }
  };

  // web3 api
  const web3Api = computed(() => {
    const { initialization, staking } = deployInfo.value;

    if (initialization.accountMnemonic && staking.networkUrl) {
      return createWeb3Api(staking.networkUrl, initialization.accountMnemonic);
    }

    return undefined;
  });

  // Generate Staking Proxy Contract
  const generateStakingProxyContract = () => {
    createConfirm({
      title: 'Confirm',
      content: 'Are you sure to generate staking proxy contract?',
      onOk: async () => {
        const api = web3Api.value;

        if (api && api.__config) {
          const { factoryContractAddress } = api.__config;
          const providerAddress = getProviderAddress(api);
          const contract = buildContract(api, web3Abi.stakeProxyFactoryAbi, factoryContractAddress);

          try {
            await runContractMethod({
              api,
              contract,
              method: 'createStakingContract',
              methodArgs: [providerAddress],
            });
          } catch (error: any) {
            console.log(error);
          }

          const agentAddress = await runContractMethod({
            api,
            contract,
            type: 'call',
            method: 'getStakingAddress',
            methodArgs: [providerAddress],
          });

          formData.agentAddress = agentAddress;
        }
      },
    });
  };

  // Approve Staking Proxy Contract
  const pledgeAmountModalVisible = ref(false);
  const pledgeAmountInModal = ref<number | undefined>();
  const approveStakingProxyContract = async () => {
    const api = web3Api.value;

    if (api && api.__config) {
      const { erc20ContractAddress } = api.__config;
      const contract = buildContract(api, web3Abi.ecr20Abi, erc20ContractAddress);
      const pledgeAmount = pledgeAmountInModal.value || 1;

      await runContractMethod({
        api,
        contract,
        method: 'approve',
        methodArgs: [formData.agentAddress, pledgeAmount],
      });

      pledgeAmountModalVisible.value = false;
      formData.pledgeAmount = pledgeAmount;
    }
  };

  // Stack Amount
  const stakeAmount = async () => {
    createConfirm({
      title: 'Confirm',
      content: `Are you sure to stake amount: ${formData.pledgeAmount} ?`,
      onOk: async () => {
        const api = web3Api.value;

        if (api && api.__config) {
          const contract = buildContract(
            api,
            web3Abi.stakeDistributionProxyAbi,
            formData.agentAddress,
          );

          const data = await runContractMethod({
            api,
            contract,
            method: 'staking',
            methodArgs: [formData.pledgeAmount],
          });

          console.log('stakeAmount data', data);
        }
      },
    });
  };

  const handleSubmit = async () => {
    await formRef.value?.validate();

    await SaveDeployInfo(applicationId.value, JSON.stringify(toRaw(deployInfo.value)));
    emits('submited', formData);
  };
</script>

<style lang="less" scoped>
  :deep(.ant-form-item) {
    .ant-input-number {
      width: 100%;
    }
  }
</style>
