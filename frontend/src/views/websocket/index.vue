<template>
  <a-spin :tip="'loading'" :spinning="loadLoading" class="spin-style">
    <div>
      <a-button class="ok-btn" block @click="showApplyResourceModal">
        apply resource
      </a-button>
      <a-button class="ok-btn" block @click="showDeployModal">
        deploy
      </a-button>
    </div>
  </a-spin>
  <apply-modal :title="'apply'" :is-show-cancel="true" ref="applyRef" :visible="applyVisible" :title-big-style="false" :close="applyClose" :loading="applyLoading">
    <div style="margin-top: 32px">
      <a-form ref="applyState" :model="applyForm" >
        <a-form-item name="cpu">
          <a-input
              v-model:value="applyForm.cpu"
              type="text"
              autocomplete="off"
              placeholder="please enter the cpu"
          />
        </a-form-item>
        <a-form-item name="memory">
          <a-input
              v-model:value="applyForm.memory"
              type="text"
              autocomplete="off"
              placeholder="please enter the memory"
          />
        </a-form-item>
        <a-form-item name="duration">
          <a-input
              v-model:value="applyForm.duration"
              type="text"
              autocomplete="off"
              placeholder="please enter the duration"
          />
        </a-form-item>
        <a-form-item name="publicKey">
          <a-input
              v-model:value="applyForm.publicKey"
              type="text"
              autocomplete="off"
              placeholder="please enter the publicKey"
          />
        </a-form-item>
        <a-button class="ok-btn" block @click="showApplyTransModel">
          apply
        </a-button>
      </a-form>
    </div>
  </apply-modal>
  <deploy-modal :title="'deploy'" :is-show-cancel="true" ref="deployRef" :visible="deployVisible" :title-big-style="false" :close="deployClose" :loading="deployLoading">
    <div style="margin-top: 32px">
      <a-form ref="applyState" :model="deployForm" >
        <a-form-item name="nodeEthereumUrl">
          <a-input
              v-model:value="deployForm.nodeEthereumUrl"
              type="text"
              autocomplete="off"
              placeholder="please enter the nodeEthereumUrl"
          />
        </a-form-item>
        <a-form-item name="ethereumUrl">
          <a-input
              v-model:value="deployForm.ethereumUrl"
              type="text"
              autocomplete="off"
              placeholder="please enter the ethereumUrl"
          />
        </a-form-item>
        <a-form-item name="ethereumNetwork">
          <a-input
              v-model:value="deployForm.ethereumNetwork"
              type="text"
              autocomplete="off"
              placeholder="please enter the ethereumNetwork"
          />
        </a-form-item>
        <a-form-item name="indexerAddress">
          <a-input
              v-model:value="deployForm.indexerAddress"
              type="text"
              autocomplete="off"
              placeholder="please enter the indexerAddress"
          />
        </a-form-item>
        <a-form-item name="mnemonic">
          <a-input
              v-model:value="deployForm.mnemonic"
              type="text"
              autocomplete="off"
              placeholder="please enter the mnemonic"
          />
        </a-form-item>
        <a-button class="ok-btn" block @click="deploy">
          deploy
        </a-button>
      </a-form>
    </div>
  </deploy-modal>
  <transaction-modal ref="applyTransRef" @signedtransaction="apply" @close="cancelLoading"/>
  <div id="terminal"/>
</template>

<script>
import 'xterm/css/xterm.css'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { AttachAddon } from 'xterm-addon-attach'
import {getCurrentInstance, onBeforeUnmount, onMounted, reactive, ref, toRefs} from "vue";
import ApplyModal from "../../components/model/index"
import DeployModal from "../../components/model/index"
import TransactionModal from "../../components/model/transactionModal"
import {useStore} from "vuex";
export default {
  name: "index",
  components: {
    ApplyModal,
    DeployModal,
    TransactionModal
  },
  setup() {
    const applyTransRef = ref();
    const api = new useStore().state.api;
    const { proxy } = getCurrentInstance();
    const state = reactive({
      loadLoading:false,
      applyLoading: false,
      deployLoading: false,
      applyVisible: false,
      applyTipVisible: false,
      deployVisible: false,
      deployTipVisible: false,
      applyForm: {
        cpu:"",
        memory: "",
        duration: "",
        publicKey: ""
      },
      deployForm: {
        nodeEthereumUrl: "mainnet:https://eth-mainnet.alchemyapi.io/v2/wHl5-FZKD68DlhCuUiOXgYh4Z0q0L5fh",   //graph_node 配置的以太坊节点
        ethereumUrl: "https://rinkeby.infura.io/v3/af7a79eb36f64e609b5dda130cd62946",     // index-service,index-agent配置的以太坊节点
        ethereumNetwork: "rinkeby",  //以太坊网络
        indexerAddress: "0x9438BbE4E7AF1ec6b13f75ECd1f53391506A12DF",   //索引人地址
        mnemonic:"please output text solve glare exit divert boil nerve eagle attack turkey"  //索引人助记词
      }
    })
    onMounted(() => {
      var term = new Terminal();
      var fitAddon = new FitAddon()
      // var socket = new WebSocket(`ws://localhost:2375/containers/1eae0f865dc9/attach/ws?logs=0&stream=1&stdin=1&stdout=1&stderr=1`)
      var socket = new WebSocket(`ws://localhost:10771/api/v1/thegraph/ws?serviceName=index-cli`)
      var attachAddon = new AttachAddon(socket)
      term.loadAddon(attachAddon)
      term.loadAddon(fitAddon)
      term.open(document.getElementById('terminal'))
      fitAddon.fit()
      term.focus()
      socket.onopen = () => { socket.send('\n') }
      window.onresize = function() { // 窗口尺寸变化时，终端尺寸自适应
        fitAddon.fit()
      }
      proxy.term = term
      proxy.socket = socket
    })
    const showApplyTransModel = async () => {
      state.loadLoading = true;
      state.applyLoading = true
      let API = await api;
      let transaction = API.tx.resourceOrder.applyFreeResource(state.applyForm.cpu,state.applyForm.memory,state.applyForm.duration,state.applyForm.publicKey,1);
      applyTransRef.value.openModal(transaction.method.toHex());
    }
    const showApplyResourceModal = () => {
      state.applyVisible = true
    }
    const showDeployModal = () => {
      // window.go.app.Application.DeleteGraphAndParams(1).then((res) => {
      //   console.log(res);
      // }).catch(() => {
      //
      // })
      state.deployVisible = true
    }
    const apply = async (krp) => {
      let API = await api;
      API.tx.resourceOrder.applyFreeResource(state.applyForm.cpu,state.applyForm.memory,state.applyForm.duration,state.applyForm.publicKey,1).signAndSend(krp,{nonce: -1},({ status, events, dispatchError }) => {
        if (status.isInBlock) {
          if (dispatchError) {
            if (dispatchError.isModule) {
              const decoded = API.registry.findMetaError(dispatchError.asModule);
              const { docs, name, section } = decoded;
              console.log(`${section}.${name}: ${docs.join(' ')}`);
              proxy.$message.error("Apply resource failed:" + docs);
              state.loadLoading = false;
              state.applyLoading = false;
            } else {
              proxy.$message.error("Apply resource:" + dispatchError.toString());
              state.loadLoading = false;
              state.applyLoading = false;
            }
          } else {
            events.forEach(({ phase, event: { data, method, section } }) => {
              if (method == 'FreeResourceApplied') {
                window.go.app.Account.SaveOrderIndex(data.toJSON()[1]).then(() => {
                  state.applyVisible = false
                  state.loadLoading =false
                  state.applyLoading = false;
                  Object.keys(state.applyForm).forEach(key => state.applyForm[key] = "")
                  proxy.$message.success("Apply resource successfully");
                })
              }
            });
          }
        }
      }).catch(() => {
        state.loadLoading = false;
        state.applyLoading = false;
        proxy.$message.error("Apply resource failed");
      })
    }
    const deploy = () => {
      // window.go.app.Setting.InitP2pSetting().then(() => {

      // })
      state.deployLoading = true
      window.go.app.Deploy.DeployTheGraph(state.deployForm.nodeEthereumUrl,state.deployForm.ethereumUrl,state.deployForm.ethereumNetwork,state.deployForm.indexerAddress,state.deployForm.mnemonic).then(() => {
        state.deployLoading = false
        state.deployVisible =false
        // Object.keys(state.deployForm).forEach(key => state.deployForm[key] = "")
        proxy.$message.success("Deploy the graph successfully");
      }).catch(() => {
        state.deployLoading = false
        proxy.$message.error("Deploy the graph failed");
      })
    }
    const applyClose = () => {
      Object.keys(state.applyForm).forEach(key => state.applyForm[key] = "")
      state.applyVisible = false
    }
    const deployClose = () => {
      // Object.keys(state.deployForm).forEach(key => state.deployForm[key] = "")
      state.deployVisible = false

    }
    const cancelLoading = () => {
      state.loadLoading = false;
    }
    onBeforeUnmount(() => {
      proxy.term.close()
      proxy.socket.dispose()
    })
    return {
      ...toRefs(state),
      showApplyResourceModal,
      showDeployModal,
      apply,
      deploy,
      applyClose,
      deployClose,
      cancelLoading,
      api,
      showApplyTransModel,
      applyTransRef
    }
  }
}
</script>

<style scoped>
#terminal{
  height: 100%;
}
</style>