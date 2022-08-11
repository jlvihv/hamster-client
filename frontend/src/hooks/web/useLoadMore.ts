import { reactive, ref } from 'vue';

export function useLoadMore(
  apiFunc: Function,
  options?: { responseHandler?: Function; perPage?: number },
) {
  const page = ref(1);
  const perPage = options?.perPage || 20;

  const items = reactive<any[]>([]);
  const isLoading = ref(false);
  const isTouchedEnd = ref(false);

  const loadMore = async (event: any) => {
    if (isLoading.value) return;

    if (event) {
      page.value++;
    }

    isLoading.value = true;

    const data = await apiFunc(page.value, perPage);
    const newItems: any[] = options?.responseHandler ? options.responseHandler(data) : data;
    items.push(...newItems);

    if (newItems.length < perPage) {
      isTouchedEnd.value = true;
    }

    isLoading.value = false;
  };

  return { page, items, isLoading, isTouchedEnd, loadMore };
}
