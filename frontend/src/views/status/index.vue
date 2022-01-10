<template>
  <a-spin :tip="'Loading'" :spinning="loadLoading" class="spin-style">
    <div class="status-container">
      <div class="tableStyle">
        <no-data-table v-if="list.data.length === 0">
          <a-table-column
              title="Protocol"
              data-index="Protocol"
              :width="100"/>
          <a-table-column title="Listen Address" :width="200" data-index="ListenAddress" />
          <a-table-column
              title="Target Address"
              :width="200"
              data-index="TargetAddress"
          />
          <a-table-column
              title="SSH Command"
              :width="200"
          />
          <a-table-column
              title="Status"
              :width="100"
              data-index="Status"
          />
          <a-table-column title="Operate" :width="100"/>
        </no-data-table>
        <p-table
            v-show="list.data.length != 0"
            :row-key="t => t.index"
            :dataList="list"
            @get-data="getList"
            ref="table-page"
        >
          <a-table-column
              title="Protocol"
              :ellipsis="true"
              data-index="Protocol"
              :width="100"/>
          <a-table-column title="Listen Address" :width="200" :ellipsis="true" data-index="ListenAddress" />
          <a-table-column
              title="Target Address"
              :width="200"
              :ellipsis="true"
              data-index="TargetAddress"
          />
          <a-table-column
              title="SSH Command"
              :width="150"
              :ellipsis="true"
          >
            <template #default="{ record }">
              <div style="display: flex;align-items: center">
                <a-tooltip placement="topLeft" :title="'ssh -i ~/.ssh/id_rsa root@127.0.0.1 -p ' + record.ListenAddress.slice(record.ListenAddress.lastIndexOf('/') + 1,record.ListenAddress.length)">
                  <span>{{strSlice('ssh -i ~/.ssh/id_rsa root@127.0.0.1 -p ' + record.ListenAddress.slice(record.ListenAddress.lastIndexOf('/') + 1,record.ListenAddress.length))}}</span>
                </a-tooltip>
                <img src="../../assets/Copy.png" style="cursor: pointer;width: 16px;height: 16px" @click="success" v-clipboard="'ssh -i ~/.ssh/id_rsa root@127.0.0.1 -p ' + record.ListenAddress.slice(record.ListenAddress.lastIndexOf('/') + 1,record.ListenAddress.length)" >
<!--                <a style="color: #4850FF" @click="success" v-clipboard="'ssh -i ~/.ssh/id_rsa root@127.0.0.1 -p ' + record.ListenAddress.slice(record.ListenAddress.lastIndexOf('/') + 1,record.ListenAddress.length)">复制</a>-->
              </div>
            </template>
          </a-table-column>
          <a-table-column
              title="Status"
              :width="100"
              data-index="Status"
          >
            <template #default="{ record }">
              <span>{{record.Status == true ? 'connected' : 'disconnected'}}</span>
            </template>
          </a-table-column>
          <a-table-column title="Operate" :width="100">
            <template #default="{ record }">
              <a style="color: #4850FF" @click="closeLink(record)">Disconnect</a>
            </template>
          </a-table-column>
        </p-table>
      </div>
    </div>
  </a-spin>
</template>

<script>
import NoDataTable from "../../components/table/NoDataTable";
import PTable from "../../components/table/PTable";
import {getCurrentInstance, onMounted, reactive, toRefs} from "vue";
import * as Wails from "@wailsapp/runtime";
export default {
  name: "index",
  components: {
    NoDataTable,
    PTable
  },
  setup() {
    const { proxy } = getCurrentInstance();
    const state = reactive({
      loadLoading: false,
      list: {
        total: 0,
        data: [],
        current: 1
      },
    })
    onMounted(() => {
      getLinkStatus();
      Wails.Events.On("Links",t=>{
        state.list.data = []
        if (t) {
          state.list.data = t
        }
      })
    })
    const getList = () => {

    }
    //get connection status
    const getLinkStatus = () => {
      state.loadLoading = true
      window.backend.P2p.GetLinkStatus().then((res)=> {
        state.list.data = []
        if (res) {
          state.list.data = res
        }
        state.loadLoading = false
      }).catch((err) => {
        proxy.$message.error(err)
        state.loadLoading = false
      })
    }
    //DISCONNECT
    const closeLink = (record) => {
      window.backend.P2p.CloseLink(record.TargetAddress).then(() => {
        getLinkStatus()
        proxy.$message.success("Disconnected successfully")
      }).catch((err) => {
        proxy.$message.error("Disconnect failed" + err)
      })
    }
    const strSlice = (data) => {
      if (data.length > 16) {
        return data.slice(0,16) + "..."
      }
      return data
    }
    const success = () => {
      proxy.$message.success("Copy successfully")
    }
    return {
      ...toRefs(state),
      getLinkStatus,
      getList,
      closeLink,
      strSlice,
      success
    }
  }
}
</script>

<style lang="scss" scoped>
.status-container {
  padding: 12px;
  background: white;
  border-radius: 8px;
  display: flex;
  justify-content: center;
  .tableStyle {
    width: 100%;
    background: white;
    border-radius: 8px;
    //height: 100%;
  }
}
.copy-btn {
  font-weight: 600;
  font-size: 12px;
  line-height: 17px;
  color: #4850FF;
  padding: 0px;
;
}
</style>
