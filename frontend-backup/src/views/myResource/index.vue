<template>
  <a-spin :tip="'loading'" :spinning="loadLoading" class="spin-style">
    <div class="tableStyle">
      <no-data-table v-if="list.data.length === 0">
        <a-table-column
            title="Resource ID"
            data-index="resource_index"
        />
        <a-table-column
            title="Resource Status"
            data-index="memory"
        />
        <a-table-column title="Creation" :width="120"/>
        <a-table-column
            :width="120"
            title="Expire"
            data-index="expireTime"
        />
        <a-table-column title="Operate" :width="90"></a-table-column>
      </no-data-table>
      <p-table
          v-show="list.data.length != 0"
          :row-key="t => t.index"
          :dataList="list"
          @get-data="getResourceList"
          ref="table-page"
      >
        <a-table-column
            :width="70"
            title="Resource ID"
            data-index="resource_index"
        />
        <a-table-column
            title="CPU"
            :width="70"
        >
          <template #default="{ record }">
            <span>{{record.config.cpu}}Core</span>
          </template>
        </a-table-column>
        <a-table-column
            title="RAM"
            :width="70"
        >
          <template #default="{ record }">
            <span>{{record.config.memory}}G</span>
          </template>
        </a-table-column>
        <a-table-column
            title="Resource Status"
        >
          <template #default="{ record }">
            <span>{{record.status}}</span>
          </template>
        </a-table-column>
        <a-table-column title="Creation" :width="110">
          <template #default="{ record }">
            <span>{{record.createTime}}</span>
          </template>
        </a-table-column>
        <a-table-column
            :width="110"
            title="Expire"
            data-index="expireTime"
        >
          <template #default="{ record }">
            <span>{{record.endTime}}</span>
          </template>
        </a-table-column>
        <a-table-column title="Operate" :width="80">
          <template #default="{ record }">
            <a style="color: #4850FF" @click="ok(record)">{{ getStatus(record) }}</a>
          </template>
        </a-table-column>
      </p-table>
    </div>
  </a-spin>
  <renew-modal :title="'renew'" :is-show-cancel="true" ref="renewRef" :visible="visible" :title-big-style="false" :close="renewClose" :loading="loading">
    <div style="margin-top: 32px">
      <a-form ref="renewState" :model="renewForm" :rules="renewRules">
        <a-form-item label="renewal duration" name="rentalDuration">
          <a-input
              v-model:value="renewForm.rentalDuration"
              type="text"
              autocomplete="off"
              placeholder="please enter the renewal period"
          />
        </a-form-item>
        <a-button class="ok-btn" block @click="showTransactionModal">
          renew
        </a-button>
      </a-form>
    </div>
  </renew-modal>
  <transaction-modal ref="transactionRef" @signedtransaction="renew" @close="cancelLoading"/>
  <transaction-modal ref="receiveRef" @signedtransaction="receivePenaltyAmount" @close="cancelLoading"/>
</template>

<script>
import {computed, onMounted, reactive, ref, toRefs, getCurrentInstance} from "vue";
import NoDataTable from "../../components/table/NoDataTable";
import PTable from "../../components/table/PTable";
import {useStore} from 'vuex'
import {timeToDay} from "@/utils/util";
import RenewModal from "../../components/model/index"
import TransactionModal from "../../components/model/transactionModal"

export default {
  name: "index",
  components: {
    NoDataTable,
    PTable,
    RenewModal,
    TransactionModal
  },
  setup() {
    const renewState = ref();
    const transactionRef = ref();
    const receiveRef = ref();
    const { proxy } = getCurrentInstance();
    const state = reactive({
      list: {
        total: 0,
        data: [],
        current: 1
      },
      loadLoading: false,
      renewForm: {
        rentalDuration: ''
      },
      loading: false,
      visible: false,
      renewIndex: ''
    });
    const renewRules = {
      rentalDuration: [
        { required: true, message: 'Please enter the renewal duration', trigger: 'blur' }
      ],
    }
    onMounted(() => {
      getResourceList();
    })
    const getStatus = computed(()=> (params) => {
      if (params.status === "Using") {
        return "renew"
      }else if (params.status === "Punished" && params.penalty_amount != "0") {
        return  "receive penalty amount"
      }else {
        return ""
      }
    })
    const getResourceList = async () => {
      let addressData =  await window.go.app.Wallet.GetWalletInfo();
      let data = await api.then(t => t.query.resourceOrder.userAgreements(addressData.address)).then(data => {
        return new Promise(function(resolve){
          resolve(data.toJSON())
        })
      }).then(data => api.then(t => t.query.resourceOrder.rentalAgreements.multi(data))).then(data => {
        return new Promise(function(resolve){
          resolve(data.map(t => t.toHuman()))
        })
      });
      let list = [];
      if (data.length > 0) {
        for (let i = 0; i < data.length; i++) {
          data[i]['createTime'] = await getDeadline(data[i].start);
          data[i]['endTime'] = await getDeadline(data[i].end);
          list.push(data[i])
        }
      }
      state.list.data = list;
    }
    const getDeadline = async (params) => {
      let apiPro = await api;
      let header = await apiPro.rpc.chain.getHeader();
      let space = parseInt(params.replace(",","")) - parseInt(header.number.toNumber());
      let now = new Date();
      return timeToDay(
          now.getTime() + space * 6 * 1000
      );
    }
    const ok = async (params) => {
      if (params.status === "Using") {
        showStakingModal(params.index);
      }else {
        showReceiveAmountModal(params.index)
      }
    }
    const showStakingModal = (index) => {
      state.visible = true;
      state.renewIndex = index;
    }
    const showReceiveAmountModal = async (index) => {
      state.loadLoading = true;
      state.renewIndex = index;
      let API = await api;
      let transaction = API.tx.resourceOrder.withdrawFaultExcution(index);
      receiveRef.value.openModal(transaction.method.toHex());
    }
    const receivePenaltyAmount = async (krp) => {
      let API = await api;
      API.tx.resourceOrder.withdrawFaultExcution(state.renewIndex).signAndSend(krp,{ nonce: -1},({ status, events, dispatchError }) => {
        if (status.isInBlock) {
          if (dispatchError) {
            if (dispatchError.isModule) {
              const decoded = API.registry.findMetaError(dispatchError.asModule);
              const { docs, name, section } = decoded;
              console.log(`${section}.${name}: ${docs.join(' ')}`);
              proxy.$message.error("Failed to claim:" + docs);
              state.loadLoading = false;
            } else {
              proxy.$message.error("Failed to claim:" + dispatchError.toString());
              state.loadLoading = false;
            }
          } else {
            state.renewIndex = '';
            getResourceList();
            state.loadLoading = false;
            proxy.$message.success("Received successfully");
          }
        }
      }).catch(() => {
        state.loadLoading = false;
        proxy.$message.error("Failed to claim");
      })
    }
    const showTransactionModal = () => {
      renewState.value
          .validate()
          .then(async () => {
            state.loading = true;
            let API = await api;
            let transaction = API.tx.resourceOrder.renewAgreement(
                state.renewIndex,
                state.renewForm.rentalDuration
            );
            transactionRef.value.openModal(transaction.method.toHex());
          })
    }
    const renew = async (krp) => {
      let API = await api;
      API.tx.resourceOrder.renewAgreement(state.renewIndex,state.renewForm.rentalDuration).signAndSend(krp,{ nonce: -1 },({ status, events, dispatchError }) => {
        if (status.isInBlock) {
          if (dispatchError) {
            if (dispatchError.isModule) {
              const decoded = API.registry.findMetaError(dispatchError.asModule);
              const { docs, name, section } = decoded;
              console.log(`${section}.${name}: ${docs.join(' ')}`);
              proxy.$message.error("Renewal failed:" + docs);
              state.loading = false;
            } else {
              proxy.$message.error("Renewal failed:" + dispatchError.toString());
              state.loading = false;
            }
          } else {
            state.loading = false;
            state.visible = false;
            state.renewIndex = ''
            renewClose();
            getResourceList();
            proxy.$message.success("Renewal succeeded");
          }
        }
      }).catch(() => {
        this.$message.error("Renewal failed");
        state.loading = false;
      })
    }
    const renewClose = () => {
      state.renewForm.rentalDuration = '';
      state.visible = false;
    }
    const cancelLoading = () => {
      state.loadLoading = false;
      state.loading = false;
    }
    const api = new useStore().state.api;
    return {
      ...toRefs(state),
      getResourceList,
      getStatus,
      ok,
      renew,
      receivePenaltyAmount,
      renewClose,
      renewRules,
      showTransactionModal,
      transactionRef,
      renewState,
      receiveRef,
      showReceiveAmountModal,
      cancelLoading,
      api
    }
  }
}
</script>

<style lang="scss" scoped>
.tableStyle {
  padding: 12px;
  background: white;
  border-radius: 8px;
  //height: 100%;
}
.staking-content {
  display: flex;
  align-items: center;
  .title {
    width: 100px;
    color: rgba(0, 0, 0, 0.85);
  }
}
.staking-footer {
  margin-top: 24px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
  .staking-btn-close {
    width: 100%;
  }
  .staking-btn-ok {
    background-color: rgb(24, 144, 255);
    color: white;
  }
}
.form-error-tip {
  color: #f5313d;;
  font-style: normal;
  font-weight: normal;
  font-size: 10px;
  line-height: 17px;
  margin-left: 80px;
}
</style>