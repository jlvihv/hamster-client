<template>
  <PageWrapper>
    <div class="search-header">
      <Card :title="t('applications.index.appList')">
        <Form layout="inline" v-model:model="searchForm" :label-col="labelCol">
          <FormItem :label="t('applications.index.nameText')" name="name">
            <Input :allowClear="true" class="input-width" v-model:value="searchForm.name" />
          </FormItem>
          <FormItem :label="t('applications.index.statusText')" name="status">
            <Select
              class="input-width"
              v-model:value="searchForm.status"
              :options="statusOptions"
            />
          </FormItem>
          <FormItem>
            <Button type="primary" @click="searchAction.onSearch">{{
              t('common.searchText')
            }}</Button>
            <Button class="ml-4" @click="searchAction.onReset">{{ t('common.resetText') }}</Button>
          </FormItem>
        </Form>
      </Card>
    </div>
    <div class="mt-4">
      <Card class="table-card">
        <div class="mx-4">
          <Button type="primary" ghost @click="addApplication"
            ><PlusOutlined class="!inline-flex" />Add</Button
          >
        </div>
        <Table
          :loading="loading"
          class="my-4"
          :row-key="(record) => record.id"
          :columns="tableColumns"
          :dataSource="tableData"
          :pagination="pagination"
        >
          <template #bodyCell="{ column, record, index }">
            <template v-if="column.dataIndex === 'action'">
              <router-link
                class="mr-3"
                :to="`/applications/${record.id}`"
                :title="t('common.lookText')"
                ><EyeOutlined
              /></router-link>
              <a
                class="mr-3 text-[#1890FF]"
                :title="t('common.editText')"
                @click="editApplication(record)"
                ><FormOutlined
              /></a>
              <Popconfirm
                :title="t('application.index.delAppInfo')"
                @confirm="tableAction.deleteApp(index, record.id)"
              >
                <a class="mr-3 text-red-600 hover:text-red-600" :title="t('common.delText')"
                  ><DeleteOutlined
                /></a>
              </Popconfirm>
              <Popconfirm
                :title="t('applications.index.disabledAppInfo')"
                @confirm="tableAction.changeStatus(index, record.id, record.status)"
              >
                <a class="mr-3" :title="t('common.disabledText')"><DisconnectOutlined /></a>
              </Popconfirm>
            </template>
          </template>
        </Table>
      </Card>
    </div>
    <!-- <appModal></appModal> -->
    <Modal
      v-model:visible="visible"
      :title="[
        formData.name != '' ? t('applications.index.editApp') : t('applications.index.addApp'),
      ]"
      @ok="handleOk"
    >
      <Form ref="formRef" :rules="formRules" v-model:model="formData" :label-col="labelCol">
        <FormItem :label="t('applications.index.nameText')" name="name">
          <Input :allowClear="true" class="input-width" v-model:value="formData.name" />
        </FormItem>
        <FormItem :label="t('applications.index.abbText')" name="abbreviation">
          <Input :allowClear="true" class="input-width" v-model:value="formData.abbreviation" />
        </FormItem>
        <FormItem :label="t('applications.index.desText')" name="describe">
          <Textarea :allowClear="true" v-model:value="formData.describe" />
        </FormItem>
      </Form>
    </Modal>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { reactive, computed, ref, onMounted } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    PlusOutlined,
    FormOutlined,
    DeleteOutlined,
    EyeOutlined,
    DisconnectOutlined,
  } from '@ant-design/icons-vue';
  import {
    Card,
    Button,
    Form,
    FormItem,
    Input,
    Select,
    Textarea,
    Table,
    Modal,
    Popconfirm,
  } from 'ant-design-vue';

  const { t } = useI18n();
  const statusOptions = reactive([]);
  const searchForm = reactive({
    status: '', //应用状态
    name: '', //应用名称
  });

  const visible = ref(false);
  // Form data
  const formRef = ref();
  const formData = reactive({
    name: '', //应用名称
    abbreviation: '', //应用简写
    describe: '', //应用描述
  });
  // Form rules
  const formRules = computed(() => ({
    name: [{ message: t('applications.index.nameText'), trigger: 'change', required: true }],
  }));

  const loading = ref(false);
  const tableData = ref([
    {
      id: 2,
      createdAt: '2022-7-4',
      name: 'name',
    },
  ]);
  const tableColumns = computed<any[]>(() => [
    {
      title: t('applications.index.noText'),
      dataIndex: 'index',
      key: 'index',
      align: 'center',
      width: '70px',
      customRender: ({ index }) => index + 1 + (pagination.current - 1) * pagination.pageSize,
    },
    {
      title: t('applications.index.nameText'),
      dataIndex: 'name',
      align: 'center',
      ellipsis: 'fixed',
      key: 'name',
    },
    {
      title: t('applications.index.abbText'),
      dataIndex: 'abbreviation',
      align: 'center',
      ellipsis: 'fixed',
      key: 'abbreviation',
    },
    {
      title: t('applications.index.desText'),
      dataIndex: 'describe',
      align: 'center',
      ellipsis: 'fixed',
      key: 'describe',
    },
    {
      title: t('applications.index.addTimeText'),
      dataIndex: 'createdAt',
      key: 'createdAt',
      ellipsis: 'fixed',
      align: 'center',
      // customRender: ({ text: date }) => formatToDateTime(date, (f) => f.datetimeWithoutSec),
    },
    {
      title: t('applications.index.statusText'),
      dataIndex: 'status',
      align: 'center',
      ellipsis: 'fixed',
      key: 'status',
      // customRender: ({ text }) => dictStore.getOptionLabel(DictCodeEnum.ApplicationStatus, text),
    },
    {
      title: t('applications.index.operText'),
      dataIndex: 'action',
      align: 'center',
      width: '200px',
    },
  ]);

  const pagination = reactive({
    // 分页配置器
    pageSize: 10, // 一页的数据限制
    current: 1, // 当前页
    total: 10, // 总数
    hideOnSinglePage: false, // 只有一页时是否隐藏分页器
    showQuickJumper: true, // 是否可以快速跳转至某页
    showSizeChanger: true, // 是否可以改变 pageSize
    pageSizeOptions: ['10', '20', '30'], // 指定每页可以显示多少条
    onShowSizeChange: (current, pagesize) => {
      // 改变 pageSize时的回调
      pagination.current = current;
      pagination.pageSize = pagesize;
      getAppList();
    },
    onChange: (current) => {
      // 切换分页时的回调，
      pagination.current = current;
      getAppList();
    },
    // showTotal: total => `总数：${total}人`, // 可以展示总数
  });

  const tableAction = {
    //启用/禁用应用状态
    async changeStatus(index, id, status) {
      console.log(index, id, status);
      // try {
      //   const newStatus = ref(DictCodeEnum.ApplicationStatus_Inactive.value); //禁用
      //   if (DictCodeEnum.ApplicationStatus_Inactive.is(status)) {
      //     newStatus.value = DictCodeEnum.ApplicationStatus_Active.value;
      //   }
      //   const result = await changeStatusApi({ id: id, status: newStatus.value });
      //   tableData.value[index].status = result.status;
      //   //设置提示信息
      //   setMessageInfo('suc');
      // } catch (error: any) {
      //   //设置提示信息
      //   setMessageInfo('error');
      // }
    },

    async deleteApp(index, id) {
      console.log(index, id);
      // try {
      //   await deleteAppApi(id);
      //   tableData.value.splice(index, 1);
      //   //设置提示信息
      //   setMessageInfo('suc');
      // } catch (error: any) {
      //   //设置提示信息
      //   setMessageInfo('error');
      // }
    },
  };

  onMounted(async () => {
    getAppList();
  });

  async function getAppList() {
    // loading.value = true;
    // try {
    //   const page = pagination.current;
    //   const pageSize = pagination.pageSize;
    //   const result = await getAppListApi({ ...formData, page, pageSize });
    //   pagination.pageSize = result.pageSize;
    //   pagination.current = result.page;
    //   pagination.total = result.total;
    //   tableData.value = result.list;
    // } catch (error: any) {
    //   setMessageInfo('error');
    // } finally {
    //   loading.value = false;
    // }
  }

  async function resetForm(resetData) {
    const keys = Object.keys(resetData);
    let obj: { [name: string]: string } = {};
    keys.forEach((item) => {
      obj[item] = '';
    });
    Object.assign(resetData, obj);
  }

  const searchAction = {
    async onReset() {
      resetForm(searchForm);
    },
    async onSearch() {
      console.log('format:', formData);
    },
  };

  async function handleOk() {
    await formRef.value.validate();
    console.log('handle Ok ', formData);
  }
  async function addApplication() {
    resetForm(formData);
    visible.value = true;
  }
  async function editApplication(data) {
    resetForm(formData);
    Object.assign(formData, data);
    visible.value = true;
  }

  const labelCol = reactive({
    style: { width: '90px' },
  });
</script>
<style lang="less" scope>
  .search-header .input-width {
    width: 180px !important;
  }

  .table-card {
    .ant-card-body {
      @apply !px-0;
    }

    .ant-table-pagination.ant-pagination {
      @apply !mx-4;
    }
  }
</style>
