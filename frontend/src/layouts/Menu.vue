<template>
  <a-menu
    id="menu"
    theme="dark"
    v-model:selectedKeys="selectedKeys"
    v-model:openKeys="openKeys"
    mode="inline"
    @click="handleClick"
  >
    <div class="menu-title-style">
      <div v-for="(item,key) in routers" :key="key">
        <a-menu-item v-if="item.isShow && !item.showLevelFlag" :key="item.name">
          <div class="menu-content">
            <span class="menu-content-front">{{
                item.meta.title
              }}</span>
          </div>
        </a-menu-item>
      </div>
    </div>
  </a-menu>


</template>

<script>
import {watch, reactive, ref, toRefs} from "vue";
import {useRouter,useRoute} from "vue-router";
import routers from "../router/index";
export default {
  name: "Menu",
  setup() {
    const currentRoute = useRoute()
    const router = useRouter();
    // const selectedKeys = ref(["computingPowerMarket"]);
    const selectedKeys = ref([currentRoute.name]);

    const openKeys = ref([currentRoute.matched[1]?.name])

    const state = reactive({
      routers: routers.options.routes[0].children,
    });
    const useRouterCurrent = reactive(useRouter());
    watch(useRouterCurrent, o => {
      openKeys.value = [currentRoute.matched[1]?.name]
      if (o.currentRoute.fullPath.indexOf("resource") != -1) {
        selectedKeys.value = ["resource"]
      }else if (o.currentRoute.fullPath.indexOf("setting") != -1) {
        selectedKeys.value = ["setting"]
      }else if (o.currentRoute.fullPath.indexOf("status") != -1) {
        selectedKeys.value = ["status"]
      } else {
        selectedKeys.value = [currentRoute.name]
      }
    });
    const handleClick = (e) => {
      console.log(e);
      router.push({ name: e.key });
    };
    return {
      ...toRefs(state),
      selectedKeys,
      openKeys,
      useRouterCurrent,
      handleClick
    };
  },
};
</script>

<style lang="scss" scoped>
.ant-menu {
  padding-left: 32px;
  line-height: normal;
  .title {
    margin-left: 10px;
  }
}
.menu-title {
  height: 41px;
  line-height: 32px;
  margin-bottom: 12px;
  display: flex;
  align-items: center;
  .menu-front-selected {
    font-size: 12px;
    font-weight: 400;
    color: #fdfeff;
    margin-left: 8px;
  }
  .menu-front {
    font-size: 12px;
    font-weight: 400;
    color: $font-color-third-content;
    margin-left: 8px;
  }
}
.menu-content-front {
  margin-left: 8px;
}
.menu-content {
  display: flex;
  align-items: center;
}
.menu-title-style {
  padding-bottom: 20px;
}
</style>
