<template>
  <div class="header">
    <div class="header-right">
      <div  v-if="address !== ''" class="address-style">
        {{address}}
        <img
            style="cursor: pointer"
            v-clipboard="address"
            @click="message.success('Copy successfully')"
            src="../assets/GrayCopy.png"
            alt=""
        />
      </div>
      <a-button v-else class="import-btn" @click="showImportJsonModal">Import Account</a-button>
    </div>
  </div>
  <!-- import json popup -->
  <import-json-modal ref="refImportJson"  @updateAccountId:updateAccountId="updateAccountId" @getAddress="getAddress" @setAddress="settingAddress" ></import-json-modal>
</template>

<script>
import {reactive, toRefs, ref, onMounted} from "vue";
import ImportJsonModal from "../components/importJson/index";
import { message } from "ant-design-vue";
import {ApiPromise, WsProvider} from "@polkadot/api";
import types from "../api/types";
import {useStore} from "vuex";
export default {
  name: "MobileHeader",
  components: {
    ImportJsonModal
  },
  setup(pro,context) {
    const state = reactive({
      address: "",
      headerInfo: {},
    });
    let accountId = ref("");
    // import json popup ref
    const refImportJson = ref();
    //import account json
    const showImportJsonModal = () => {
      refImportJson.value.openModal();
    }
    const store = new useStore();
    const initApi = () => {
      window.go.app.Setting.GetSetting().then(res => {
        store.commit('setUrl',res.WsUrl);
        const wsProvider = new WsProvider(res.WsUrl);
        const newApi = ApiPromise.create({provider: wsProvider,types});
        store.commit('setApi',newApi);
      })
    }
    onMounted(() => {
      getAddress()
      initApi()
    })
    //get account address
    const getAddress = () => {
      window.go.app.Wallet.GetWalletInfo().then(res => {
        state.address = res.address
        context.emit("changeLoading")
      }).catch(() => {
        state.address = ""
        context.emit("changeLoading")
      })
    }
    // update address after importing json
    function updateAccountId() {
      accountId.value = JSON.parse(
          window.localStorage.getItem("accountJson")
      ).address;
    }
    //unbind
    const forgotAddress = () => {

    }
    //settingAddress
    const settingAddress = () => {
      context.emit("settingAddress")
    }
    return {
      ...toRefs(state),
      accountId,
      refImportJson,
      showImportJsonModal,
      updateAccountId,
      forgotAddress,
      getAddress,
      settingAddress,
      message
    };
  },
};
</script>

<style lang="scss">
.header {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  .import-btn {
    font-weight: 600;
    font-size: 12px;
    line-height: 17px;
    color: #4850FF;
    padding: 0px;
    margin-right: 23px;
  }
  .header-right {
    display: flex;
    align-items: center;
    .address-style {
      padding: 8px 12px;
      margin-right: 8px;
      background: #F0F3FF;
      border-radius: 31px;
      line-height: 17px;
    }
  }
}
</style>
