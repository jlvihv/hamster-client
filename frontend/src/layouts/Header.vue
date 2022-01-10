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
      <a-button class="logout" @click="logout">Logout</a-button>
    </div>
  </div>
  <!-- import json popup -->
  <import-json-modal ref="refImportJson"  @updateAccountId:updateAccountId="updateAccountId" @getAddress="getAddress" @setAddress="settingAddress" ></import-json-modal>
</template>

<script>
import {getCurrentInstance, reactive, toRefs, ref, onMounted} from "vue";
import {useRouter} from "vue-router";
import ImportJsonModal from "../components/importJson/index";
import { message } from "ant-design-vue";
export default {
  name: "MobileHeader",
  components: {
    ImportJsonModal
  },
  setup(pro,context) {
    const router = useRouter();
    const { proxy } = getCurrentInstance();
    const state = reactive({
      address: "",
      headerInfo: {},
    });
    let accountId = ref("");
    // import json popup ref
    const refImportJson = ref();
    const logout = () => {
      window.backend.Login.Logout().then(() => {
        router.push("/login")
      }).catch((err) => {
        proxy.$message.error(err)
      })
    };
    //import account json
    const showImportJsonModal = () => {
      refImportJson.value.openModal();
    }
    onMounted(() => {
      getAddress()
    })
    //get account address
    const getAddress = () => {
      window.backend.Wallet.GetWalletInfo().then(res => {
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
      logout,
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
  .logout {
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
