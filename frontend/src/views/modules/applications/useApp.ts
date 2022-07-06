import { computed } from 'vue';
import { useI18n } from '/@/hooks/web/useI18n';

const { t } = useI18n();

// option data
export function useOptionsContent() {
  const statusOptions = computed(() => [
    { value: 2, label: t('applications.index.all') },
    { value: 0, label: t('applications.index.undeployed') },
    { value: 1, label: t('applications.index.deployed') },
  ]);

  return { statusOptions };
}
