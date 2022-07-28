import { ApiPromise, WsProvider, SubmittableResult } from '@polkadot/api';
import { formatBalance as formatBalanceUtil } from '@polkadot/util';
import { Keyring } from '@polkadot/keyring';

// export types
export type { SubmittableResult } from '@polkadot/api';

export async function createPolkadotApi(wsUrl: string, callback?: (api: ApiPromise) => any) {
  const api = new ApiPromise({ provider: new WsProvider(wsUrl) });
  await api.isReadyOrError.catch((e) => {
    console.log('connect error', e);
    api.disconnect();
  });
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
  return new Promise((resolve, reject) => {
    if (orderNumber.toJSON() > 0) {
      const unsubscribe = api.tx.resourceOrder
        .releaseApplyFreeResource(orderNumber.toJSON())
        .signAndSend(keyringPair, (result) => {
          if (result.status.isInBlock) {
            console.log(`Transaction included at blockHash ${result.status.asInBlock}`);
            resolve(result);
          } else if (result.status.isFinalized) {
            console.log(`Transaction finalized at blockHash ${result.status.asFinalized}`);
            result.events
              .filter(({ event }) => api.events.system.ExtrinsicFailed.is(event))
              .forEach(({ event }) => {
                const [error] = event.data;
                const decoded = api.registry.findMetaError(result.dispatchError.asModule);
                const { docs, method, section } = decoded;
                console.log(`${section}.${method}: ${docs.join(' ')}`);
                reject(error);
              });
          }

          if (result.isCompleted) {
            unsubscribe();
          }
        });
    } else {
      resolve(undefined);
    }
  });
}

// Apply resource order
export function applyResourceOrder(
  api,
  keyringPair,
  options: { cpu?: number; memory?: number; leaseTerm: number; publicKey: string; type?: number },
) {
  return new Promise((resolve, reject) => {
    const { cpu = 4, memory = 8, leaseTerm, publicKey, type = 1 } = options;

    const unsubscribe = api.tx.resourceOrder
      .applyFreeResource(cpu, memory, leaseTerm, publicKey, type)
      .signAndSend(keyringPair, (result) => {
        if (result.status.isInBlock) {
          console.log(`Transaction included at blockHash ${result.status.asInBlock}`);
          resolve(result);
        } else if (result.status.isFinalized) {
          console.log(`Transaction finalized at blockHash ${result.status.asFinalized}`);
          result.events
            .filter(({ event }) => api.events.system.ExtrinsicFailed.is(event))
            .forEach(({ event }) => {
              const [error] = event.data;
              const decoded = api.registry.findMetaError(result.dispatchError.asModule);
              const { docs, method, section } = decoded;
              console.log(`${section}.${method}: ${docs.join(' ')}`);
              reject(error);
            });
        }

        if (result.isCompleted) {
          unsubscribe();
        }
      })
      .catch((e) => {
        reject(e);
      });
  });
}
