import { useI18n } from '/@/hooks/web/useI18n';

interface DictionaryItem {
  code: number | string;
  children: Array<
    Omit<DictionaryItem, 'children'> & {
      children: DictionaryItem[] | null;
    }
  >;
}

export class DictCode {
  codes: [string, number] | [string];
  value: number | string;

  // Save all code data here
  static data: DictionaryItem[] = [];

  constructor(code: string, subCode?: number) {
    if (subCode != null) {
      const item = DictCode.data.find((x) => x.code === code);
      if (item) item.children.push({ code: subCode, children: [] });

      this.codes = [code, subCode];
      this.value = subCode;
    } else {
      DictCode.data.push({ code, children: [] });

      this.codes = [code];
      this.value = code;
    }
  }

  getLabel() {
    const { t } = useI18n();
    return t(`enums.dictCodeEnum.${this.codes.join('_')}`);
  }

  getOptions() {
    const { t } = useI18n();
    const item = DictCode.data.find((x) => x.code === this.value);

    if (!item) return [];

    return item.children.map((x) => ({
      label: t(`enums.dictCodeEnum.${[this.value, x.code].join('_')}`),
      value: x.code,
    }));
  }

  getOptionLabel(value: number) {
    const options = this.getOptions();
    const option = options.find((x) => x.value === value);

    return option?.label;
  }

  valueOf() {
    return this.value;
  }

  is(value?: string | number | null) {
    if (value == null) {
      return false;
    } else {
      return value === this.value;
    }
  }
}

export const DictCodeEnum = {
  ApplicationDeployStatus: new DictCode('ApplicationDeployStatus'),
  ApplicationDeployStatus_All: new DictCode('ApplicationDeployStatus', 0),
  ApplicationDeployStatus_Running: new DictCode('ApplicationDeployStatus', 1),
  ApplicationDeployStatus_Deploying: new DictCode('ApplicationDeployStatus', 2),
  ApplicationDeployStatus_DeploymentFailed: new DictCode('ApplicationDeployStatus', 3),
  ApplicationDeployStatus_Offline: new DictCode('ApplicationDeployStatus', 4),
  ApplicationQueueStatus: new DictCode('ApplicationQueueStatus'),
  ApplicationQueueStatus_None: new DictCode('ApplicationQueueStatus', 0),
  ApplicationQueueStatus_Running: new DictCode('ApplicationQueueStatus', 1),
  ApplicationQueueStatus_Succeeded: new DictCode('ApplicationQueueStatus', 2),
  ApplicationQueueStatus_Failed: new DictCode('ApplicationQueueStatus', 3),
};
