<template>
  <div class="flex justify-between mb-[20px]">
    <div>
      <SvgIcon v-if="showBack" @click="onBack" class="cursor-pointer" color="#858B92" size="20" name="left" />
    </div>
    <div>
      <SvgIcon class="cursor-pointer" color="#858B92" size="20" name="people" />
      <SvgIcon class="cursor-pointer ml-[20px]" color="#858B92" size="20" name="setting" />
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { toRefs } from 'vue';
  import { SvgIcon } from '/@/components/Icon';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const props = defineProps({
    showBack: Boolean,
  });
  const { showBack } = toRefs(props);

  const onBack = async () => {
    router.push({ path: '/applications/index' });
  };
</script>