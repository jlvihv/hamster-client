<template>
  <Card>
    <Descriptions :column="1" :title="t('applications.see.rewardInfo')" bordered>
      <DescriptionsItem :label="t('applications.reward.account')">
        {{ deployInfo.staking.agentAddress }}
      </DescriptionsItem>
      <DescriptionsItem :label="t('applications.reward.income')">
        <label> {{ income }}</label>
        <Button class="ml-4" type="primary" @click="withdraw">withdraw</Button>
      </DescriptionsItem>
    </Descriptions>
  </Card>
</template>

<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Card, Button, Descriptions, DescriptionsItem } from 'ant-design-vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { computed, ref } from 'vue';
  import { buildContract, createWeb3Api, runContractMethod, web3Abi } from '/@/utils/web3Util';

  // defines
  const props = defineProps({
    deployInfo: Object as PropType<Recordable>,
  });

  const { t } = useI18n();
  const { createConfirm } = useMessage();

  const income = ref(0);

  // web3 api
  const web3Api = computed(() => {
    const { initialization, staking } = props.deployInfo;
    //'clarify height path primary quantum already turtle plate rely hollow frequent exile'
    const accountMnemonic = initialization.accountMnemonic;
    // 'https://rinkeby.infura.io/v3/bab2a1a435b04c07a488d847cf6788f7'
    const networkUrl = staking.networkUrl;

    if (accountMnemonic && networkUrl) {
      return createWeb3Api(networkUrl, accountMnemonic);
    }

    return undefined;
  });

  const getIncome = async () => {
    console.log('get income start');

    const api = web3Api.value;

    const address = props.deployInfo.staking.agentAddress;

    if (api && api.__config) {
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
        }
      },
      iconType: 'warning',
    });
  };

  getIncome().then(() => {
    console.log('get income end');
  });
</script>

<style lang="less" scoped>
  :deep(.ant-descriptions) {
    .ant-descriptions-item-label {
      width: 30%;
    }
  }
</style>
