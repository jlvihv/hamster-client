<template>
  <div>
    <div class="stake-box p-[20px]">
      <div class="flex">
        <img :src="addressAvatar" class="bg-[#D8D8D8] rounded-[50%] w-[60px] h-[60px]" />
        <div class="ml-[16px]">
          <div>{{ t('applications.see.withdraw') }}</div>
          <div class="text-[20px] font-bold">{{ shortAddress }}</div>
        </div>
      </div>
      <div class="grid grid-cols-2 text-center mt-[20px]">
        <div class="right-line">
          <div>
            <SvgIcon color="#63A0FA" size="16" name="with" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.withdraw') }}</label>
          </div>
          <div>
            <label class="text-[18px] font-bold">{{ unStakeAmount }}</label>
            <label class="text-[12px]">{{ t('applications.see.grt') }}</label>
          </div>
        </div>
        <div>
          <div>
            <SvgIcon color="#63A0FA" size="16" name="time" />
            <label class="text-[#7B8082] ml-[6px]">{{ t('applications.see.thawLeft') }}</label>
          </div>
          <div class="text-[18px] font-bold">1D 0H 19M</div>
        </div>
      </div>
    </div>
    <div class="px-[20px] pb-[60px]">
      <Button
        class="mt-[20px]"
        type="primary"
        size="large"
        @click="withdraw"
        :loading="withdrawButtonLoading"
      >
        {{ t('applications.see.withdraw') }}
      </Button>
    </div>
  </div>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SvgIcon } from '/@/components/Icon';
  import { Button } from 'ant-design-vue';
  import { computed, ref } from 'vue';
  import { buildContract, createWeb3Api, runContractMethod, web3Abi } from '/@/utils/web3Util';
  import { useMessage } from '/@/hooks/web/useMessage';
  // defines
  const props = defineProps({
    unStakeAmount: String,
    addressAvatar: String,
    shortAddress: String,
    deployInfo: Object as PropType<Recordable>,
  });
  const { t } = useI18n();
  const emits = defineEmits(['close-drawer', 'query-un-stake', 'query-stake']);
  const { createErrorModal } = useMessage();
  const withdrawButtonLoading = ref(false);
  const web3Api = computed(() => {
    const { initialization, staking } = props.deployInfo;
    const accountMnemonic = initialization.accountMnemonic;
    const networkUrl = staking.networkUrl;
    if (accountMnemonic && networkUrl) {
      return createWeb3Api(networkUrl, accountMnemonic);
    }
    return undefined;
  });
  const withdraw = async () => {
    const api = web3Api.value;
    if (api) {
      withdrawButtonLoading.value = true;
      const address = props.deployInfo?.staking.agentAddress;
      const contract = buildContract(api, web3Abi.stakeDistributionProxyAbi, address);
      try {
        await runContractMethod({
          api,
          contract,
          method: 'withdraw',
          methodArgs: [],
          type: 'send',
        });
        emits('close-drawer');
        emits('query-un-stake');
        emits('query-stake');
        withdrawButtonLoading.value = false;
      } catch (e: any) {
        withdrawButtonLoading.value = false;
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
