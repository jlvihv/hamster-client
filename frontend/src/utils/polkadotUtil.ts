import { ApiPromise, WsProvider, SubmittableResult } from '@polkadot/api';
import { formatBalance as formatBalanceUtil } from '@polkadot/util';
import { Keyring } from '@polkadot/keyring';

// export types
export type { SubmittableResult } from '@polkadot/api';

export async function createPolkadotApi(wsUrl: string, callback?: (api: ApiPromise) => any) {
  const api = await ApiPromise.create({ provider: new WsProvider(wsUrl) });

  // If callback provided,
  // auto disconnect when callback process finished
  if (callback) {
    callback(api);
    api.disconnect();
  }

  return api;
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

// Create keyring pair by json file and password
export function createKeyPair(json: any, password: string) {
  const kr = new Keyring({ type: 'sr25519' });
  const krp = kr.addFromJson(json);

  // verify password
  try {
    krp.decodePkcs8(password);
  } catch (error: any) {
    return false;
  }

  return krp;
}

// Cancel resource order
export async function cancelResourceOrder(api, keyringPair) {
  const orderNumber = await api.query.resourceOrder.applyUsers(keyringPair.address);

  console.log('orderNumber', orderNumber);

  if (orderNumber.toJSON() > 0) {
    const unsubscribe = api.tx.resourceOrder
      .releaseApplyFreeResource(orderNumber.toJSON())
      .signAndSend(
        keyringPair,
        handleTxResults({
          txSuccessCb: (result) => {
            console.log(result);
          },
          unsubscribe: () => unsubscribe(),
        }),
      );
  }
}

// Apply resource order
export async function applyResourceOrder(
  api,
  keyringPair,
  options: { cpu?: number; memory?: number; leaseTerm: number; publicKey: string; type?: number },
  callback?: typeof handleTxResults,
) {
  const { cpu = 4, memory = 8, leaseTerm, publicKey, type = 1 } = options;

  // Cancel first
  await cancelResourceOrder(api, keyringPair);

  return api.tx.resourceOrder
    .applyFreeResource(cpu, memory, leaseTerm, publicKey, type)
    .signAndSend(keyringPair, callback);
}
