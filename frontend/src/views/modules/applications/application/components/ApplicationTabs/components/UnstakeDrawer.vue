<template>
  <Form>
    <div class="stake-box p-[20px]">
      <div class="flex">
        <img :src="addressAvatar" class="bg-[#D8D8D8] rounded-[50%] w-[60px] h-[60px]" />
        <div class="ml-[16px]">
          <div>{{ t('applications.see.unstake') }}</div>
          <div class="text-[20px] font-bold">{{ shortAddress }}</div>
        </div>
      </div>
      <div class="grid grid-cols-2 text-center mt-[20px]">
        <div class="right-line">
          <div>
            <SvgIcon color="#63A0FA" size="16" name="grt" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.yStake') }}</label>
          </div>
          <div>
            <label class="text-[18px] font-bold">{{ stakeAmount }}</label>
            <label class="text-[12px]">{{ t('applications.see.grt') }}</label>
          </div>
        </div>
        <div>
          <div>
            <SvgIcon color="#63A0FA" size="16" name="balance" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.wBalance') }}</label>
          </div>
          <div>
            <label class="text-[18px] font-bold">{{ addressBalance }}</label>
            <label class="text-[12px]">{{ t('applications.see.grt') }}</label>
          </div>
        </div>
      </div>
    </div>

    <div class="px-[20px] pb-[60px]">
      <div class="font-bold my-[10px]">{{ t('applications.reward.unStakeAmount') }}</div>
      <Input
        class="border !border-[#043CC1] rounded-[8px] h-[60px] px-[10px]"
        v-model:value="inputUnStakeAmount"
        @change="inputUnStakeChange"
      >
        <template #suffix>
          <div>
            <label class="text-[#7B8082] mr-[10px]">{{ t('applications.see.grt') }}</label>
            <label
              class="bg-[#63A0FA] px-[20px] py-[8px] rounded-[4px] text-white"
              @click="maxClick"
            >
              {{ t('applications.see.max') }}
            </label>
          </div>
        </template>
      </Input>
      <!--      <div class="flex justify-between mt-[10px]">-->
      <!--        <div>{{ t('applications.see.minAmount') }}</div>-->
      <!--        <div>-->
      <!--          <label class="text-[18px] font-bold mr-[3px]">10000000000</label-->
      <!--          >{{ t('applications.see.grt') }}-->
      <!--        </div>-->
      <!--      </div>-->
      <!--      <div class="flex justify-between my-[10px]">-->
      <!--        <div>{{ t('applications.see.endPeriod') }}</div>-->
      <!--        <div>-->
      <!--          <label class="text-[18px] font-bold mr-[3px]">28</label>{{ t('applications.see.days') }}-->
      <!--        </div>-->
      <!--      </div>-->
      <Button
        class="mt-[10px]"
        type="primary"
        size="large"
        @click="unStake"
        :loading="unStakeButtonLoading"
        >{{ t('applications.see.unstake') }}</Button
      >
    </div>
  </Form>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { Button, Input, Form } from 'ant-design-vue';
  import { computed, ref } from 'vue';
  import { buildContract, createWeb3Api, runContractMethod, web3Abi } from '/@/utils/web3Util';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { BigNumber } from 'ethers';
  import { thawingPeriod } from '/@/utils/constant';
  import { UpdateThinkingTime } from '/@wails/go/app/Application';
  // defines
  const props = defineProps({
    stakeAmount: String,
    addressBalance: String,
    addressAvatar: String,
    shortAddress: String,
    deployInfo: Object as PropType<Recordable>,
    applicationId: Number,
  });
  const inputUnStakeAmount = ref();
  const unStakeButtonLoading = ref(false);
  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const emits = defineEmits(['close-drawer', 'query-un-stake']);
  const web3Api = computed(() => {
    const { initialization, staking } = props.deployInfo;
    const accountMnemonic = initialization.accountMnemonic;
    const networkUrl = staking.networkUrl;
    if (accountMnemonic && networkUrl) {
      return createWeb3Api(networkUrl, accountMnemonic);
    }
    return undefined;
  });
  const maxClick = () => {
    inputUnStakeAmount.value = props.stakeAmount;
  };
  const inputUnStakeChange = () => {
    inputUnStakeAmount.value = inputUnStakeAmount.value.replace(/[^\d.]/g, '');
  };
  const unStake = async () => {
    const api = web3Api.value;
    if (api && api.__config) {
      unStakeButtonLoading.value = true;
      const address = props.deployInfo?.staking.agentAddress;
      const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);
      try {
        await runContractMethod({
          api,
          contract,
          method: 'unstake',
          methodArgs: [api.utils.toWei(inputUnStakeAmount.value.toString())],
          type: 'send',
        });
        const blockNumber = await api.eth.getBlockNumber();
        const currentBlockNumber = BigNumber.from(blockNumber);
        const period = BigNumber.from(thawingPeriod);
        const untilTime = currentBlockNumber.add(period).add(BigNumber.from('1'));
        await UpdateThinkingTime(props.applicationId, untilTime.toNumber());
        inputUnStakeAmount.value = '';
        emits('close-drawer');
        emits('query-un-stake');
        unStakeButtonLoading.value = false;
      } catch (e: any) {
        unStakeButtonLoading.value = false;
        createErrorModal({
          title: t('common.errorTip'),
          content: e.message,
        });
        console.info('Approve Staking Proxy Contract Error', e);
      }
    }
  };
</script>
<style lang="less" scoped>
  .stake-box {
    box-shadow: 0px 0px 5px 0px rgba(150, 150, 150, 0.2);
  }

  .right-line::after {
    border-color: rgba(216, 216, 216, 0.4);
    top: 30%;
  }

  :deep(.ant-btn-primary) {
    @apply w-full;
    border-radius: 24px;
  }
</style>
