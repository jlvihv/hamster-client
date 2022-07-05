import { resultSuccess, resultError } from '../../helper';
import * as Setting from '/@wails/go/app/Setting';

const settingInfo: { wsUrl?: string } = {};

export default {
  SettingWsUrl: (wsUrl: string) => {
    if (!wsUrl) {
      return resultError('Failed to save');
    }
    settingInfo.wsUrl = wsUrl;
    return resultSuccess(true);
  },
  GetSetting: () => {
    return settingInfo ? resultSuccess(settingInfo) : resultError('wsUrl record not found');
  },
} as Partial<typeof Setting>;
