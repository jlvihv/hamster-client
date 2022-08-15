<template>
  <Button @click="onClick" class="!border-none !bg-transparent" :style="btnWH" :type="type" :shape="shape" :target="target" :loading="loading" :htmlType="htmlType" :ghost="ghost" :disabled="disabled" :danger="danger" :block="block">
    <template #icon>
      <SvgIcon 
        :class="className"
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

  const props = defineProps({
    block: propTypes.bool.def(false),
    danger: propTypes.bool.def(false),
    disabled: propTypes.bool.def(false),
    ghost: propTypes.bool.def(false),
    class: propTypes.string.def(''),
    htmlType: propTypes.string.def('button'),
    icon: propTypes.string.def(''),
    shape: propTypes.string.def('circle'),
    size: propTypes.string.def('middle'),
    target: propTypes.string.def(''),
    type: propTypes.string.def('default'),
  });
  const { block, danger, disabled, ghost, htmlType, icon, shape, size, target, type, class:className } = toRefs(props);

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
