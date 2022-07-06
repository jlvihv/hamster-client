import { defineStore } from 'pinia';
import { store } from '/@/store';
import { GetWalletInfo, SaveWallet, DeleteWallet } from '/@wails/go/app/Wallet';
import { GetSetting } from '/@wails/go/app/Setting';

interface SettingState {
  walletInfo?: { address: string; address_json: string };
  config?: Recordable;
}

export const useSettingStore = defineStore({
  id: 'app-setting',
  state: (): SettingState => ({
    walletInfo: undefined,
    config: undefined,
  }),
  getters: {},
  actions: {
    async getWalletInfoAction() {
      const wallet = await GetWalletInfo();
      if (!(wallet instanceof Error)) this.walletInfo = wallet;
    },
    async saveWalletAction(address: string, json: string) {
      await SaveWallet(address, json);
      await this.getWalletInfoAction();
    },
    async deleteWalletAction() {
      await DeleteWallet();
      this.walletInfo = undefined;
    },
    async getConfigAction() {
      const config = await GetSetting();
      if (!(config instanceof Error)) this.config = config;
    },
  },
});

// Need to be used outside the setup
export function useSettingStoreWithOut() {
  return useSettingStore(store);
}
