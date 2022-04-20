<template>
  <div class="setting-container">
    <a-spin :tip="tip" :spinning="loading">
      <div class="wallet-connect" v-if="address != ''">
        <div class="wallet-connect-title">
          <div>
            <div class="wallet-title-text">
              <span>ADDRESS:</span>
            </div>
            <div class="wallet-connect-address" style="margin-top: 4px">
              <!-- copy password -->
              <span
              >{{ address }}
            <img
                style="cursor: pointer"
                v-clipboard="address"
                @click="message.success('Copy successfully')"
                src="../../assets/GrayCopy.png"
                alt=""
            />
          </span>
            </div>
          </div>
          <div style="margin-top: 8px">
            <span class="wallet-title-text">ACCOUNT BALANCE:</span>
            <div class="wallet-connect-address" style="margin-top: 4px">{{ accountAmount }}</div>
          </div>
        </div>
        <!-- UNBIND -->
        <a-button @click="forgotAddress" class="wallet-connect-btn">Unbind</a-button>
      </div>
      <a-form ref="settingState" :model="settingForm" :rules="settingRules">
        <a-form-item name="publicKey">
          <a-textarea v-model:value="settingForm.publicKey" placeholder="please enter public key" :rows="5"/>
        </a-form-item>
      </a-form>
      <div style="display: flex;align-items: center">
        <span style="width: 160px">Please input WsUrl:</span>
        <a-input type="text" v-model:value="msg" placeholder="edit me" @change="editApi"/>
      </div>
      <div class="node-address-style" v-if="settingData.peerId != ''">
        <span class="font-style">Node address：</span>
        <span>
          <a-tooltip placement="top">
              <template #title>
                <span>{{"/ip4/127.0.0.1/tcp/" + settingData.port + "/p2p/" + settingData.peerId}}</span>
              </template>
              <span>{{stringSplice("/ip4/127.0.0.1/tcp/" + settingData.port + "/p2p/" + settingData.peerId,65)}}</span>
          </a-tooltip>
        </span>
      </div>
      <div class="save-button-style">
        <a-button class="setting-button ok-btn" @click="setting">Save</a-button>
      </div>
    </a-spin>
  </div>
</template>

<script>
import {getCurrentInstance, onMounted, reactive, ref, toRefs} from "vue";
import { message } from "ant-design-vue";
import {useStore} from 'vuex'
// test
import { ApiPromise,WsProvider } from "@polkadot/api";
import types from "../../api/types";
export default {
  name: "index",
  setup(pro,context) {
    const { proxy } = getCurrentInstance();
    const settingState = ref();
    const store = new useStore();
    const state = reactive({
      settingForm: {
        publicKey: ""
      },
      loading: false,
      address: "",
      accountAmount: "0 Uint",
      settingData: {
        peerId: "",
        port: ""
      },
      msg: store.state.wsUrl
    })
    const settingRules = {
      publicKey: [
        {
          required: true,
          message: "Please enter public key",
          trigger: "blur",
        },
      ],
      privateKey: [
        {
          required: true,
          message: "Please enter private key",
          trigger: "blur",
        },
      ]
    }
    onMounted(() => {
      getSetting();
      getAddress();
    })
    const stringSplice = (str,length) => {
      if (str.length > length) {
        return str.substring(0,length) + "..."
      } else {
        return str
      }
    }
    const getSetting = () => {
      window.go.app.Setting.GetSetting().then(res => {
        state.settingForm.publicKey = res.PublicKey;
        state.settingData.peerId = res.PeerId;
        state.settingData.port = res.Port;
      })
    }
    const setting = () => {
      settingState.value.validate().then(() => {
        window.go.app.Setting.Setting(state.settingForm.publicKey).then(res => {
          if (res) {
            proxy.$message.success("Configured successfully")
          }
        }).catch((err) => {
          proxy.$message.error(err)
        })
      }).catch(() => {
        return false
      })
    }
    //UNBIND
    const forgotAddress = () => {
      state.loading = true
      window.go.app.Wallet.DeleteWallet();
      state.address = ""
      setTimeout(() => {
        context.emit("getAddress")
      },1000)
    }
    //get account address
    const getAddress = () => {
      window.go.app.Wallet.GetWalletInfo().then(res => {
        state.address = res.address
        api.then(t => {
          t.query.system.account(res.address).then(res =>{
            state.accountAmount = res.data.free.toHuman()
          })
        })
      }).catch(() => {
        state.address = ""
      })
    }
    const api = new useStore().state.api;
    const editApi = () => {
      if (state.msg === '') {
        proxy.$message.warning("WsUrl cannot be empty");
        return;
      }
      store.commit('setUrl',state.msg);
      const wsProvider = new WsProvider(state.msg);
      const newApi = ApiPromise.create({provider: wsProvider,types});
      store.commit('setApi',newApi);
    }
    return {
      ...toRefs(state),
      settingState,
      settingRules,
      setting,
      forgotAddress,
      getAddress,
      message,
      stringSplice,
      store,
      api,
      editApi
    }
  },
}


</script>

<style lang="scss" >
.setting-container {
  padding: 12px;
  background: white;
  height: 100%;
  border-radius: 8px;
  .form-error-tip {
    color: #f5313d;
    font-size: 10px;
    line-height: 17px;
    margin-left: 80px;
  }
  .wallet-connect {
    width: 100%;
    height: auto;
    padding: 12px;
    border: 1px solid #f6f8fc;
    box-sizing: border-box;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    .wallet-title-text {
      height: 22px;
      font-style: normal;
      font-weight: 600;
      font-size: 16px;
      line-height: 22px;
      color: #202839;
    }
    .wallet-connect-address {
      height: 20px;
      font-style: normal;
      font-weight: normal;
      font-size: 14px;
      line-height: 20px;
      color: #adb6ca;
    }
    .wallet-connect-btn {
      width: 92px;
      height: 36px;
      background: #4850ff;
      border-radius: 4px;
      font-style: normal;
      font-weight: normal;
      font-size: 14px;
      line-height: 20px;
      color: #ffffff;
      margin-left: auto;
    }
  }
  .setting-button {
    width: 500px;
    margin-bottom: 16px;
  }
  .save-button-style {
    display: flex;
    justify-content: center;
    margin-top: 20px;
  }
  .node-address-style {
    .font-style {
      font-size: 12px;
      font-weight: 600;
    }
  }
}
</style>
