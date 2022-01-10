<template>
  <div class="home">
    <a-table
        bordered
        :dataSource="dataSource"
        :columns="columns"
        :pagination="false"
    >
      <template #operation="{ record }">
          <UseResource v-if="record.status===0" :id="record.ID">Use</UseResource>
          <a v-if="record.status===1" @click="deleteUse(record.ID)">Remove</a>
      </template>
    </a-table>
  </div>
</template>

<script>
import {onMounted, ref} from "vue";
import * as Wails from '@wailsapp/runtime';
import UseResource from "@/components/UseResource.vue";
import {message} from "ant-design-vue";
export default {
  name: "Home",
  components: {
    UseResource
  },
  setup() {
    let dataSource = ref([])
    onMounted(() =>{
      Wails.Events.On("resources",t=>{
        dataSource.value = t.map(t => {
          t.key = t.ID
          return t
        })
      })
    })
    function deleteUse(id){
      window.backend.WailsApi.DeleteUseResource(id).then(() => {
        message.success('Succeed');
      });
    }
    return {
      dataSource,
      columns: [
        {
          title: 'resource ID',
          dataIndex: 'ID',
          key: 'ID',
        },
        {
          title: 'resource node ID',
          dataIndex: 'peerId',
          key: 'peerId',
        },
        {
          title: 'cpu core',
          dataIndex: 'cpu',
          key: 'cpu',
        },
        {
          title: 'RAM',
          dataIndex: 'memory',
          key: 'memory',
        },
        {
          title: 'system type',
          dataIndex: 'vmType',
          key: 'vmType',
        },
        {
          title: 'system image',
          dataIndex: 'systemImage',
          key: 'systemImage',
        },
        {
          title: 'founder',
          dataIndex: 'creator',
          key: 'creator',
        },
        {
          title: 'user',
          dataIndex: 'user',
          key: 'user',
        },
        {
          title: 'system status',
          dataIndex: 'status',
          key: 'status',
        },
        {
          title: 'expire time',
          dataIndex: 'expireTime',
          key: 'expireTime',
        },
        {
          title: 'operate',
          dataIndex: 'operation',
          key: 'operation',
          slots: {
            customRender: 'operation',
          },
        },
      ],
      deleteUse
    };
  },
};
</script>
