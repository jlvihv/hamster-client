<template>
<div class="login-container">
  <login-model :title="'Welcome to Hamster client'" :is-show-cancel="false" ref="loginRef" :visible="visible">
    <div class="model-content-login">
      <a-form ref="codeState" :model="codeForm" :rules="codeRules">
        <a-form-item name="mobile">
          <a-input
              v-model:value="codeForm.mobile"
              :disabled="disabled"
              type="text"
              autocomplete="off"
              placeholder="Please enter phone number"
              @change="changeMobile"
          />
        </a-form-item>
      </a-form>
      <a-form ref="loginState" :model="loginForm" :rules="loginRules">
        <a-form-item name="smsCode">
          <a-input v-model:value="loginForm.smsCode" type="text" autocomplete="off" placeholder="Please enter verification code">
            <template #suffix>
                <span v-show="getCodeButton" class="getCode" @click="getCode">
                  Get Verification Code
                </span>
              <span v-show="!getCodeButton" class="getCode">Retry Later: {{ time }}</span>
            </template>
          </a-input>
        </a-form-item>
        <a-button class="ok-btn" block @click="toLogin" :loading="loading">
          LogIn
        </a-button>
      </a-form>
    </div>
  </login-model>
</div>
</template>

<script>
import LoginModel from "../components/model/index"
import {getCurrentInstance, reactive, ref, toRefs} from "vue"
import {useRouter} from "vue-router";
export default {
  name: "login",
  components: {
    LoginModel
  },
  setup() {
    const router = useRouter();
    const { proxy } = getCurrentInstance();
    const codeState = ref();
    const loginState = ref();
    const state = reactive({
      visible: true,
      loginForm: {
        mobile: "",
        smsCode: ""
      },
      codeForm: {
        mobile:""
      },
      loading: false,
      getCodeButton: true,
      disabled: false,
      time: 60,
      timer: null,
    })
    const codeRules = {
      mobile: [
        {
          required: true,
          message: "Please enter phone number",
          trigger: "blur",
        },
        {
          pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
          message: "Please enter a valid phone number",
          trigger: "change",
        },
      ],
    }
    const loginRules = {
      smsCode: [
        {
          required: true,
          message: "Please enter verification code",
          validator: function (rule, value) {
            if (value === "") {
              return Promise.reject("Please enter verification code");
            } else if (value) {
              return Promise.resolve();
            }
          },
          trigger: "change",
        },
      ],
    }
    const getCode = () => {
      state.loginForm.smsCode = "";
      codeState.value
          .validate()
          .then(() => {
            state.getCodeButton = false;
            state.disabled = true;
            state.timer = setInterval(() => {
              state.time--;
              if (state.time <= 0) {
                state.time = 60;
                state.getCodeButton = true;
                state.disabled = false;
                clearInterval(state.timer);
              }
            }, 1000);
            window.backend.Login.GetCode(state.codeForm.mobile).then(() => {
              proxy.$message.success("SMS sent successfully");
            }).catch((err) => {
              state.disabled = false;
              proxy.$message.error(err);
            })
          }).catch(() => {
            return false;
      })
    }
    const toLogin = () => {
      // router.push("/resource")
      const result = codeState.value
          .validate()
          .then(() => {
            return true;
          })
          .catch(() => {
            return false;
          });

      if (!result) return;
      state.loginForm.mobile = state.codeForm.mobile;
      loginState.value
          .validate()
          .then(() => {
            state.loading = true;
            window.backend.Login.Login(state.loginForm.mobile,state.loginForm.smsCode).then(() => {
              proxy.$message.success("Login successful");
              state.loading = false;
              state.visible = false;
              router.push("/resource")
            }).catch((err) => {
              proxy.$message.error(err);
              state.loading = false;
            })
          }).catch(() => {
              return false;
      })
    }
    const changeMobile = () => {
      state.loginForm.smsCode = "";
      state.loginForm.mobile = "";
    }
    return {
      ...toRefs(state),
      getCode,
      codeState,
      loginRules,
      codeRules,
      toLogin,
      loginState,
      changeMobile
    }
  }
}
</script>

<style lang="scss" scoped>
.login-container {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: url("../assets/loginBack.png");
}
.model-content-login {
  margin-top: 64px;
}
.getCode {
  font-weight: 400;
  color: #4850ff;
  &:first-child {
    cursor: pointer;
  }
}
</style>
