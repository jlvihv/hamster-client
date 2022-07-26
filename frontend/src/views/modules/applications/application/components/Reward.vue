<template>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.rewardInfo')" bordered>
      <DescriptionsItem :label="t('applications.reward.account')">
        <LoadingOutlined v-if="isRefreshing" />
        <template v-else>
          {{ deployInfo.staking.agentAddress }}
          <Button class="ml-3" type="primary" @click="getIncome">
            {{ t('applications.reward.refresh') }}
          </Button>
        </template>
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.reward.stake')">
        <label> {{ stakeTotal }}</label>
        <Button class="ml-3" type="primary" @click="showUnStakeModal" v-if="stakeTotal != '0'">
          {{ t('applications.reward.unStake') }}
        </Button>
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.reward.unStaking')">
        <label> {{ unStakeAmount }}</label>
        <Button class="ml-3" type="primary" @click="withdrawStake" v-if="unStakeAmount != '0'">
          {{ t('applications.reward.withdraw') }}
        </Button>
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.reward.income')">
        <label> {{ income }}</label>
        <Button class="ml-3" type="primary" @click="withdraw" v-if="income != '0'">
          {{ t('applications.reward.withdraw') }}
        </Button>
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.reward.balance')">
        {{ balance }}
      </DescriptionsItem>
    </Descriptions>
  </Card>
  <Modal
    v-model:visible="unStakeModalVisible"
    :title="t('applications.reward.unStakeModalTitle')"
    :footer="null"
  >
    <Form layout="vertical" :model="formData" :rules="formRules" ref="formRef">
      <FormItem :label="t('applications.reward.unStakeAmount')" name="unStakeParam">
        <Input
          :allowClear="true"
          class="input-width"
          v-model:value="formData.unStakeParam"
          :placeholder="t('applications.reward.unStakePlaceholder')"
          @change="stakeAmountChange"
        />
      </FormItem>
      <FormItem class="text-center">
        <Button
          class="w-32 mt-6 ml-4"
          type="primary"
          @click="unStake"
          :loading="unStakeModalLoading"
        >
          {{ t('common.confirmText') }}
        </Button>
      </FormItem>
    </Form>
  </Modal>
</template>

<script lang="ts" setup>
  import { computed, reactive, ref, watchEffect } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    Card,
    Button,
    Descriptions,
    DescriptionsItem,
    Modal,
    Form,
    FormItem,
    Input,
  } from 'ant-design-vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { LoadingOutlined } from '@ant-design/icons-vue';
  import {
    buildContract,
    createWeb3Api,
    getProviderAddress,
    runContractMethod,
    web3Abi,
  } from '/@/utils/web3Util';
  import { createRule } from '/@/utils/formUtil';

  // defines
  const props = defineProps({
    deployInfo: Object as PropType<Recordable>,
  });

  const { t } = useI18n();

  const { createConfirm, createErrorModal } = useMessage();

  const income = ref('0');
  const balance = ref(0);
  const isRefreshing = ref(false);
  //un stake
  const stakeTotal = ref('0');
  const unStakeAmount = ref('0');
  const formData = reactive<{
    unStakeParam: number;
  }>({});
  const unStakeModalVisible = ref(false);
  const unStakeModalLoading = ref(false);
  const formRef = ref();
  const formRules = computed(() => ({
    unStakeParam: [createRule(t('applications.reward.unStakePlaceholder'))],
  }));
  const stakeAmountChange = () => {
    formData.unStakeParam = formData.unStakeParam.replace(/[^\-?\d.]/g, '');
  };
  // web3 api
  const web3Api = computed(() => {
    const { initialization, staking } = props.deployInfo;
    const accountMnemonic = initialization.accountMnemonic;
    const networkUrl = staking.networkUrl;

    if (accountMnemonic && networkUrl) {
      return createWeb3Api(networkUrl, accountMnemonic);
    }

    return undefined;
  });

  const getIncome = async () => {
    console.log('get income start');

    isRefreshing.value = true;
    const api = web3Api.value;

    const address = props.deployInfo.staking.agentAddress;

    if (api && api.__config) {
      const erc20 = buildContract(api, web3Abi.ecr20Abi, api.__config.erc20ContractAddress);

      const ethAddress = getProviderAddress(api);
      const balance_data = await runContractMethod({
        api,
        contract: erc20,
        method: 'balanceOf',
        methodArgs: [ethAddress],
        type: 'call',
      });

      console.log('balance_data:', balance_data);

      balance.value = api.utils.fromWei(balance_data.toString());

      isRefreshing.value = false;

      const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);
      try {
        const data = await runContractMethod({
          api,
          contract,
          method: 'gainIncome',
          methodArgs: [],
          type: 'call',
        });

        console.log('get income : ', data);
        income.value = api.utils.fromWei(data.toString());
      } catch (e: any) {
        income.value = '0';
        console.info(e.message);
      }
    }

    console.log('get income end');
  };
  const getStakeTotal = async () => {
    const api = web3Api.value;
    if (api && api.__config) {
      const address = props.deployInfo.staking.agentAddress;
      const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);
      try {
        const data = await runContractMethod({
          api,
          contract,
          method: 'getStakingAmount',
          methodArgs: [],
          type: 'call',
        });
        stakeTotal.value = api.utils.fromWei(data.toString(), 'micro');
      } catch (e: any) {
        stakeTotal.value = '0';
        console.info(e.message);
      }
    }
    console.info('get total stake end');
  };
  const getUnStakeAmount = async () => {
    const api = web3Api.value;
    if (api && api.__config) {
      const address = props.deployInfo.staking.agentAddress;
      const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);
      try {
        const data = await runContractMethod({
          api,
          contract,
          method: 'getUnStakingAmount',
          methodArgs: [],
          type: 'call',
        });
        unStakeAmount.value = api.utils.fromWei(data.toString(), 'micro');
      } catch (e: any) {
        unStakeAmount.value = '0';
        console.info(e.message);
      }
    }
    console.info('get total stake end');
  };
  const withdrawStake = () => {
    createConfirm({
      title: 'Confirm',
      content: 'Are you sure to withdraw ?',
      onOk: async () => {
        const api = web3Api.value;
        const address = props.deployInfo.staking.agentAddress;
        if (api && api.__config) {
          const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);

          await runContractMethod({
            api,
            contract,
            method: 'withdraw',
            methodArgs: [],
            type: 'send',
          });
          await getUnStakeAmount();
        }
      },
      iconType: 'warning',
    });
  };
  const showUnStakeModal = () => {
    unStakeModalVisible.value = true;
  };
  const unStake = async () => {
    await formRef.value?.validate();
    unStakeModalLoading.value = true;
    const api = web3Api.value;
    const address = props.deployInfo.staking.agentAddress;
    if (api && api.__config) {
      const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);
      try {
        await runContractMethod({
          api,
          contract,
          method: 'unstake',
          methodArgs: [api.utils.toWei(formData.unStakeParam.toString(), 'micro')],
          type: 'send',
        });
        unStakeModalLoading.value = false;
        unStakeModalVisible.value = false;
        formData.unStakeParam = '';
      } catch (e: any) {
        unStakeModalLoading.value = false;
        createErrorModal({
          title: t('common.errorTip'),
          content: e.message,
        });
      }
    }
  };
  const withdraw = () => {
    createConfirm({
      title: 'Confirm',
      content: 'Are you sure to withdraw ?',
      onOk: async () => {
        const api = web3Api.value;
        const address = props.deployInfo.staking.agentAddress;
        if (api && api.__config) {
          const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);

          const data = await runContractMethod({
            api,
            contract,
            method: 'withdrawIncome',
            methodArgs: [],
            type: 'send',
          });

          console.log('get income :', data);
          getIncome().then(() => {
            console.log('refresh income');
          });
        }
      },
      iconType: 'warning',
    });
  };
  watchEffect(() => {
    if (props.deployInfo.staking.agentAddress) {
      getStakeTotal();
      getUnStakeAmount();
      setTimeout(getIncome, 100);
    }
  });
</script>

<style lang="less" scoped>
  :deep(.ant-descriptions) {
    .ant-descriptions-item-label {
      width: 30%;
    }
  }
</style>
