<template>
  <div class="grid grid-cols-4 gap-8 text-center mt-3 mx-3">
    <div
      class="bg-white rounded-[8px] py-[20px] px-[40px] duration-500 hover:scale-110"
      v-for="item in subgraphs"
      :key="item.id"
    >
      <img :src="item.image" class="w-full rounded-[8px]" />
      <div class="text-[14px] text-[#BFC6D4] mt-[6px]">{{ shortenAddress(item.owner.id) }}</div>
      <div class="text-[#222222] mt-[6px] font-bold truncate">{{ item.displayName }}</div>
      <div class="text-[14px] text-[#BFC6D4] mt-[6px]">
        Signal: {{ formatfromWei(item.currentSignalledTokens) }} GRT
      </div>
      <Button
        class="text-[#63A0FA] text-[14px] mt-[20px] border border-[#63A0FA] rounded-[4px] h-[30px] !min-w-[100px]"
        :loading="subgraphDeployLoading[getSubgraphIpfsHash(item)]"
        :disabled="deployedSubgraphIdentifiers.includes(getSubgraphIpfsHash(item))"
        @click="handleDeploySubgraph(item)"
      >
        {{
          deployedSubgraphIdentifiers.includes(getSubgraphIpfsHash(item))
            ? t('applications.see.deployed')
            : t('applications.see.start')
        }}
      </Button>
    </div>
  </div>
  <div class="text-center my-[40px]" v-if="!isTouchedEnd">
    <Button
      class="!h-[60px] w-[200px]"
      size="large"
      type="primary"
      shape="round"
      @click="loadSubgraphList"
      :loading="isSubgraphsLoading"
    >
      {{ t('common.moreText') }}
    </Button>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, reactive, ref, computed, toRefs } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useLoadMore } from '/@/hooks/web/useLoadMore';
  import { createSubgraphClient, fetchSubgraphs } from '/@/utils/thegraphUtil/subgraph';
  import { shortenAddress, pluginConfigs } from '/@/utils/thegraphUtil';
  import { GraphStart, GraphRules } from '/@wails/go/app/Graph';
  import { formatfromWei } from '/@/utils/web3Util';
  import { Button } from 'ant-design-vue';

  const props = defineProps({
    application: Object as PropType<Recordable>,
  });
  const { application } = toRefs(props);

  const { t } = useI18n();

  const deployedSubgraphs = ref([]);
  const deployedSubgraphIdentifiers = computed(() =>
    deployedSubgraphs.value.map((x) => x.identifier),
  );
  const subgraphDeployLoading = reactive<Record<string, boolean>>({});

  const getSubgraphIpfsHash = (item: any) => item.currentVersion.subgraphDeployment.ipfsHash;
  const handleDeploySubgraph = async (item: any) => {
    const deploymentId = getSubgraphIpfsHash(item);

    subgraphDeployLoading[deploymentId] = true;

    try {
      await GraphStart(application.value.id, deploymentId);
      await loadDeployedSubgraphs();
    } catch (error: any) {
      console.log('Deployed Failed', deploymentId);
    } finally {
      subgraphDeployLoading[deploymentId] = false;
    }
  };

  // client for get subgraph lists
  const nodeType = application.value.selectNodeType;
  const plugin =
    pluginConfigs.find(({ value }) => value === nodeType) ||
    pluginConfigs.find(({ value }) => value === 'thegraph_rinkeby');
  const listClient = createSubgraphClient(plugin.url);

  const {
    isTouchedEnd,
    items: subgraphs,
    isLoading: isSubgraphsLoading,
    loadMore: loadSubgraphList,
  } = useLoadMore((page, perPage) => fetchSubgraphs(listClient, page, perPage), {
    responseHandler: (data) => data.data.subgraphs,
    perPage: 20,
  });

  const loadDeployedSubgraphs = async () => {
    // Get deployed subgraphs
    const { info } = await GraphRules(application.value.id);
    deployedSubgraphs.value = info;
  };

  onMounted(() => {
    loadSubgraphList();
    loadDeployedSubgraphs();
  });
</script>
