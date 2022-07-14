<template>
  <PageWrapper>
    <div class="rounded-[10px] bg-white">
      <div class="text-[16px] font-bold p-[24px]">{{ t('applications.deploy.appDeploy') }}</div>
      <div class="bg-[#F5F5F5]">
        <Steps type="navigation" labelPlacement="vertical" :current="currentStep">
          <Step
            v-for="item in steps.length"
            :key="item"
            :title="t(`applications.deploy.stepTitle${item}`)"
          />
        </Steps>
      </div>
      <div class="p-[24px]">
        <component
          :is="steps[currentStep]"
          :applicationId="applicationId"
          v-model:deployInfo="deployInfo"
          @submited="handleNext"
        />
      </div>
    </div>
    <Modal
      width="80%"
      :bodyStyle="previewModalBodyStyle"
      v-model:visible="previewModalVisible"
      :title="t('applications.deploy.previewModalTitle')"
      :okText="t('common.deployText')"
      @ok="passwordModalVisible = true"
    >
      <DeployInfo :deployInfo="deployInfo" />
    </Modal>
    <Modal
      v-model:visible="passwordModalVisible"
      :title="t('applications.deploy.passwordModalTitle')"
      :okText="t('common.okText')"
      :okButtonProps="{ disabled: !password }"
      @ok="handleDeploy"
    >
      <Form layout="vertical">
        <FormItem :label="t('applications.deploy.password')" name="pledgeAmount">
          <InputPassword v-model:value="password" />
        </FormItem>
      </Form>
    </Modal>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { reactive, ref, onMounted } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useRoute, useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useMessage } from '/@/hooks/web/useMessage';
  import InitForm from './components/InitForm.vue';
  import StakeForm from './components/StakeForm.vue';
  import DeployForm from './components/DeployForm.vue';
  import DeployInfo from '../application/components/DeployInfo.vue';
  import { GetDeployInfo, DeployTheGraph } from '/@wails/go/app/Deploy';
  import { useSettingStore } from '/@/store/modules/setting';
  import {
    createPolkadotApi,
    createKeyPair,
    applyResourceOrder,
    handleTxResults,
  } from '/@/utils/polkadotUtil';
  import { Steps, Step, Modal, Form, FormItem, InputPassword } from 'ant-design-vue';

  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const settingStore = useSettingStore();
  const router = useRouter();
  const { params } = useRoute();
  const applicationId = Number(params.id);

  const previewModalVisible = ref(false);
  const previewModalBodyStyle = reactive({
    height: '360px',
    overflow: 'auto',
  });

  // steps
  const steps = [InitForm, StakeForm, DeployForm];
  const currentStep = ref(0);
  const handleNext = () => {
    if (currentStep.value === steps.length - 1) {
      previewModalVisible.value = true;
    } else {
      currentStep.value++;
    }
  };

  // deployInfo
  const deployInfo = reactive<{
    initialization: Recordable;
    staking: Recordable;
    deployment: Recordable;
  }>({
    initialization: {},
    staking: {},
    deployment: {},
  });

  // Get saved deployInfo from API
  onMounted(async () => {
    const { data } = await GetDeployInfo(applicationId);
    if (data) Object.assign(deployInfo, data);

    // Run settingStore actions
    settingStore.getWalletInfoAction();
    settingStore.getConfigAction();
  });

  // Input password when confirm deploying
  const passwordModalVisible = ref(false);
  const password = ref('');

  const handleDeploy = async () => {
    const json = settingStore.walletInfo?.addressJson;
    const wsUrl = settingStore.config?.wsUrl;

    if (!json) {
      console.log('Please import your account before deployment.');
      return;
    }

    if (!wsUrl) {
      console.log('Please config the wsUrl in setting.');
      return;
    }

    const polkadotApi = await createPolkadotApi(wsUrl);
    const keyPair = createKeyPair(JSON.parse(json), password.value);

    if (!keyPair) {
      createErrorModal({
        content: t('applications.deploy.passwordError'),
      });
      return;
    }

    const { leaseTerm, publicKey } = deployInfo.initialization;
    const unsubscribe = await applyResourceOrder(
      polkadotApi,
      keyPair,
      { leaseTerm, publicKey },
      handleTxResults({
        txSuccessCb: async (result) => {
          console.log(result);

          // Call deploy API
          await DeployTheGraph(applicationId, {});
          router.push('/applications/' + applicationId);
        },
        txFailedCb: (error) => {
          console.log(error);
          createErrorModal({
            content: t('applications.deploy.deployFailed'),
          });
        },
        unsubscribe: () => unsubscribe(),
      }),
    );
  };
</script>

<style lang="less" scoped>
  :deep(.input-width) {
    width: 50%;
  }
</style>
