<template>
  <PageWrapper>
    <div class="search-header">
      <Card :title="t('applications.index.appList')">
        <Form layout="inline" v-model:model="searchForm">
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
      <Card class="application-table-card">
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
                :title="t('applications.index.delAppInfo')"
                @confirm="deleteApp(index, record.id)"
              >
                <a class="mr-3 text-red-600 hover:text-red-600" :title="t('common.delText')"
                  ><DeleteOutlined
                /></a>
              </Popconfirm>
            </template>
          </template>
        </Table>
      </Card>
    </div>
    <Modal
      v-model:visible="visible"
      :title="[
        operateType === 'edit' ? t('applications.index.editApp') : t('applications.index.addApp'),
      ]"
      :okText="t('common.confirmText')"
      @ok="handleOk"
    >
      <Form ref="formRef" :rules="formRules" v-model:model="formData" :label-col="labelCol">
        <FormItem :label="t('applications.index.nameText')" name="name">
          <Input :allowClear="true" class="input-width" v-model:value="formData.name" />
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
  import { useOptionsContent } from '.././useApp';
  import { PageWrapper } from '/@/components/Page';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import {
    ApplicationList,
    AddApplication,
    UpdateApplication,
    DeleteApplication,
  } from '/@wails/go/app/Application';
  import { PlusOutlined, FormOutlined, DeleteOutlined, EyeOutlined } from '@ant-design/icons-vue';
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
  const { notification, createErrorModal } = useMessage();
  const { statusOptions } = useOptionsContent();
  const searchForm = reactive({
    status: 2, //Application status
    name: '', //Application name
  });

  const visible = ref(false);
  const operateType = ref('add');
  // Form data
  const formRef = ref();
  const formData = reactive<{ id?: number; name: string; describe: string }>({
    name: '', //Application name
    describe: '', //Application description
  });
  // Form rules
  const formRules = computed(() => ({
    name: [{ message: t('applications.index.nameText'), trigger: 'change', required: true }],
  }));

  const loading = ref(false);
  const tableData = ref([]);
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
      customRender: ({ text: date }) => formatToDateTime(date, (f) => f.datetimeWithoutSec),
    },
    {
      title: t('applications.index.statusText'),
      dataIndex: 'status',
      align: 'center',
      ellipsis: 'fixed',
      key: 'status',
      customRender: ({ text }) =>
        statusOptions.value.find((option) => option.value === text)?.label,
    },
    {
      title: t('applications.index.operText'),
      dataIndex: 'action',
      align: 'center',
      width: '200px',
    },
  ]);

  const pagination = reactive({
    // Paging configurator
    pageSize: 10, // One-page data limit
    current: 1, // Current page
    total: 10, // Total
    hideOnSinglePage: false, // Whether to hide the paginator when only one page is available
    showQuickJumper: true, // Is it possible to jump quickly to a page
    showSizeChanger: true, // Is it possible to change the pageSize
    pageSizeOptions: ['10', '20', '30'], // Specify how many items can be displayed per page
    onShowSizeChange: (current, pagesize) => {
      // Callback when changing pageSize
      pagination.current = current;
      pagination.pageSize = pagesize;
      getAppList();
    },
    onChange: (current) => {
      // callback when switching paging.
      pagination.current = current;
      getAppList();
    },
    // showTotal: total => `total：${total}`, // Total number of displays possible
  });

  async function deleteApp(index, id) {
    try {
      await DeleteApplication(id);
      tableData.value.splice(index, 1);
      setMessageInfo('suc');
    } catch (error: any) {
      setMessageInfo('error');
    }
  }

  onMounted(async () => {
    getAppList();
  });

  async function getAppList() {
    loading.value = true;
    try {
      const page = pagination.current;
      const pageSize = pagination.pageSize;
      const result = await ApplicationList(page, pageSize, searchForm.name, searchForm.status);
      pagination.total = result.total;
      tableData.value = result.items;
    } catch (error: any) {
      setMessageInfo('error');
    } finally {
      loading.value = false;
    }
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
      searchForm.status = 2;
      getAppList();
    },
    async onSearch() {
      pagination.current = 1;
      getAppList();
    },
  };

  async function handleOk() {
    await formRef.value.validate();
    try {
      if (operateType.value === 'add') {
        const result = await AddApplication(formData);
        setApplicationReload(result);
      } else {
        const result = await UpdateApplication(formData);
        setApplicationReload(result);
      }
    } catch (error: any) {
      setMessageInfo('error');
    } finally {
      visible.value = false;
    }
  }
  async function setApplicationReload(result) {
    if (result === true) {
      setMessageInfo('suc');
      getAppList();
    }
  }
  async function addApplication() {
    resetForm(formData);
    operateType.value = 'add';
    visible.value = true;
  }
  async function editApplication(data) {
    const { id, name, describe } = data;
    Object.assign(formData, { id, name, describe });

    operateType.value = 'edit';
    visible.value = true;
  }
  async function setMessageInfo(infoType) {
    if (infoType == 'error') {
      createErrorModal({
        title: t('common.errorTip'),
        content: t('common.operateFailText'),
      });
    } else {
      notification.success({
        message: t('common.operateSucText'),
        duration: 3,
      });
    }
  }

  const labelCol = reactive({
    style: { width: '90px' },
  });
</script>

<style lang="less" scoped>
  :deep(.input-width) {
    width: 180px !important;
  }

  .application-table-card {
    :deep(.ant-card-body) {
      @apply !px-0;
    }

    :deep(.ant-table-pagination.ant-pagination) {
      @apply !mx-4;
    }
  }
</style>
