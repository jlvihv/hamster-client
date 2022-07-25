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
      <DescriptionsItem :label="t('applications.reward.income')">
        <label> {{ income }}</label>
        <Button class="ml-3" type="primary" @click="withdraw">
          {{ t('applications.reward.withdraw') }}
        </Button>
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.reward.balance')">
        {{ balance }}
      </DescriptionsItem>
    </Descriptions>
  </Card>
</template>

<script lang="ts" setup>
  import { computed, ref, watchEffect } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Button, Descriptions, DescriptionsItem } from 'ant-design-vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { LoadingOutlined } from '@ant-design/icons-vue';
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
  const isRefreshing = ref(false);

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

      const data = await runContractMethod({
        api,
        contract,
        method: 'gainIncome',
        methodArgs: [],
        type: 'call',
      });

      console.log('get income : ', data);
      income.value = api.utils.fromWei(data.toString());
    }

    console.log('get income end');
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
