<template>
  <modal-component ref="modalComponent" :title="'Import json'" :tip="'UPLOADING...'" :loading="loadFlag" :visible="visible" :close="importClose" >
    <a-upload :fileList="defaultFileList" :before-upload="handleImport">
      <div class="click-import-json">
          <span class="level-three-content-text">click to import json</span>
      </div>
    </a-upload>
    <div v-if="jsonValue.address" class="accountInfo">
      <div class="json-account-name">
        <span>NAME: {{ jsonValue.meta.name }}</span>
      </div>
      <div class="accountInfo-address">
          <span class="level-three-content-text" style="line-height: 17px"
          >{{ jsonValue.address }}
            <img
                @click="message.success('COPY SUCCESSFULLY')"
                style="cursor: pointer"
                v-clipboard="jsonValue.address"
                src="../../assets/GrayCopy.png"
            /></span>
      </div>
      <div class="level-three-content-text account-password">
        <span>PASSWORD：</span>
      </div>
      <a-input
          :type="changeHidden"
          v-model:value="password"
          placeholder="Please enter password"
      >
        <template #suffix>
          <img
              v-if="changeHidden === 'password'"
              style="cursor: pointer"
              @click="isHiddenPwd(false)"
              src="../../assets/HiddenPassword.png"
          />
          <img
              v-else
              style="cursor: pointer"
              @click="isHiddenPwd(true)"
              src="../../assets/ShowPassword.png"
          />
        </template>
      </a-input>
    </div>
    <div class="modal-footer">
      <a-button class="cancel-btn" @click="handleCancel">
        Cancel
      </a-button>
      <a-button
          class="ok-btn"
          style="margin-left: 12px"
          @click="getKeyringPair()"
      >
        OK
      </a-button>
    </div>
  </modal-component>
</template>
<script>
import {defineComponent, ref} from "vue";
import { Keyring } from "@polkadot/keyring";
import { message } from "ant-design-vue";
import ModalComponent from "../../components/model/index";
export default defineComponent({
  name: "ImportJsonModal",
  emits: ["updateAccountId:updateAccountId"],
  components: {
    ModalComponent
  },
  setup(proxy,context) {
    // pop up flag
    let visible = ref(false);
    let modalComponent = ref();
    // file list is empty
    const defaultFileList = ref([]);
    // store json data
    let jsonValue = ref({});
    // password
    let password = ref("");
    // loadFlag
    let loadFlag = ref(false);
    // input box type
    let changeHidden = ref("password");

    // show/hidden password
    function isHiddenPwd(hiddenPwdFlg) {
      if (hiddenPwdFlg) {
        changeHidden.value = "password";
      } else {
        changeHidden.value = "text";
      }
    }
    const importClose = () => {
      visible.value = false
    }
    // get the uploaded json file
    const handleImport = (file) => {
      const reader = new window.FileReader();
      reader.readAsText(file);
      reader.onloadend = () => {
        jsonValue.value = {};
        try {
          jsonValue.value = JSON.parse(reader.result);
        } catch (e) {
          message.error("the file format is incorrect please upload again");
          return false;
        }
        if (
          !jsonValue.value.address ||
          !jsonValue.value.encoded ||
          !jsonValue.value.encoding ||
          !jsonValue.value.meta
        ) {
          message.error("please import json in correct format");
          jsonValue.value = {};
        }
      };

      return false;
    };

    // verify json password
    function getKeyringPair() {
      // determine whether to import json
      if (!jsonValue.value.address) {
        message.error("please import json first");
        return;
      }
      // whether to enter a password
      if (password.value === "") {
        message.error("password can not be blank");
        return;
      }
      // becomes loading
      loadFlag.value = true;
      let json = jsonValue.value;
      let pwd = password.value;

      setTimeout(() => {
        try {
          const kr = new Keyring({
            type: "sr25519",
          });
          const krp = kr.addFromJson(json);
          // verify password
          krp.decodePkcs8(pwd);
          // backend
          window.backend.Wallet.SaveWallet(json.address, JSON.stringify(json)).then(() => {
            context.emit('getAddress');
            context.emit('setAddress');
          })
          handleCancel();
        } catch (error) {
          loadFlag.value = false;
          message.error("wrong password");
        }
      }, 100);

    }

    // close popup
    function handleCancel() {
      jsonValue.value = {};
      password.value = "";
      changeHidden.value = "password";
      loadFlag.value = false;
      visible.value = false;
    }

    // open popup
    function openModal() {
      jsonValue.value = {};
      password.value = "";
      changeHidden.value = "password";
      loadFlag.value = false;
      visible.value = true;
    }

    return {
      handleImport,
      fileList: ref([]),
      handleCancel,
      openModal,
      visible,
      defaultFileList,
      getKeyringPair,
      password,
      jsonValue,
      message,
      loadFlag,
      changeHidden,
      isHiddenPwd,
      modalComponent,
      importClose
    };
  },
});
</script>

<style lang="scss">
.accountInfo {
  width: 100%;
  margin-bottom: 20px;
}

.json-account-name {
  font-style: normal;
  font-weight: 600;
  font-size: 12px;
  line-height: 17px;
  color: #58637b;
}
.accountInfo-address {
  width: 100%;
  margin: 8px 0 20px;
  line-height: 17px;
}
.account-password {
  line-height: 17px;
  margin-bottom: 4px;
}

.import-json-title {
  width: 100%;
  height: 22px;
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  .import-json-close {
    margin-left: auto;
  }
}
.modal-footer {
  display: flex;
  align-items: center;
}
.click-import-json {
  margin-top: 16px;
  cursor: pointer;
  width: 100%;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
  border: 1px dashed #edeff1;
  border-radius: 4px;
}

.ant-upload.ant-upload-select {
  width: 100%;
}
</style>
