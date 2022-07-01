<template>
  <div class="no-data-table">
    <a-table :data-source="list">
      <slot></slot>
    </a-table>
    <div
      style="min-height: 600px;display: flex;align-items: center;justify-content: center"
    >
      <div>
        <img :src="image" />
      </div>
    </div>
  </div>
</template>

<script>
import {reactive, toRefs} from "vue";

export default {
  name: "NoDataTable",
  setup() {
    const state = reactive({
      list: [],
      image:require("../../assets/empty.png")
    })
    return {
      ...toRefs(state)
    }
  }
};
</script>

<style lang="scss">
$table-head-color: #b6bac5;
.no-data-table {
  height: 100%;
  .ant-table-thead > tr > th {
    color: $table-head-color;
    border: 0px;
    background: #f9faff;
    font-weight: 600;
  }
  .ant-table-placeholder {
    display: none !important;
  }
}
</style>
