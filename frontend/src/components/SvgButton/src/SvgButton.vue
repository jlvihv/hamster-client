<template>
  <Button
    v-bind="$attrs"
    :loading="loading"
    @click="onClick"
    class="!border-none !bg-transparent !p-0"
    :style="btnWH"
  >
    <template #icon>
      <SvgIcon :class="iconClass" :size="size" :name="icon" />
    </template>
  </Button>
</template>
<script lang="ts" setup>
  import { toRefs, reactive, ref } from 'vue';
  import { propTypes } from '/@/utils/propTypes';
  import { SvgIcon } from '/@/components/Icon';
  import { Button } from 'ant-design-vue';

  const props = defineProps({
    iconClass: propTypes.string.def(''),
    icon: propTypes.string.def(''),
    size: propTypes.string.def(''),
  });
  const { size } = toRefs(props);

  const emits = defineEmits(['click']);
  const loading = ref(false);
  async function onClick() {
    //The callback method needs to be called to modify the loading value
    loading.value = true;

    try {
      emits('click', () => {
        loading.value = false;
      });
    } catch (error: any) {
      loading.value = false;
    }
  }

  const btnWH = reactive({ height: size.value + 'px', width: size.value + 'px' });
</script>
