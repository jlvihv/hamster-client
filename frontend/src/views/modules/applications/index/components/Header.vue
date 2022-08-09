<template>
  <div class="flex justify-between mb-[20px] relative">
    <div>
      <SvgIcon
        v-if="showBack"
        @click="onBack"
        class="cursor-pointer"
        color="#858B92"
        size="20"
        name="left"
      />
    </div>
    <div>
      <SvgIcon
        class="cursor-pointer"
        color="#858B92"
        size="20"
        name="people"
        @click="
          showPeople = true;
          showSetting = false;
        "
      />
      <SvgIcon
        class="cursor-pointer ml-[20px]"
        color="#858B92"
        size="20"
        name="setting"
        @click="
          showSetting = true;
          showPeople = false;
        "
      />
    </div>
    <div v-if="showPeople" @mouseleave="showPeople = false" class="pop-div">
      <div class="top-div right-[50px]"></div>
      <div class="border-box">
        <div class="title-div">
          <SvgIcon size="16" color="#63A0FA" name="address" />
          {{ t('applications.index.address') }}
        </div>
        <div class="text-[12px] mb-[20px] whitespace-normal break-all"
          >5EKQCEm834AMHZ2CvdWeQUP5QrzBJzYa5PQrzBJzYa5PnhByhitgK9RS5o</div
        >
        <div class="title-div">
          <SvgIcon size="16" color="#63A0FA" name="balance" />
          {{ t('applications.index.balance') }}
        </div>
        <div class="text-[12px] mb-[20px]">300000.00 Unit</div>
        <div class="button-div">
          <Button type="primary">{{ t('applications.index.changeWallet') }}</Button>
        </div>
      </div>
    </div>
    <div v-if="showSetting" @mouseleave="onMouseLeave($event)" class="pop-div">
      <div class="top-div right-[10px]"></div>
      <div class="border-box">
        <div class="title-div">
          <SvgIcon size="16" color="#63A0FA" name="ws" />
          {{ t('applications.index.wsUrl') }}
        </div>
        <div>
          <Select
            class="w-full"
            :allowClear="true"
            :placeholder="t('applications.index.wsUrlPlaceholder')"
            :options="urlOptions"
          />
        </div>
        <div class="button-div">
          <Button type="primary">{{ t('common.saveText') }}</Button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { toRefs, reactive, ref } from 'vue';
  import { SvgIcon } from '/@/components/Icon';
  import { useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Button, Select } from 'ant-design-vue';

  const { t } = useI18n();
  const router = useRouter();
  const props = defineProps({
    showBack: Boolean,
  });
  const { showBack } = toRefs(props);
  const urlOptions = reactive([
    { label: '5EKQCEm834AMHZ2CvdWeQUP5QrzBJzYa5PQrzBJzYa5PnhByhitgK9RS5o', value: '1' },
  ]);
  const showSetting = ref(false);
  const showPeople = ref(false);

  const onBack = async () => {
    router.push({ path: '/applications/index' });
  };

  async function onMouseLeave(e) {
    var currTargetEl = e.relatedTarget || e.toElement;
    const targetClassName = currTargetEl.className
    if (targetClassName.indexOf('ant-select-dropdown') !== -1) return; 
    showSetting.value = false;
  }
</script>
<style lang="less" scoped>
  .pop-div {
    @apply absolute z-50 top-9;
    width: 250px;
    right: -10px;
  }

  .border-box {
    @apply bg-white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0px 0px 4px 0px rgba(31, 31, 35, 0.2);
  }

  .top-div {
    @apply absolute;
    top: -40px;
    width: 20px;
    height: 40px;
    border-bottom: 10px solid #fff;
    border-left: 10px solid transparent;
    border-right: 10px solid transparent;
  }

  .title-div {
    @apply font-bold;
    margin-bottom: 10px;
  }

  .button-div {
    @apply text-center;
    margin-top: 30px;
    margin-bottom: 30px;
  }

  :deep(.ant-select-single:not(.ant-select-customize-input) .ant-select-selector) {
    height: auto !important;
  }

  :deep(.ant-select-single .ant-select-selector .ant-select-selection-item) {
    line-height: 20px !important;
    padding-top: 5px;
    padding-bottom: 5px;
    white-space: normal;
    word-break: break-all;
  }
</style>
