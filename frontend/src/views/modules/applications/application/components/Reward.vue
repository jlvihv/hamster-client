<template>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.rewardInfo')" bordered>
      <DescriptionsItem :label="t('applications.reward.account')">
        {{ deployInfo.staking.agentAddress }}
        <Button class="mr-1" type="primary" :loading="refresh" @click="getIncome">refresh</Button>
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.reward.income')">
        <label> {{ displayIncome }}</label>
        <Button class="ml-4" type="primary" @click="withdraw">withdraw</Button>
      </DescriptionsItem>
      <DescriptionsItem label="Balance">
        {{ displayBalance }}
      </DescriptionsItem>
    </Descriptions>
  </Card>
</template>

<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Button, Descriptions, DescriptionsItem } from 'ant-design-vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { computed, ref, watch } from 'vue';
  import {
    buildContract,
    createWeb3Api,
    getProviderAddress,
    runContractMethod,
    web3Abi,
  } from '/@/utils/web3Util';

  // defines
  const props = defineProps({
    deployInfo: Object as PropType<Recordable>,
  });

  const { t } = useI18n();
  const { createConfirm } = useMessage();

  const income = ref(0);
  const balance = ref(0);
  const refresh = ref(false);

  // web3 api
  const web3Api = computed(() => {
    console.log(props.deployInfo);
    const { initialization, staking } = props.deployInfo;
    //'clarify height path primary quantum already turtle plate rely hollow frequent exile'
    const accountMnemonic = initialization.accountMnemonic;
    // 'https://rinkeby.infura.io/v3/bab2a1a435b04c07a488d847cf6788f7'
    const networkUrl = staking.networkUrl;

    console.log('accountMnemonic:', accountMnemonic);
    console.log('networkUrl: ', networkUrl);

    if (accountMnemonic && networkUrl) {
      return createWeb3Api(networkUrl, accountMnemonic);
    }

    return undefined;
  });

  const getIncome = async () => {
    console.log('get income start');

    refresh.value = true;
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

      balance.value = balance_data;

      refresh.value = false;

      const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);

      const data = await runContractMethod({
        api,
        contract,
        method: 'gainIncome',
        methodArgs: [],
        type: 'call',
      });

      console.log('get income : ', data);
      income.value = data;
    }
  };

  const withdraw = async () => {
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

  const displayIncome = computed(() => {
    const api = web3Api.value;
    console.log('income to wei:', income.value.toString());
    return api?.utils.fromWei(income.value.toString());
  });

  const displayBalance = computed(() => {
    const api = web3Api.value;
    return api?.utils.fromWei(balance.value.toString());
  });

  watch(
    () => props.deployInfo.staking.agentAddress,
    (_) => {
      getIncome().then(() => {
        console.log('get income end');
      });
    },
  );
</script>

<style lang="less" scoped>
  :deep(.ant-descriptions) {
    .ant-descriptions-item-label {
      width: 30%;
    }
  }
</style>
