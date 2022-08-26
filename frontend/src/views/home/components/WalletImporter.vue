<template>
  <div class="text-[#222222] text-[40px] font-bold">{{ t('home.importWallet') }}</div>
  <div class="mt-[40px] text-center" v-if="!formData.fileList?.length">
    <div class="flex items-center justify-center my-[40px]">
      <component
        :is="isWails && isMacOS ? UploadDragger : Upload"
        name="file"
        accept=".json"
        class="upload-dnd-zone !bg-transparent"
        v-model:fileList="formData.fileList"
        :maxCount="1"
        :beforeUpload="beforeUpload"
        :openFileDialogOnClick="false"
      >
        <div class="text-center px-3 cursor-pointer" v-if="isWails && isMacOS">
          <img :src="walletImage" class="inline-block w-[200px]" />
          <div v-if="isWails && isMacOS" class="mt-3 text-gray-500">{{
            t('home.dragAccountToUpload')
          }}</div>
        </div>
        <img :src="walletImage" class="inline-block w-[200px]" v-else />
      </component>
    </div>
    <Upload
      name="file"
      accept=".json"
      v-model:fileList="formData.fileList"
      :maxCount="1"
      :beforeUpload="beforeUpload"
    >
      <SvgIcon class="text-primary cursor-pointer" size="56" name="importWallet" />
    </Upload>
  </div>
  <div class="mt-[40px] text-center" v-else>
    <Form
      layout="vertical"
      ref="formRef"
      :model="formData"
      :rules="formRules"
      v-bind="$attrs"
      class="!mb-[40px] text-left"
      :class="prefixCls"
    >
      <FormItem class="my-[20px]" name="file">
        <Upload v-model:fileList="formData.fileList" @remove="handleFileRemove" />
      </FormItem>
      <FormItem class="my-[20px]">
        <div class="font-bold">{{ t('home.name') }}: {{ formData.jsonContent.meta.name }}</div>
        <div>{{ formData.jsonContent.address }}</div>
      </FormItem>
      <FormItem name="password">
        <InputPassword :placeholder="t('home.pwd')" v-model:value="formData.password" />
      </FormItem>
    </Form>
    <SvgButton @click="handleSubmit" iconClass="text-primary" size="56" icon="next" />
  </div>
</template>

<script lang="ts" setup>
  import { ref, reactive, computed, onMounted } from 'vue';
  import { SvgIcon } from '/@/components/Icon';
  import { SvgButton } from '/@/components/SvgButton';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { createKeyPair } from '/@/utils/polkadotUtil';
  import { useSettingStore } from '/@/store/modules/setting';
  import { isWails, isMacOS } from '/@/utils/is';
  import walletImage from '/@/assets/images/wallet.png';
  import { Upload, UploadDragger, Form, FormItem, InputPassword } from 'ant-design-vue';

  const emits = defineEmits(['submit']);

  const { t } = useI18n();
  const { prefixCls } = useDesign('import-wallet');
  const { createErrorModal } = useMessage();
  const settingStore = useSettingStore();

  // get Wallet Info on mounted
  onMounted(() => settingStore.getWalletInfoAction());

  // Form
  const formRef = ref();

  const formData = reactive<{
    fileList?: Recordable[];
    password?: string;
    jsonContent?: Recordable;
    jsonString?: string;
  }>({});

  const formRules = computed(() => ({
    password: [{ message: t('home.pleaseEnterPassword'), trigger: 'change', required: true }],
  }));

  const handleSubmit = async (callback) => {
    await formRef.value?.validate();

    const { jsonContent, jsonString, password } = formData;

    if (!jsonContent || !jsonContent.address || !password) return;

    // Check password
    const isMatch = createKeyPair(jsonContent, password);

    if (isMatch) {
      // call api to save wallet info
      await settingStore.saveWalletAction(jsonContent.address, jsonString, password);
      emits('submit', formData);
    } else {
      createErrorModal({ title: t('common.errorTip'), content: t('home.passwordError') });
    }
    callback();
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
          createErrorModal({ title: t('common.errorTip'), content: t('home.correctFormatJson') });

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
  @prefix-cls: ~'@{namespace}-import-wallet';

  .@{prefix-cls} {
    .ant-upload {
      @apply w-full;
    }

    .ant-btn-dashed {
      @apply w-full h-10;
      border-style: solid;
    }

    .ant-form-item {
      @apply block mb-3;
    }

    .ant-input-affix-wrapper {
      @apply !rounded-[8px] !h-[42px];
    }
  }

  .upload-dnd-zone {
    .ant-upload {
      @apply cursor-default;
    }
  }
</style>
