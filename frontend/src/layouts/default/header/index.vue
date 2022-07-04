<template>
  <div class="float-right">
    <Button type="link" size="large" @click="showModal">import account</Button>
    <Modal
      v-model:visible="visible"
      title="Import json"
      ok-text="OK"
      cancel-text="Cancel"
      @ok="handleOk"
    >
      <Upload
        v-model:file-list="fileList"
        name="file"
        :multiple="true"
        action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
        @change="handleChange"
      >
        <Button type="dashed"> click to import json </Button>
      </Upload>
      <div v-if="contentvisible">
        <Descriptions>
          <DescriptionsItem label="NAME:">user</DescriptionsItem>
        </Descriptions>
        <Form layout="vertical">
          <FormItem label="PASSWORD:">
            <Input placeholder="Please enter password" />
          </FormItem>
        </Form>
      </div>
    </Modal>
  </div>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import {
    Button,
    Modal,
    Upload,
    Form,
    FormItem,
    Input,
    Descriptions,
    DescriptionsItem,
  } from 'ant-design-vue';
  import { message } from 'ant-design-vue';
  const visible = ref<boolean>(false);
  const contentvisible = ref<boolean>(false);
  const showModal = () => {
    visible.value = true;
  };

  const handleOk = (e: MouseEvent) => {
    console.log(e);
    visible.value = false;
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

  const handleChange = (info: FileInfo) => {
    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList);
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

  const fileList = ref([]);
</script>
<style lang="less" scoped>
  :deep(.ant-upload) {
    @apply w-full;
  }

  :deep(.ant-btn-dashed) {
    @apply w-full;
  }

  .ant-form-item {
    @apply mb-0;
  }

  :deep(.ant-modal-footer) {
    @apply text-center;
    text-align: center !important;
  }
</style>
