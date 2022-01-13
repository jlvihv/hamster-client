<template>
  <div class="home">
    <div class="home-container">
      <a-layout class="layout-container">
        <a-layout-sider width="142px" class="menu-left-style">
          <div class="logo">
            <img src="@/assets/logo.png" width="100" height="15"/>
          </div>
          <div class="custom-menu">
            <Menu />
          </div>
        </a-layout-sider>
        <a-layout>
          <a-layout-header class="header">
            <Header ref="pcHeader" @settingAddress="getSettingAddress" @changeLoading="changeLoading"> </Header>
          </a-layout-header>
          <a-layout-content class="body-content">
            <router-view v-slot="{Component}" @getAddress="getAddress">
              <keep-alive :include="['setting','resource']">
                <component :is="Component" ref="settingRef"/>
              </keep-alive>
            </router-view>
<!--            <router-view @getAddress="getAddress" ></router-view>-->
          </a-layout-content>
        </a-layout>
      </a-layout>
    </div>
  </div>
</template>

<script>
import Header from "./Header";
import Menu from "./Menu";
import {onMounted,ref} from "vue";
export default {
  name: "MobileLayout",
  components: {
    Header,
    Menu,
  },
  setup() {
    const pcHeader = ref()
    const settingRef = ref()
    onMounted(() => {
    })
    //get the header address
    const getAddress = () => {
      pcHeader.value.getAddress()
    }
    //after importing set the acquisition address
    const getSettingAddress = () => {
      settingRef.value.getAddress()
      settingRef.value.isSettingPublicKey()
    }
    //modify the loading of the set
    const changeLoading = () => {
      settingRef.value.loading = false
    }
    return  {
      getAddress,
      pcHeader,
      settingRef,
      getSettingAddress,
      changeLoading
    }
  }
};
</script>

<style lang="scss" scoped>
//.home-container {
//  height: 100vh;
//}
.body-content {
  //height: calc(100vh - 64px);
  overflow-x: hidden;
  overflow-y: scroll;
  padding: 12px;
  background: #f6f6f9;
}
.header {
  height: 64px;
  background-color: #ffffff;
}
.logo {
  cursor: pointer;
  color: #ffffff;
  height: 20px;
  margin: 32px 27px 32px 27px;
  img {
    height: 100%;
  }
}
.custom-menu {
  height: calc(100vh - 84px);
  overflow-y: scroll;
}
.menu-left-style {
  width: 142px;
  background: #202946;
}
.ant-layout-header {
  padding: 0;
}
</style>
