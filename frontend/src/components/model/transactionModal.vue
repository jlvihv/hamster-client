<template>
  <basic-modal :title="'transaction'" :is-show-cancel="true" ref="renewRef" :visible="visible" :title-big-style="false" :close="transactionClose" :loading="loading">
    <div style="margin-top: 32px">
      <div style="margin-bottom: 8px">
        <span style="margin-bottom: 4px">transaction hash:</span>
        <div>{{ transactionHash }}</div>
      </div>
      <a-form ref="transactionState" :model="transactionForm" :rules="transactionRules">
        <a-form-item label="password" name="password">
          <a-input
              v-model:value="transactionForm.password"
              type="password"
              autocomplete="off"
              placeholder="please enter the account password"
          />
        </a-form-item>
        <a-button class="ok-btn" block @click="sign">
          sign the transaction
        </a-button>
      </a-form>
    </div>
  </basic-modal>
</template>

<script>
import BasicModal from './index'
import {reactive, ref, toRefs,getCurrentInstance} from "vue";
import { Keyring } from "@polkadot/keyring";
export default {
  name: "transactionModal",
  components: {
    BasicModal
  },
  setup(props, context) {
    const state = reactive({
      loading: false,
      visible: false,
      transactionHash: '',
      transactionForm: {
        password: ''
      }
    });
    const { proxy } = getCurrentInstance();
    const transactionState = ref();
    const transactionRules = {
      password: [
        { required: true, message: 'Please enter the account password', trigger: 'blur' }
      ]
    };
    const sign = () => {
      transactionState.value
      .validate()
      .then(async () => {
        state.loading = true;
        let addressData = await window.go.app.Wallet.GetWalletInfo();
        setTimeout(() => {
          try {
            let accountJson = JSON.parse(addressData.address_json);
            const kr = new Keyring({
              type: "sr25519",
            });
            const krp = kr.addFromJson(accountJson);
            krp.decodePkcs8(state.transactionForm.password);
            state.loading = false;
            state.visible = false;
            state.transactionForm.password = '';
            state.transactionHash = '';
            context.emit("signedtransaction", krp);
          } catch (err) {
            state.loading = false;
            proxy.$message.error('Password error')
          }
        }, 100);
      }).catch((error) => {
        console.log(error);
        proxy.$message.error('Parameter verification error')
      })
    }
    const transactionClose = () => {
      state.visible = false;
      state.transactionForm.password = '';
      state.transactionHash = ''
      context.emit("close");
    }
    const openModal = (hash) => {
      state.transactionHash = hash;
      state.visible = true;
    }
    return {
      ...toRefs(state),
      transactionRules,
      sign,
      openModal,
      transactionClose,
      transactionState
    }
  }
}
</script>

<style scoped>

</style>