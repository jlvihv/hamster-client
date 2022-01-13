<template>
  <div>
    <a @click="showModal">add user</a>
    <a-modal v-model:visible="visible" title="ADD" @ok="handleOk">
        <a-textarea allow-clear v-model:value="value" placeholder="Please enter user public key" />
    </a-modal>
  </div>
</template>
<script>
import {defineComponent, ref} from 'vue';
import { message } from 'ant-design-vue';

export default defineComponent({
  props:{
    id: Number
  },
  setup(props) {
    let value = ref("")
    const visible = ref(false);

    const showModal = () => {
      visible.value = true;
    };

    const handleOk = () => {
      console.info(window.backend)
      let param = { user:value.value,ID:props.id }
      window.backend.WailsApi.UseResource(param).then(() => {
        message.success('SUCCESSFUL');
        visible.value = false;
      });
    };

    return {
      value,
      visible,
      showModal,
      handleOk,
    };
  },
});
</script>