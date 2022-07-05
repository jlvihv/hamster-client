<template>
  <div class="float-right">
    <Button type="link" size="large" @click="showModal">
      {{ t('routes.header.importAccount') }}
    </Button>
    <Modal
      v-model:visible="visible"
      :title="t('routes.header.importJson')"
      :ok-text="t('routes.header.ok')"
      :cancel-text="t('routes.header.cancle')"
      @ok="handleOk"
    >
      <Form layout="vertical" ref="formRef" :model="formData" :rules="formRules">
        <FormItem :name="t('routes.header.file')">
          <Upload
            accept=".json"
            v-model:fileList="formData.fileList"
            name="t('routes.header.file')"
            @change="handleChange"
            :beforeUpload="beforeUpload"
          >
            <Button type="dashed"> {{ t('routes.header.buttonImportJson') }} </Button>
          </Upload>
        </FormItem>
        <FormItem v-if="contentvisible">
          <Descriptions>
            <DescriptionsItem :label="t('routes.header.name')">
              {{ t('routes.header.user') }}
            </DescriptionsItem>
          </Descriptions>
        </FormItem>
        <FormItem
          :name="t('routes.header.password')"
          v-if="contentvisible"
          :label="t('routes.header.upperPassword')"
        >
          <Input
            :placeholder="t('routes.header.PleaseEnterPassword')"
            v-model:value="formData.password"
          />
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>
<script lang="ts" setup>
  import { ref, reactive, computed } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    Button,
    Modal,
    Upload,
    Form,
    FormItem,
    Input,
    Descriptions,
    DescriptionsItem,
    message,
  } from 'ant-design-vue';

  const { t } = useI18n();
  const visible = ref<boolean>(false);
  const contentvisible = ref<boolean>(false);
  const formRef = ref();

  const formData = reactive<{
    fileList?: string;
    password?: string;
    fileContent?: string;
  }>({});

  const formRules = computed(() => ({
    fileList: [{ message: t('routes.header.uploadFile'), trigger: 'change', required: true }],
    password: [
      { message: t('routes.header.PleaseEnterPassword'), trigger: 'change', required: true },
    ],
  }));

  const showModal = () => {
    visible.value = true;
  };

  const handleOk = async () => {
    await formRef.value?.validate();
    visible.value = false;
    console.log(formData.fileList[0], formData.password, formData.fileContent);
  };

  interface FileItem {
    uid: string;
    name?: string;
    status?: string;
    response?: string;
    url?: string;
  }

  interface FileInfo {
    file: FileItem;
    fileList: FileItem[];
  }

  const beforeUpload = async (file) => {
    const reader = new FileReader();
    reader.readAsText(file);

    reader.onload = (e) => {
      formData.fileContent = JSON.parse(e.target?.result);
    };

    // Prevent upload
    return false;
  };

  const handleChange = async (info: FileInfo) => {
    if (info.file.status !== 'uploading') {
      // console.log(info.file, info.fileList);
      contentvisible.value = true;
    }
    if (info.file.status === 'done') {
      message.success(`${info.file.name} file uploaded successfully`);
      contentvisible.value = true;
    } else if (info.file.status === 'error') {
      message.error(`${info.file.name} file upload failed.`);
      contentvisible.value = true;
    }
  };
</script>

<style lang="less" scoped>
  :deep(.ant-upload) {
    @apply w-full;
  }

  :deep(.ant-btn-dashed) {
    @apply w-full h-10;
  }

  .ant-form-item {
    @apply mb-0;
  }

  :deep(.ant-modal-footer) {
    @apply text-center;
    text-align: center !important;
  }
</style>
