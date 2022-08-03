<template>
  <div>{{ status }}</div>
</template>

<script lang="ts" setup>
  import { ExceptionEnum } from '/@/enums/exceptionEnum';

  defineProps({
    status: {
      type: Number as PropType<number>,
      default: ExceptionEnum.PAGE_NOT_FOUND,
    },

    title: {
      type: String as PropType<string>,
      default: '',
    },

    subTitle: {
      type: String as PropType<string>,
      default: '',
    },

    full: {
      type: Boolean as PropType<boolean>,
      default: false,
    },
  });
</script>
