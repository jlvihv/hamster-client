<template>
  <Form layout="vertical" :model="formData" :rules="formRules" ref="formRef">
    <FormItem :label="t('applications.deploy.leaseTerm')" name="leaseTerm">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value.number="formData.leaseTerm"
        :placeholder="t('applications.deploy.leaseTermPlaceholder')"
      />
    </FormItem>
    <FormItem :label="t('applications.deploy.publicKey')" name="publicKey">
      <Textarea
        class="input-width"
        v-model:value="formData.publicKey"
        :placeholder="t('applications.deploy.publicKeyPlaceholder')"
      />
    </FormItem>
    <FormItem :label="t('applications.deploy.importAccount')" name="accountMnemonic">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value="formData.accountMnemonic"
        :placeholder="t('applications.deploy.importAccountPlaceholder')"
      />
    </FormItem>
    <FormItem class="text-center">
      <Button @click="goBack" class="w-20 mt-6 ml-4">{{ t('common.cancelText') }}</Button>
      <Button class="w-20 mt-6 ml-4" type="primary" @click="handleSubmit">
        {{ t('common.nextText') }}
      </Button>
    </FormItem>
  </Form>
</template>

<script lang="ts" setup>
  import { ref, toRaw, toRefs, computed } from 'vue';
  import { useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { SaveDeployInfo } from '/@wails/go/app/Deploy';
  import { createRule } from '/@/utils/formUtil';
  import { Form, FormItem, Button, Input, Textarea } from 'ant-design-vue';

  // defines
  const props = defineProps({
    applicationId: Number,
    deployInfo: Object as PropType<Recordable>,
  });
  const { applicationId, deployInfo } = toRefs(props);

  const emits = defineEmits(['update:deployInfo', 'submited']);

  const { t } = useI18n();
  const { createErrorModal } = useMessage();
  const router = useRouter();

  const goBack = () => router.push('/applications/' + applicationId.value);

  const formRef = ref();
  const formData: {
    leaseTerm?: string;
    publicKey?: string;
    accountMnemonic?: string;
  } = deployInfo.value.initialization;

  const formRules = computed(() => ({
    leaseTerm: [createRule(t('applications.deploy.leaseTermPlaceholder'))],
    publicKey: [createRule(t('applications.deploy.publicKeyPlaceholder'))],
    accountMnemonic: [createRule(t('applications.deploy.importAccountPlaceholder'))],
  }));

  const handleSubmit = async () => {
    await formRef.value?.validate();

    try {
      await SaveDeployInfo(applicationId.value, JSON.stringify(toRaw(deployInfo.value)));
    } catch (error: any) {
      createErrorModal({
        title: t('common.errorTip'),
        content: t('applications.deploy.saveFailed'),
      });

      console.log(error);
      return;
    }

    emits('submited', formData);
  };
</script>
