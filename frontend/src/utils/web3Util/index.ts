import Web3 from 'web3';
import HDWalletProvider from '@truffle/hdwallet-provider';
import configAbi from './abi/config.json';
import ecr20Abi from './abi/ecr20.json';
import stakeDistributionProxyAbi from './abi/stake-distribution-proxy.json';
import stakeProxyFactoryAbi from './abi/stake-proxy-factory.json';

type Web3Type = Web3 & {
  __config?: typeof web3Configs[number];
};

export const web3Abi = {
  configAbi,
  ecr20Abi,
  stakeDistributionProxyAbi,
  stakeProxyFactoryAbi,
};

export const web3Configs = [
  {
    endpoint: 'http://164.92.80.238:8545',
    stakeContractAddress: '0x6189463cfcaD694E0b93e53E29C9070734c25D48',
    factoryContractAddress: '0x75930F4bC45eacc658B6DC111Bbbc664f66449CC',
    erc20ContractAddress: '0xEd2Fed4571597f5b954D4BC212b28422B720b49c',
    configContractAddress: '0x24EA2E2cbc8A0D772dCDFCA304f1b6Fc34fACeCF',
  },
];

export function createWeb3Api(
  endpoint: string,
  mnemonic: string,
  callback?: (api: Web3Type) => any,
): Web3Type {
  const provider = new HDWalletProvider({
    mnemonic: { phrase: mnemonic },
    providerOrUrl: endpoint,
  });

  const api = new Web3(provider);
  const config = web3Configs.find((x) => x.endpoint === endpoint) || web3Configs[0];

  Object.assign(api, { __config: config });

  if (callback) {
    callback(api);
    provider.engine.stop();
  }

  return api;
}

export function getProviderAddress(api: Web3Type) {
  return (api.currentProvider as HDWalletProvider).getAddress(0);
}

export function buildContract(api: Web3Type, abi: Recordable, address: string) {
  return new api.eth.Contract(abi as any, address);
}

export function runContractMethod(options: {
  type: 'send' | 'call';
  api: Web3Type;
  contract: ReturnType<typeof buildContract>;
  method: string;
  methodArgs: any[];
}) {
  const { type = 'send', api, contract, method, methodArgs } = options;
  const ethAddress = getProviderAddress(api);

  return contract.methods[method](...methodArgs)[type]({ from: ethAddress });
}
