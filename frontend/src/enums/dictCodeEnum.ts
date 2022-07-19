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
  ApplicationDeployStatus_All: new DictCode('ApplicationDeployStatus', 2),
  ApplicationDeployStatus_NotDeployed: new DictCode('ApplicationDeployStatus', 0),
  ApplicationDeployStatus_WaitResource: new DictCode('ApplicationDeployStatus', 3),
  ApplicationDeployStatus_InProgress: new DictCode('ApplicationDeployStatus', 4),
  ApplicationDeployStatus_Deployed: new DictCode('ApplicationDeployStatus', 1),
  ApplicationDeployStatus_DeployFailed: new DictCode('ApplicationDeployStatus', 5),
};
