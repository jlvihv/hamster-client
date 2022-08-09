<template>
  <PageWrapper>
    <Header />
    <div class="grid grid-cols-3 gap-8 mt-3 mx-3 text-center">
      <div
        class="bg-white rounded-[20px] relative duration-500 hover:scale-110"
        v-for="(item, index) in [1, 2, 3, 4, 5]"
        :key="index"
      >
        <div
          class="text-[#2E3C43] text-[12px] px-[10px] absolute right-0 top-[20px] h-[20px] rounded-l-[100px] bg-color"
        >
          Thegraph Test
        </div>
        <img
          :src="`src/assets/images/application-bg-${(index % 4) + 1}.png`"
          class="w-full rounded-t-[20px]"
        />
        <div class="text-[20px] font-bold my-[10px]">Thegraph0{{ item }}</div>
        <div class="text-[18px]">
          <SvgIcon class="text-primary" size="20" name="grt" />200000000000 GRT
        </div>
        <div class="text-[#6A7EAF] text-[18px] mt-[10px] mb-[50px]">Running</div>
      </div>
      <router-link to="/applications/new">
        <div class="rounded-[20px] h-full border-2 border-[#043CC1] border-dashed flex items-center justify-center cursor-pointer duration-500 hover:scale-110">
          <SvgIcon class="text-primary rounded-[50%]" size="60" name="add" />
        </div>
      </router-link>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { reactive, onMounted } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { SvgIcon } from '/@/components/Icon';
  import Header from './components/Header.vue';
  import { ApplicationList } from '/@wails/go/app/Application';

  const applications = reactive([]);
  const pagination = reactive({ page: 1, pageSize: 20 });

  const loadApplications = async () => {
    const { items } = await ApplicationList(pagination.page, pagination.pageSize);

    applications.push(...items);
  };

  onMounted(loadApplications);
</script>

<style lang="less" scoped>
  .bg-color {
    background: rgba(255, 255, 255, 0.6);
  }
</style>
