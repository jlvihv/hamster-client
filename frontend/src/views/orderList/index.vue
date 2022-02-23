<template>
  <a-spin :tip="'loading'" :spinning="loadLoading" class="spin-style">
    <div class="tableStyle">
      <no-data-table v-if="list.data.length === 0">
        <a-table-column
            title="Order ID"
        />
        <a-table-column
            title="Resource ID"
            data-index="resource_index"
        />
        <a-table-column title="Order Amount"/>
        <a-table-column title="Creation" :width="120"/>
        <a-table-column
            title="Order Status"
        />
        <a-table-column title="Operate" :width="90"></a-table-column>
      </no-data-table>
      <p-table
          v-show="list.data.length != 0"
          :row-key="t => t.index"
          :dataList="list"
          @get-data="getOrderList"
          ref="table-page"
      >
        <a-table-column
            title="Order ID"
            data-index="index"
            :width="72"
        />
        <a-table-column
            title="Resource ID"
            data-index="resource_index"
            :width="90"
        />
        <a-table-column title="Order Amount" data-index="price"/>
        <a-table-column title="Creation" :sorter="(a, b) => a.index - b.index" :width="125">
          <template #default="{ record }">
            <span>{{setTime(record.time.secs.toString().replace(/["|’|,|\“|\”|\，]/g, ""))}}</span>
          </template>
        </a-table-column>
        <a-table-column
            :width="105"
            title="Order Status"
        >
          <template #default="{ record }">
        <span style="vertical-align: middle;" v-if="record.status === 'Pending'">
          <img src="../../assets/doing.png">
          Processing</span>
            <span style="vertical-align: middle;" v-if="record.status === 'Finished'">
          <img src="../../assets/finished.png">
          Completed </span>
            <span style="vertical-align: middle;" v-if="record.status === 'Canceled'">
          <img src="../../assets/cancel.png">
          Cancelled </span>
          </template>
        </a-table-column>
        <a-table-column title="Operate" :width="95">
          <template #default="{ record }">
            <a style="color: #4850FF" @click="showCancelOrderModel(record.index)">{{ record.status === 'Pending' ? 'Cancel Order' : ''}}</a>
          </template>
        </a-table-column>
      </p-table>
    </div>
  </a-spin>
  <transaction-modal ref="cancelOrderRef" @signedtransaction="cancelOrder" @close="cancelLoading"/>
</template>

<script>
import {getCurrentInstance, onMounted, reactive, ref, toRefs} from "vue";
import NoDataTable from "../../components/table/NoDataTable";
import PTable from "../../components/table/PTable";
import api from "../../api";
import TransactionModal from "../../components/model/transactionModal"

export default {
  name: "index",
  components: {
    NoDataTable,
    PTable,
    TransactionModal
  },
  setup() {
    const cancelOrderRef = ref();
    const { proxy } = getCurrentInstance();
    const state = reactive({
      loadLoading: false,
      cancelIndex: '',
      list: {
        total: 0,
        data: [],
        current: 1
      },
    });
    onMounted(() => {
      getOrderList();
    })
    const setTime = (str) => {
      let n = parseInt(str) * 1000;
      let D = new Date(n);
      let year = D.getFullYear();

      let month = D.getMonth() + 1;
      month = month < 10 ? ('0' + month) : month;

      let day = D.getDate();
      day = day < 10 ? ('0' + day) : day;

      let hours = D.getHours();
      hours = hours < 10 ? ('0' + hours) : hours;

      let minutes = D.getMinutes();
      minutes = minutes < 10 ? ('0' + minutes) : minutes;

      let now_time = year + '-' + month + '-' + day + ' ' + hours + ':' + minutes;
      return now_time;
    }
    const showCancelOrderModel = async (index) => {
      state.loadLoading = true;
      state.cancelIndex = index;
      let API = await api;
      let transaction = API.tx.resourceOrder.cancelOrder(index);
      cancelOrderRef.value.openModal(transaction.method.toHex());
    }
    const cancelOrder = async (krp) => {
      let API = await api;
      API.tx.resourceOrder.cancelOrder(state.cancelIndex).signAndSend(krp,{nonce: -1},({ status, events, dispatchError }) => {
        if (status.isInBlock) {
          if (dispatchError) {
            if (dispatchError.isModule) {
              const decoded = API.registry.findMetaError(dispatchError.asModule);
              const { docs, name, section } = decoded;
              console.log(`${section}.${name}: ${docs.join(' ')}`);
              proxy.$message.error("Order cancellation failed:" + docs);
              state.loadLoading = false;
            } else {
              proxy.$message.error("Order cancellation failed:" + dispatchError.toString());
              state.loadLoading = false;
            }
          } else {
            proxy.$message.success("Order cancelled successfully");
            getOrderList(() => state.loadLoading = false);
          }
        }
      }).catch(() => {
        state.loadLoading = false;
        proxy.$message.error("Order cancellation failed");
      })
    }
    const getOrderList = async (lodding) => {
      let addressData =  await window.go.app.Wallet.GetWalletInfo();
      api.then(t => t.query.resourceOrder.userOrders(addressData.address))
          .then(data => {
            return new Promise(function (resolve) {
              resolve(data.toHuman())
            })
          })
          .then(data => api.then(t => t.query.resourceOrder.resourceOrders.multi(data)))
          .then(data => {
            return new Promise(function (resolve) {
              resolve(data.map(t => t.toHuman()))
            })
          }).then(data => {
        let list = []
        for (let i in data) {
          list.push(data[i])
        }
        state.list.data = list.reverse();
        console.log(state.list.data);
        if (lodding) {
          lodding()
        }
      })
    }
    const cancelLoading = () => {
      state.loadLoading = false;
    }
    return {
      ...toRefs(state),
      getOrderList,
      setTime,
      cancelOrder,
      showCancelOrderModel,
      cancelOrderRef,
      cancelLoading
    }
  }
}
</script>

<style scoped>

</style>