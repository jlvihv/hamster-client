<template>
  <a-spin :tip="'loading'" :spinning="loadLoading" class="spin-style">
    <div class="tableStyle">
      <no-data-table v-if="list.data.length === 0">
        <a-table-column
            title="Resource ID"
            data-index="resource_index"
        />
        <a-table-column title="System"/>
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
        <a-table-column title="Operate" :width="90"></a-table-column>
      </no-data-table>
      <p-table
          v-show="list.data.length != 0"
          :row-key="t => t.index"
          :dataList="list"
          @get-data="getList"
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
        <a-table-column title="Creation" :width="120">
          <template #default="{ record }">
            <span>{{record.createTime}}</span>
          </template>
        </a-table-column>
        <a-table-column
            :width="120"
            title="Expire"
            data-index="expireTime"
        >
          <template #default="{ record }">
            <span>{{record.endTime}}</span>
          </template>
        </a-table-column>
        <a-table-column title="Operate" :width="80">
          <template #default="{ record }">
            <a style="color: #4850FF" @click="link(record)">Link</a>
          </template>
        </a-table-column>
      </p-table>
    </div>
  </a-spin>
  <tip-modal :title="'message notification'" :is-show-cancel="true" ref="loginRef" :visible="visible" :title-big-style="false" :close="close">
    <div class="tip-content">
      <span class="font-style">You have not configured the public key, please go to "Setting" to configure the public key</span>
    </div>
    <div class="tip-button">
      <a-button class="ok-btn" @click="goSetting">go to configuration</a-button>
    </div>
  </tip-modal>
  <link-modal :title="'link'" :is-show-cancel="true" ref="linkRef" :visible="linkVisible" :title-big-style="false" :close="linkClose" :loading="loading">
    <div style="margin-top: 32px">
      <a-form ref="linkState" :model="linkForm" :rules="linkRules">
        <a-form-item name="peerId">
          <a-input
              v-model:value="linkForm.peerId"
              type="text"
              autocomplete="off"
          />
        </a-form-item>
        <a-form-item name="port">
          <a-input
              v-model:value="linkForm.port"
              type="text"
              autocomplete="off"
              placeholder="please enter the port number"
          />
        </a-form-item>
        <a-button class="ok-btn" block @click="toLink">
          link
        </a-button>
      </a-form>
    </div>
  </link-modal>
  <link-tip :title="'initial configuration'" :is-show-cancel="true" :visible="linkTipVisible" :title-big-style="false" :close="linkTipClose" :loading="tipLoading">
    <div style="margin-top: 32px;margin-bottom: 20px">
      <span>You have not initialized the configuration, please initialize the configuration</span>
    </div>
    <a-button class="ok-btn" block @click="setting">
      initial configuration
    </a-button>
  </link-tip>
</template>

<script>
import {getCurrentInstance, onMounted, reactive, ref, toRefs} from "vue";
import NoDataTable from "../../components/table/NoDataTable";
import PTable from "../../components/table/PTable";
import {timeToDay} from "../../utils/util";
import TipModal from "../../components/model/index"
import {useRouter} from "vue-router";
import LinkModal from "../../components/model/index"
import LinkTip from "../../components/model/index"
import {useStore} from "vuex";

export default {
  name: "index",
  components: {
    NoDataTable,
    PTable,
    TipModal,
    LinkModal,
    LinkTip
  },
  setup() {
    const router = useRouter();
    const linkState = ref();
    const { proxy } = getCurrentInstance();
    const state = reactive({
      loadLoading:false,
      tipLoading: false,
      loading: false,
      visible: false,
      linkVisible: false,
      linkTipVisible: false,
      linkForm: {
        port:"",
        peerId: ""
      },
      list: {
        total: 0,
        data: [],
        current: 1
      },
    })
    const linkRules = {

    }
    //p2p connection
    const toLink = () => {
      linkState.value.validate().then(() => {
        state.loading = true
        window.go.app.P2p.Link(parseInt(state.linkForm.port),state.linkForm.peerId).then(() => {
          state.loading = false;
          proxy.$message.success("connection succeeded")
          state.linkVisible = false
        }).catch((err) => {
          proxy.$message.error(err)
          state.loading = false
        })
      }).catch(() => {
        return false
      })
    }
    //get my resources
    const getList = async () => {
      let addressData =  await window.go.app.Wallet.GetWalletInfo();
      let API = await api;
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
          let t = await API.query.provider.resources(data[i].resource_index);
          if (t.toJSON() != null) {
            data[i]['status']= t.toJSON().status
            list.push(data[i])
          }
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
    const link = (record) => {
      //query if p2p is configured
      window.go.app.P2p.IsP2PSetting().then(res => {
        if (res) {
          state.linkForm.peerId = record.peer_id
          state.linkVisible = true
        }else {
          //initial configuration
          state.linkTipVisible = true
        }
      })
    }
    onMounted(() => {
      isSettingPublicKey();
      getList()
    })
    const getResourceList = () => {
      state.loadLoading = true
      window.go.app.Resource.GetResources().then(res => {
        state.list.data = res
        state.loadLoading = false
      }).catch(() => {
        state.loadLoading = false
      })
    }
    const close = () => {
      state.visible = false
    }
    const linkClose = () => {
      state.linkVisible = false
    }
    const linkTipClose = () => {
      state.linkTipVisible = false;
    }
    const getAddress = () => {

    }
    const isSettingPublicKey = () => {
      window.go.app.Account.IsAccountSetting().then(res => {
        if (res) {
          getList()
          // getResourceList()
        }else {
          state.visible = true
        }
      })
    }
    //p2p Initialization Configuration
    const setting = () => {
      state.tipLoading = true
      window.go.app.Setting.InitP2pSetting().then(() => {
        proxy.$message.success("configured successfully")
        state.tipLoading = false
        state.linkTipVisible = false
      }).catch(() => {
        proxy.$message.error("configuration failed")
        state.tipLoading = false
      })
    }
    const goSetting = () => {
      router.push("/setting")
    }
    const api = new useStore().state.api;
    return {
      ...toRefs(state),
      getList,
      link,
      getResourceList,
      timeToDay,
      close,
      goSetting,
      isSettingPublicKey,
      linkClose,
      linkState,
      linkRules,
      toLink,
      linkTipClose,
      setting,
      getAddress,
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
.tip-content {
  margin-top: 20px;
  .font-style {
    color: #58637B;
    font-size: 14px;
    line-height: 17px;
  }
}
.tip-button {
  margin-top: 20px;
}
</style>
