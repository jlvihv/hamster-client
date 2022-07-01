<template>
<div class="model-container">
  <a-modal
      class="login-model"
      :visible="visible"
      :width="480"
      :centered="true"
      :footer="null"
      :maskClosable="false"
      :closable="false"
  >
      <span>
      </span>
    <a-spin :tip="tip" :spinning="loading">
      <div class="modal-header">
        <div :class="titleBigStyle ? 'title-font' : 'title-font-style'">{{title}}</div>
        <div v-show="isShowCancel">
          <img class="close" @click="close" src="../../assets/Close.png" />
        </div>
      </div>
      <div class="modal-content">
        <slot></slot>
      </div>
    </a-spin>
  </a-modal>
</div>
</template>

<script>
import {defineComponent} from "vue";

export default defineComponent({
  name: "index",
  props: {
    isShowCancel: {
      type: Boolean,
      default: true
    },
    cancelFunction: {
      type: Function,
      default: function () {}
    },
    title: {
      type: String,
      default: ""
    },
    tip: {
      type: String,
      default: ""
    },
    loading: {
      type: Boolean,
      default: false,
    },
    visible: {
      type: Boolean,
      default: false
    },
    titleBigStyle: {
      type: Boolean,
      default: true
    },
    close: {
      type: Function,
      default: function (){}
    }
  }
})
</script>

<style lang="scss">
.ant-modal-body {
  padding: 0px !important;
}
.close {
  cursor: pointer;
}
.modal-header {
  display: flex;
  justify-content: space-between;
  padding:32px 32px 0px 32px;
  .title-font {
    color: #000000;
    font-weight: 600;
    font-size: 32px;
    line-height: 45px;
  }
  .title-font-style {
    color: #000000;
    font-weight: 600;
    font-size: 20px;
    line-height: 28px;
  }
}
.modal-content {
  padding: 0px 32px 32px 32px;
}

</style>
