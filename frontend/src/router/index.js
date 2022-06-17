import {createMemoryHistory, createRouter} from 'vue-router'
import Layout from "../layouts/index"
import Resource from "../views/resource/index"
import Setting from "../views/setting/index"
import Status from "../views/status/index"
import ResourceMarket from "@/views/resourceMarket";
import OrderList from "@/views/orderList";
import MyResource from "@/views/myResource";
import Websocket from "@/views/websocket"

const routes = [
  {
    path: '/',
    name: 'Home',
    showLevelFlag: false,
    isShow: false,
    component: Layout,
    redirect: "/my-resource",
    children: [
      {
        path: '/resource-market',
        name: 'resource-market',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: 'Resource Market'
        },
        component: ResourceMarket
      },
      {
        path: '/my-resource',
        name: 'my-resource',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "My Resource"
        },
        // You can only use pre-loading to add routes, not the on-demand loading method.
        component: MyResource
      },
      {
        path: '/order-list',
        name: 'order-list',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "Order List"
        },
        component: OrderList
      },
      {
        path: '/link',
        name: 'link',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "Link",
        },
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
      },
      {
        path: '/Deploy',
        name: 'Deploy',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: 'Deploy',
        },
        component: Websocket
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
