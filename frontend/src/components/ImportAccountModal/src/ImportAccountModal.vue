<template>
  <template v-if="settingStore.walletInfo?.address">
    <div class="text-primary">{{ settingStore.walletInfo?.address }}</div>
  </template>
  <template v-else>
    <slot name="trigger" :showModal="showModal">
      <Button type="link" size="large" @click="showModal">
        {{ t('importAccount') }}
      </Button>
    </slot>
  </template>
  <Modal
    v-model:visible="modalVisible"
    :class="prefixCls"
    :title="t('importAccount')"
    :ok-text="tt('common.okText')"
    :cancel-text="tt('common.closeText')"
    @ok="handleModalOK"
  >
    <Form layout="vertical" ref="formRef" :model="formData" :rules="formRules">
      <FormItem name="fileList">
        <Upload
          accept=".json"
          v-model:fileList="formData.fileList"
          :maxCount="1"
          :beforeUpload="beforeUpload"
          @remove="handleFileRemove"
        >
          <Button type="dashed">
            {{ t('buttonImportJson') }}
          </Button>
        </Upload>
      </FormItem>
      <template v-if="formData.jsonContent">
        <FormItem>
          <div class="font-bold text-primary">{{
            `${t('name')}: ${formData.jsonContent.meta.name}`
          }}</div>
          <div>{{ formData.jsonContent.address }}</div>
        </FormItem>
        <FormItem name="password" :label="t('upperPassword')">
          <InputPassword
            :placeholder="t('PleaseEnterPassword')"
            v-model:value="formData.password"
          />
        </FormItem>
      </template>
    </Form>
  </Modal>
</template>

<script lang="ts" setup>
  import { ref, reactive, computed, onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { isJSONAndPasswordMatch } from '/@/utils/polkadotUtil';
  import { useSettingStore } from '/@/store/modules/setting';
  import { Button, Modal, Upload, Form, FormItem, InputPassword } from 'ant-design-vue';

  const { t } = useI18n('components.importAccountModal');
  const { t: tt } = useI18n();
  const { prefixCls } = useDesign('import-account-modal');
  const { createErrorModal } = useMessage();
  const settingStore = useSettingStore();

  // get Wallet Info on mounted
  onMounted(() => settingStore.getWalletInfoAction());

  const modalVisible = ref<boolean>(false);
  const formRef = ref();

  const formData = reactive<{
    fileList?: Recordable[];
    password?: string;
    jsonContent?: Recordable;
    jsonString?: string;
  }>({});

  const formRules = computed(() => ({
    fileList: [{ message: t('uploadFile'), trigger: 'change', required: true }],
    password: [{ message: t('PleaseEnterPassword'), trigger: 'change', required: true }],
  }));

  const showModal = () => {
    modalVisible.value = true;
  };

  const handleModalOK = async () => {
    await formRef.value?.validate();

    const { jsonContent, password } = formData;

    if (!jsonContent || !jsonContent.address || !password) return;

    // Check password
    const isMatch = isJSONAndPasswordMatch(jsonContent, password);

    if (isMatch) {
      // call api to save wallet info
      await settingStore.saveWalletAction(jsonContent.address, jsonString);
      modalVisible.value = false;
    } else {
      createErrorModal({ content: t('passwordError') });
    }
  };

  const beforeUpload = (file: File, fileList: Recordable[]) =>
    new Promise((resolve) => {
      const reader = new FileReader();

      reader.readAsText(file);
      reader.onload = (e: any) => {
        const fileContent = e.target?.result;
        formData.jsonString = fileContent;
        let json: Recordable | undefined;

        try {
          json = JSON.parse(fileContent);
        } catch (e: any) {
          // Nothing
        }

        if (!json || !json.address || !json.encoded || !json.encoding || !json.meta) {
          createErrorModal({ content: t('correctFormatJson') });

          fileList.pop();
          handleFileRemove();
          resolve(false);
          return;
        }

        // Assign json content,
        // false means preventing upload
        formData.jsonContent = json;
        resolve(false);
      };
    });

  // Reset fromData when json file removed
  const handleFileRemove = () => {
    formData.fileList = undefined;
    formData.jsonContent = undefined;
    formData.password = undefined;
  };
</script>

<style lang="less">
  @prefix-cls: ~'@{namespace}-import-account-modal';

  .@{prefix-cls} {
    .ant-upload {
      @apply w-full;
    }

    .ant-btn-dashed {
      @apply w-full h-10;
    }

    .ant-form-item {
      @apply mb-3;
    }
  }
</style>
