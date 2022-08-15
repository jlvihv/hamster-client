<template>
  <Button v-bind="$attrs" :loading="loading" @click="onClick" class="!border-none !bg-transparent" :style="btnWH">
    <template #icon>
      <SvgIcon 
        :class="class"
        :size="size"
        :name="icon"
      />
    </template>
  </Button>
</template>
<script lang="ts" setup>
  import { toRefs, reactive, ref } from 'vue';
  import { propTypes } from '/@/utils/propTypes';
  import { SvgIcon } from '/@/components/Icon';
  import { Button } from 'ant-design-vue';

  const props =  defineProps({
    class: propTypes.string.def(''),
    icon: propTypes.string.def(''),
    size: propTypes.string.def(''),
  });
  const { size } = toRefs(props);

  const emits = defineEmits(['click']);
  const loading = ref(false);
  async function onClick() {  //The callback method needs to be called to modify the loading value
    loading.value = true;
    emits('click', () => {
      loading.value = false;
    })
    
  }

  const btnWH = reactive({ height: size.value + 'px', width: size.value + 'px' });
</script>
