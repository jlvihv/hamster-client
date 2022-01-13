import {createMemoryHistory, createRouter} from 'vue-router'
import Layout from "../layouts/index"
import Resource from "../views/resource/index"
import Setting from "../views/setting/index"
import Status from "../views/status/index"

const routes = [
  {
    path: '/',
    name: 'Home',
    showLevelFlag: false,
    isShow: false,
    component: Layout,
    redirect: "/resource",
    children: [
      {
        path: '/resource',
        name: 'resource',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "Resource"
        },
        // You can only use pre-loading to add routes, not the on-demand loading method.
        component: Resource
      },
      {
        path: '/setting',
        name: 'setting',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "Setting",
        },
        // You can only use pre-loading to add routes, not the on-demand loading method.
        component: Setting
      },
      {
        path: '/status',
        name: 'status',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "Status"
        },
        // You can only use pre-loading to add routes, not the on-demand loading method.
        component: Status
      }
    ]
  }
]

const router = createRouter({
  mode: "abstract",
  history: createMemoryHistory(),
  routes
})

export default router
