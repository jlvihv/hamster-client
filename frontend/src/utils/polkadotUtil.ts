import { ApiPromise, WsProvider, SubmittableResult } from '@polkadot/api';
import { formatBalance as formatBalanceUtil } from '@polkadot/util';
import { Keyring } from '@polkadot/keyring';

// export types
export type { SubmittableResult } from '@polkadot/api';

export function createPolkadotApi(wsUrl: string, callback?: (api: ApiPromise) => any) {
  const apiPromise = ApiPromise.create({ provider: new WsProvider(wsUrl) });

  // If callback provided,
  // auto disconnect when callback process finished
  if (callback) {
    apiPromise.then((api) => {
      callback(api);
      api.disconnect();
    });
  }

  return apiPromise;
}

export function handleTxResults(options: {
  txFailedCb?: (result: SubmittableResult) => any;
  txSuccessCb?: (result: SubmittableResult) => any;
  txUpdateCb?: (result: SubmittableResult) => any;
  unsubscribe: () => any;
}) {
  const { txFailedCb, txSuccessCb, txUpdateCb, unsubscribe } = options;
  return (result: SubmittableResult): void => {
    if (!result || !result.status) {
      return;
    }

    // const status = result.status.type.toLowerCase() as QueueTxStatus;
    // console.log(`${handler}: status :: ${JSON.stringify(result)}`);
    txUpdateCb?.(result);

    if (result.status.isFinalized || result.status.isInBlock) {
      result.events
        .filter(({ event: { section } }) => section === 'system')
        .forEach(({ event: { method } }): void => {
          if (method === 'ExtrinsicFailed') {
            txFailedCb?.(result);
          } else if (method === 'ExtrinsicSuccess') {
            txSuccessCb?.(result);
          }
        });
    } else if (result.isError) {
      txFailedCb?.(result);
    }

    if (result.isCompleted) {
      unsubscribe();
    }
  };
}

// we can override formatBalance in feature,
// we must wrap formatBalanceUtil here to let vite compile to es2015
type FormatBalanceParamsType = Parameters<typeof formatBalanceUtil>;
export function formatBalance(
  mount: FormatBalanceParamsType[0],
  options: FormatBalanceParamsType[1] = {},
) {
  // Decimals 12 is static, we can get it by polkadotApi.registry.chainDecimals
  const defaultOptions = { forceUnit: '-', decimals: 12, withUnit: true };
  return formatBalanceUtil(mount, { ...defaultOptions, ...options });
}

// Check if match for json file and password
export function isJSONAndPasswordMatch(json: any, password: string) {
  const kr = new Keyring({ type: 'sr25519' });
  const krp = kr.addFromJson(json);

  // verify password
  try {
    krp.decodePkcs8(password);
  } catch (error: any) {
    return false;
  }

  return true;
}
