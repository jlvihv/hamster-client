export const pluginConfigs = [
  {
    label: 'Thegraph(Mainnet)',
    value: 'thegraph_mainnet',
    url: 'https://gateway.thegraph.com/network',
  },
  {
    label: 'Thegraph(Rinkeby)',
    value: 'thegraph_rinkeby',
    url: 'https://api.thegraph.com/subgraphs/name/graphprotocol/graph-network-testnet',
  },
  { label: 'Thegraph(Görli)', value: 'thegraph_gorli', disabled: true },
  { label: 'Thegraph(Ethereum)', value: 'thegraph_ethereum', disabled: true },
  { label: 'Polygon(Ethereum)', value: 'polygon_ethereum', disabled: true },
];

export const deployData = {
  nodeEthereumUrl: 'mainnet:https://main-light.eth.linkpool.io',
  ethereumUrl: 'https://rinkeby.infura.io/v3/62d7b5f33ae443e784919f1c2afe24a3',
  ethereumNetwork: 'rinkeby',
  indexerAddress: '',
};

// Make address shorter,
// Example: "0xd5f6e31199220a0d5334cad2b6ecd70c8f1a6b79" => "0xd5f6—1a6b79"
export function shortenAddress(address: string, digit = 6) {
  return address.slice(0, digit) + '—' + address.slice(-digit);
}
