<template>
  <Form layout="vertical" :model="formData" :rules="formRules" ref="formRef">
    <FormItem :label="t('applications.deploy.nodeMainnet')" name="nodeEthereumUrl">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value="formData.nodeEthereumUrl"
        :placeholder="t('applications.deploy.nodeEthereumUrlPlaceholder')"
      />
    </FormItem>
    <FormItem :label="t('applications.deploy.ethereumUrl')" name="ethereumUrl">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value="formData.ethereumUrl"
        :placeholder="t('applications.deploy.ethereumUrlPlaceholder')"
      />
    </FormItem>
    <FormItem :label="t('applications.deploy.ethereumNetwork')" name="ethereumNetwork">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value="formData.ethereumNetwork"
        :placeholder="t('applications.deploy.ethereumNetworkPlaceholder')"
      />
    </FormItem>
    <FormItem :label="t('applications.deploy.indexerAddress')" name="indexerAddress">
      <Input
        :allowClear="true"
        class="input-width"
        v-model:value="formData.indexerAddress"
        :placeholder="t('applications.deploy.indexerAddressPlaceholder')"
      />
    </FormItem>
    <FormItem class="text-right">
      <Button class="ml-4" type="primary" @click="handleSubmit">
        {{ t('common.nextText') }}
      </Button>
    </FormItem>
  </Form>
</template>

<script lang="ts" setup>
  import { computed, ref, toRefs, toRaw } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { SaveDeployInfo } from '/@wails/go/app/Deploy';
  import { createRule } from '/@/utils/formUtil';
  import { Form, FormItem, Button, Input } from 'ant-design-vue';

  // defines
  const props = defineProps({
    applicationId: Number,
    deployInfo: Object as PropType<Recordable>,
  });
  const { applicationId, deployInfo } = toRefs(props);

  const emits = defineEmits(['update:deployInfo', 'submited']);

  const { t } = useI18n();

  const formRef = ref();
  const formData: {
    nodeEthereumUrl?: string;
    ethereumUrl?: string;
    ethereumNetwork?: string;
    indexerAddress?: string;
  } = deployInfo.value.deployment;

  const formRules = computed(() => ({
    nodeEthereumUrl: [createRule(t('applications.deploy.nodeEthereumUrlPlaceholder'))],
    ethereumUrl: [createRule(t('applications.deploy.ethereumUrlPlaceholder'))],
    ethereumNetwork: [createRule(t('applications.deploy.ethereumNetworkPlaceholder'))],
    indexerAddress: [createRule(t('applications.deploy.indexerAddressPlaceholder'))],
  }));

  const handleSubmit = async () => {
    await formRef.value?.validate();

    console.log(formData);

    await SaveDeployInfo(applicationId.value, toRaw(deployInfo.value));
    emits('submited', formData);
  };
</script>
