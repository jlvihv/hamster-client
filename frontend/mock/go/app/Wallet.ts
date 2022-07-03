import { resultSuccess, resultError } from '../../helper';
import * as Wallet from '/@wails/go/app/Wallet';

let walletInfo: { address: string; address_json: string } | undefined;

export default {
  GetWalletInfo: () => {
    return walletInfo ? resultSuccess(walletInfo) : resultError('Wallet record not found');
  },
  DeleteWallet: () => {
    walletInfo = undefined;
    return resultSuccess(walletInfo);
  },
  SaveWallet: (address: string, addressJson: string) => {
    walletInfo = { address, address_json: addressJson };
    return resultSuccess(walletInfo);
  },
} as Partial<typeof Wallet>;
