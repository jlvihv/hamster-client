<template>
  <PageWrapper>
    <Header :showBack="true"></Header>
    <div class="text-center text-[30px] font-bold mb-[40px]">{{
      t('applications.new.addTitle')
    }}</div>
    <div class="bg-white rounded-[20px] mx-[40px] mb-[100px] py-[40px] px-[90px]">
      <Form
        :model="formData"
        :rules="formRules"
        ref="formRef"
        :labelCol="{ style: { width: '200px' } }"
      >
        <FormItem :label="t('applications.new.name')" name="nodeEthereumUrl">
          <Input
            :allowClear="true"
            class="input-width"
            v-model:value="formData.name"
            :placeholder="t('applications.new.namePlaceholder')"
          />
        </FormItem>
        <FormItem :label="t('applications.new.leaseTerm')" name="ethereumUrl">
          <Input
            :allowClear="true"
            class="input-width"
            v-model:value="formData.name"
            :placeholder="t('applications.new.leaseTermPlaceholder')"
          />
        </FormItem>
        <FormItem :label="t('applications.new.nodeType')" name="ethereumNetwork">
          <Select
            :allowClear="true"
            class="input-width"
            v-model:value="formData.name"
            :placeholder="t('applications.new.nodeTypePlaceholder')"
            :options="typeOptions"
          />
        </FormItem>
        <FormItem :label="t('applications.new.thegrapIndexer')" name="indexerAddress">
          <Input
            s
            :allowClear="true"
            class="input-width"
            v-model:value="formData.name"
            :placeholder="t('applications.new.thegrapIndexerPlaceholder')"
          />
        </FormItem>
        <FormItem :label="t('applications.new.stakingAmount')" name="indexerAddress">
          <Input
            :allowClear="true"
            class="input-width"
            v-model:value="formData.name"
            :placeholder="t('applications.new.stakingAmountPlaceholder')"
          />
        </FormItem>
        <FormItem class="text-center">
          <Button
            size="large"
            class="w-32 mt-6 ml-4 !rounded-[30px]"
            type="primary"
            @click="handleSubmit"
          >
            {{ t('common.createText') }}
          </Button>
        </FormItem>
      </Form>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { ref, reactive, computed } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { createRule } from '/@/utils/formUtil';
  import Header from '../index/components/Header.vue';
  import { Form, FormItem, Input, Select, Button } from 'ant-design-vue';

  const { t } = useI18n();

  const typeOptions = reactive([{ label: 'type', value: 'type' }]);
  const formRef = ref();
  const formData = reactive({
    name: '',
  });

  const formRules = computed(() => ({
    name: [createRule(t('applications.new.name'))],
  }));

  const handleSubmit = async () => {
    await formRef.value?.validate();
  };
</script>
<style lang="less" scoped>
  :deep(.ant-form-item-label > label) {
    @apply !h-[42px];
  }

  :deep(.ant-input-affix-wrapper) {
    @apply !rounded-[8px] !h-[42px];
  }

  :deep(.ant-select-selector) {
    @apply !rounded-[8px] !h-[42px];
  }

  :deep(.ant-select-selection-item) {
    @apply !leading-[40px];
  }
</style>
