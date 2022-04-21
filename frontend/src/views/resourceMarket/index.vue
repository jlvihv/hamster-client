<template>
  <a-spin :tip="'loading'" :spinning="loadLoading" class="spin-style">
    <div class="resource-market">
      <no-data-table v-if="list.data.length === 0">
        <a-table-column
            title="Resource ID"
            data-index="index"
        />
        <a-table-column
            title="Unit Price"
        />
        <a-table-column title="System"/>
        <a-table-column title="CPU Model"/>
        <a-table-column
            title="CPU"
            data-index="cpu"
            :width="70"
        />
        <a-table-column
            title="RAM"
            data-index="memory"
            :width="70"
        />
        <a-table-column title="Creation" :width="120"/>
        <a-table-column
            :width="120"
            title="Expire"
            data-index="expireTime"
        />
        <a-table-column title="Operate" :width="80"></a-table-column>
      </no-data-table>
      <p-table
          v-show="list.data.length != 0"
          :row-key="t => t.index"
          :dataList="list"
          @get-data="getResourceList"
          ref="table-page"
      >
        <a-table-column
            title="Resource ID"
            data-index="index"
        />
        <a-table-column title="Unit Price">
          <template #default="{ record }">
            <span>{{ record.rental_info.rent_unit_price }}</span>
          </template>
        </a-table-column>
        <a-table-column title="System">
          <template #default="{ record }">
            <span>{{ record.config.system }}</span>
          </template>
        </a-table-column>
        <a-table-column title="CPU Model">
          <template #default="{ record }">
            <span>{{ record.config.cpu_model }}</span>
          </template>
        </a-table-column>
        <a-table-column title="CPU">
          <template #default="{ record }">
            <span>{{ record.config.cpu }}</span>
          </template>
        </a-table-column>
        <a-table-column title="Memory">
          <template #default="{ record }">
            <span>{{ record.config.memory + ' G'}}</span>
          </template>
        </a-table-column>
        <a-table-column title="Rental duration" :sorter="(a, b) => a.index - b.index" :width="125">
          <template #default="{ record }">
            <span>{{ record.rental_info.rent_duration}}</span>
          </template>
        </a-table-column>
        <a-table-column title="Operate" :width="95">
          <template #default="{ record }">
            <a style="color: #4850FF" @click="showBuyResourceModal(record)">Purchase</a>
          </template>
        </a-table-column>
      </p-table>
    </div>
  </a-spin>
  <buy-modal :title="'Purchase resources'" :is-show-cancel="true" ref="renewRef" :visible="visible" :title-big-style="false" :close="buyClose" :loading="loading">
    <div style="margin-top: 32px">
      <a-form ref="buyFormRef" :model="buyResourceForm" :rules="buyResourceRules">
        <a-form-item name="rentalDuration">
          <h6 style="margin-bottom: 8px">Rental duration:</h6>
          <a-input
              oninput="value=value.replace(/[^\d]/g,'');if(value>=4294967295)value=4294967295"
              v-model:value="buyResourceForm.rentalDuration"
              type="text"
              autocomplete="off"
              placeholder="please enter the renewal period"
              @change="getFee"
          />
        </a-form-item>
        <a-form-item name="publicKey">
          <h6 style="margin-bottom: 8px">User public key:</h6>
          <a-textarea v-model:value="buyResourceForm.publicKey" placeholder="Please enter user public key" :rows="3" />
        </a-form-item>
        <h6>Rental price: {{showPrice}}</h6>
        <h6>Service Charge: {{showFee}}</h6>
        <a-button class="ok-btn" block @click="showTransactionModal">
          purchase
        </a-button>
      </a-form>
    </div>
  </buy-modal>
  <transaction-modal ref="buyResourceRef" @signedtransaction="buyResource" @close="cancelLoading"/>
</template>

<script>
import {getCurrentInstance, onMounted, reactive, ref, toRefs} from "vue";
import NoDataTable from "../../components/table/NoDataTable";
import PTable from "../../components/table/PTable";
import TransactionModal from "../../components/model/transactionModal";
import BuyModal from "../../components/model/index";
import {useStore} from "vuex";
import {ApiPromise, WsProvider} from "@polkadot/api";
import types from "../../api/types";

export default {
  name: "index",
  components: {
    NoDataTable,
    PTable,
    TransactionModal,
    BuyModal,
  },
  setup: function () {
    const buyResourceRef = ref();
    const buyFormRef = ref();
    const { proxy } = getCurrentInstance();
    const state = reactive({
      showPrice: "0.0000 Uint",
      uintPrice: 0,
      showFee: '0.0000 Uint',
      loading: false,
      visible: false,
      loadLoading: false,
      resourceIndex: '',
      buyResourceForm: {
        rentalDuration: '',
        publicKey: ''
      },
      list: {
        total: 0,
        data: [],
        current: 1
      },
    });
    const store = new useStore();
    const initApi = async () => {
      let res = await window.go.app.Setting.GetSetting()
      store.commit('setUrl',res.WsUrl);
      const wsProvider = new WsProvider(res.WsUrl);
      const newApi = ApiPromise.create({provider: wsProvider,types});
      store.commit('setApi',newApi);
    }
    initApi()
    const buyResourceRules = {
      rentalDuration: [
        { required: true, message: 'Please enter the renewal duration', trigger: 'blur' }
      ],
      publicKey: [
        { required: true, message: 'Please enter the user public key', trigger: 'blur' }
      ],
    };
    onMounted(() => {
      getResourceList();
    })
    const getFee = async () => {
      let API = await api;
      let addressData = await window.go.app.Wallet.GetWalletInfo();
      if (addressData) {
        let n = Number(state.buyResourceForm.rentalDuration);
        if (!isNaN(n)) {
          // eslint-disable-next-line no-undef
          let price = state.uintPrice * BigInt(state.buyResourceForm.rentalDuration);
          let API = await api;
          let data = await API.registry.createType('Balance', price);
          state.showPrice = data.toHuman();
        }
        if (state.buyResourceForm.rentalDuration <= 4294967295) {
          let res = await API.tx.resourceOrder.createOrderInfo(
              state.resourceIndex,
              state.buyResourceForm.rentalDuration,
              state.buyResourceForm.publicKey).paymentInfo(addressData.address);
          state.showFee = res.partialFee.toHuman();
        } else {
          state.showFee = "0.0000 Unit";
        }
      } else {
        state.showFee = "0.0000 Unit";
      }
    }
    const getResourceList = async () => {
      let API = await api;
      let data = await API.query.provider.resources.entries();
      let list = [];
      data.forEach(([k, v]) => {
        let resource = v.toHuman();
        if (resource.status === "Unused") {
          let unit_price = v.value.rental_info.rent_unit_price.toBigInt();
          resource['unitPrice'] = unit_price;
          list.push(resource)
        }
      });
      state.list.data = list;
    }
    const showTransactionModal =  () => {
      buyFormRef.value.validate().then(async () => {
        state.loading = true;
        let API = await api;
        let transaction = API.tx.resourceOrder.createOrderInfo(state.resourceIndex,state.buyResourceForm.rentalDuration,state.buyResourceForm.publicKey);
        buyResourceRef.value.openModal(transaction.method.toHex());
      })
    }
    const showBuyResourceModal = async (params) => {
      let addressData =  await window.go.app.Wallet.GetWalletInfo();
      if (addressData) {
        state.resourceIndex = params.index;
        state.uintPrice = params.unitPrice
        state.visible = true;
      } else {
        proxy.$message.warning('Please import user JSON');
      }
    }
    const buyResource = async (krp) => {
      let API = await api;
      API.tx.resourceOrder.createOrderInfo(state.resourceIndex,state.buyResourceForm.rentalDuration,state.buyResourceForm.publicKey).signAndSend(krp,{nonce: -1},({ status, events, dispatchError }) => {
        if (status.isInBlock) {
          if (dispatchError) {
            if (dispatchError.isModule) {
              const decoded = API.registry.findMetaError(dispatchError.asModule);
              const { docs, name, section } = decoded;
              console.log(`${section}.${name}: ${docs.join(' ')}`);
              proxy.$message.error("Failed to purchase resources:" + docs);
              state.loading = false;
            } else {
              proxy.$message.error("Failed to purchase resources:" + dispatchError.toString());
              state.loading = false;
            }
          } else {
            state.loading = false;
            state.resourceIndex = '';
            state.uintPrice = 0;
            proxy.$message.success('Successfully purchased resources');
            getResourceList();
            state.visible = false;
          }
        }
      }).catch((error) => {
        console.log(error);
        state.loading = false;
        proxy.$message.error('Failed to purchase resources');
      })
    }
    const buyClose = () => {
      state.buyResourceForm.publicKey = "";
      state.buyResourceForm.rentalDuration = "";
      state.uintPrice = 0;
      state.resourceIndex = "";
      state.visible = false;
    }
    const cancelLoading = () =>  {
      state.loading = false;
    }
    const api = new useStore().state.api;
    return {
      ...toRefs(state),
      getResourceList,
      buyResourceRef,
      showBuyResourceModal,
      buyResource,
      buyClose,
      buyFormRef,
      buyResourceRules,
      showTransactionModal,
      getFee,
      cancelLoading,
      api
    }
  }
}
</script>

<style scoped>

</style>