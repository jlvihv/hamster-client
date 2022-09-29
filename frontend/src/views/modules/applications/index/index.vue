<template>
  <PageWrapper>
    <Header />
    <div class="grid grid-cols-3 gap-8 mx-3 mt-3 text-center">
      <router-link
        to="/applications/select"
        class="rounded-[20px] h-full border-2 border-[#043CC1] border-dashed flex items-center justify-center cursor-pointer transition duration-500 hover:scale-110 min-h-[300px]"
      >
        <SvgIcon class="text-primary rounded-[50%]" size="60" name="add" />
      </router-link>
      <div
        class="bg-white rounded-[20px] relative duration-500 hover:scale-110 cursor-pointer"
        v-for="(item, index) in applications"
        :key="item.id"
        @click="$router.push(`/applications/${item.id}`)"
      >
        <div
          class="text-[#2E3C43] text-[12px] px-[10px] absolute right-0 top-[20px] h-[20px] rounded-l-[100px] bg-color"
        >
          {{ item.selectNodeType }}
        </div>
        <img
          :src="getImageURL(`application-bg-${(index % 4) + 1}.png`)"
          class="w-full rounded-t-[20px]"
        />
        <div class="text-[20px] font-bold my-[10px]">{{ item.name }}</div>
        <Tooltip>
          <template #title v-if="item.grtIncome == 0 ? false : true">
            <span>{{ item.grtIncome }}</span>
          </template>
          <div class="text-[18px]">
            <SvgIcon class="text-primary" size="20" name="grt" />
            {{ formatIncome(Number(item.grtIncome)) }} GRT
          </div>
        </Tooltip>
        <div class="text-[#6A7EAF] text-[18px] mt-[10px] mb-[50px]">
          {{ DictCodeEnum.ApplicationDeployStatus.getOptionLabel(item.status) }}
        </div>
      </div>
    </div>
    <div class="text-center my-[40px]" v-if="!isTouchedEnd && !isFromGuide">
      <Button
        class="!h-[60px] w-[200px]"
        size="large"
        type="primary"
        shape="round"
        @click="loadApplications"
        :loading="isLoading"
      >
        {{ t('common.moreText') }}
      </Button>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { onMounted } from 'vue';
  import { useRoute } from 'vue-router';
  import { PageWrapper } from '/@/components/Page';
  import { SvgIcon } from '/@/components/Icon';
  import Header from './components/Header.vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useLoadMore } from '/@/hooks/web/useLoadMore';
  import { ApplicationList } from '/@wails/go/app/Application';
  import { DictCodeEnum } from '/@/enums/dictCodeEnum';
  import { useAssets } from '/@/hooks/web/useAssets';
  import { Button, Tooltip } from 'ant-design-vue';
  import { formatIncome } from '/@/utils/thegraphUtil/grt';

  const { t } = useI18n();
  const route = useRoute();
  const { getImageURL } = useAssets();

  // check if redirect from home(guide page)
  const isFromGuide = route.query.from === 'guide';

  const {
    items: applications,
    isLoading,
    isTouchedEnd,
    loadMore: loadApplications,
  } = useLoadMore(
    (page, perPage) =>
      ApplicationList(page, perPage, undefined, DictCodeEnum.ApplicationDeployStatus_All.value),
    {
      responseHandler: (data) => data.items,
      perPage: 20,
    },
  );

  onMounted(loadApplications);
</script>

<style lang="less" scoped>
  .bg-color {
    background: rgba(255, 255, 255, 0.6);
  }
</style>
