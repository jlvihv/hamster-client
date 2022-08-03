<template> Guide Page </template>

<script lang="ts" setup>
  import { onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import { useSettingStore } from '/@/store/modules/setting';

  const router = useRouter();
  const settingStore = useSettingStore();

  onMounted(() => {
    // Redirect to applidation index if wallet binded
    if (settingStore.walletInfo) {
      router.push('/applications/index');
    }
  });
</script>
