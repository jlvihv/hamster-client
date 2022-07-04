<template>
  <Modal v-model:visible="visible" title="Basic Modal" @ok="handleOk">
    <p>Some contents...</p>
    <p>Some contents...</p>
    <p>Some contents...</p>
  </Modal>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  // import { useI18n } from '/@/hooks/web/useI18n';
  // import { useMessage } from '/@/hooks/web/useMessage';
  // import { DictCodeEnum } from '/@/enums/dictCodeEnum';
  import { Modal } from 'ant-design-vue';

  // const { t } = useI18n();
  // const { createApiErrorModal, notification } = useMessage();

  // const emits = defineEmits(['operateComplete']);

  const visible = ref(false);
  // const formRef = ref();
  // const formData = reactive({
  //   developerId: 0, //开发者id
  //   auditStatus: DictCodeEnum.IdentificationStatus_CertificateFailed.value, //审核状态
  //   auditFailureReason: '', //审核失败原因
  // });

  // const formRules = computed(() => {
  //   const requiredRule = (message: string) => ({ required: true, trigger: 'change', message });

  //   return {
  //     auditStatus: [requiredRule(t('auditStatus'))],
  //     auditFailureReason: [requiredRule(t('notApprovedReason'))],
  //   };
  // });

  async function handleOk() {
    //   // await formRef.value?.validate();
    //   // formData.developerId = Number(developerId.value);
    //   // const { auditStatus, ...baseParam } = formData;
    //   // const param = reactive({
    //   //   auditStatus: true,
    //   // });
    //   // if (auditStatus === DictCodeEnum.IdentificationStatus_CertificateFailed.value) {
    //   //   param.auditStatus = false;
    //   // }
    //   // try {
    //   //   await auditDeveloperApi({ ...param, ...baseParam });
    //   //   emits('operateComplete', auditStatus);
    //   //   closeModal();
    //   //   notification.success({
    //   //     message: tt('common.operateSucText'),
    //   //     duration: 3,
    //   //   });
    //   // } catch (error) {
    //   //   createApiErrorModal(error);
    //   // }
  }
  // const labelCol = reactive({
  //   style: { width: '90px' },
  // });
</script>

<style lang="less"></style>
