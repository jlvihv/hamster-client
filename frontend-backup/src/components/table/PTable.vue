<template>
  <div class="pt-table">
    <a-table
        :data-source="dataList.data"
        :pagination="pagination"
        :rowKey="rowKey"
    >
      <slot></slot>
    </a-table>
  </div>
</template>

<script>
import {watch} from "vue";

export default {
  name: "p-table",
  emits:['change'],
  props: {
    pagination: {
      type: Boolean,
      default: false
    },
    rowKey: {
      type: Function,
      default: t => t
    },
    pageSize: {
      type: Number,
      default: 10
    },
    dataList: {
      type: Object
    },
    queryFunction: {
      type: Function
    },
    params: {
      type: Object
    },
  },
  setup(props) {
    watch(() => props.dataList, (val) => {
      this.list = val.data;
    })
  }
};
</script>
<style lang="scss">
$table-head-color: #b6bac5;
$table-background-color: #fafbfd;
$line-color: #f7f8fa;
.pt-table {
  .custom-pagination {
    display: flex;
    justify-content: center;
    margin-top: 8px;
    .custom-pagination-all {
      display: flex;
      align-items: center;
      .total-style {
        color: #697581;
        width: 120px;
        height: 41px;
        border: 1px solid #edeff1;
        margin-right: 20px;
        border-radius: 4px;
        padding: 12px 0px;
        display: flex;
        justify-content: center;
      }
      .jump-page {
        margin-left: 12px;
        display: flex;
        align-items: center;
        .go-to {
          margin-right: 12px;
        }
        .jump-input {
          border: 1px solid #edeff1;
          height: 41px;
          padding: 12px 0px;
          width: 120px;
          text-align: center;
        }
        .page {
          margin-left: 12px;
        }
      }
    }
  }
}
.ant-table-thead > tr > th {
  font-size: 12px;
  line-height: 17px;
  color: $table-head-color;
  background: #f9faff;
  border-bottom: 0px;
}

.ant-table-tbody > tr > td {
  font-weight: 600;
  font-size: 12px;
  line-height: 17px;
  color: #697581;
  border-bottom: 1px solid $line-color;
}
.ant-pagination-prev {
  width: 120px;
  height: 41px;
  .ant-pagination-item-link {
    color: #cacdd3;
    background-color: #f0f3ff;
    border: 0px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
.ant-pagination-next {
  width: 120px;
  height: 41px;
  .ant-pagination-item-link {
    color: #cacdd3;
    background-color: #f0f3ff;
    border: 0px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
.ant-pagination-next:hover .ant-pagination-item-link {
  color: #ffffff;
  background: #4850ff;
  border-radius: 5px;
}
.ant-pagination-prev:hover .ant-pagination-item-link {
  color: #ffffff;
  background: #4850ff;
  border-radius: 5px;
}
</style>
