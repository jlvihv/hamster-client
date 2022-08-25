export const pluginConfigs = [
  {
    label: 'Thegraph(Mainnet)',
    value: 'thegraph_mainnet',
    url: 'https://gateway.thegraph.com/network',
  },
  {
    label: 'Thegraph(Görli)',
    value: 'thegraph_gorli',
    url: 'https://api.thegraph.com/subgraphs/name/graphprotocol/graph-network-goerli',
  },
  {
    label: 'Thegraph(Rinkeby)',
    value: 'thegraph_rinkeby',
    url: 'https://api.thegraph.com/subgraphs/name/graphprotocol/graph-network-testnet',
    disabled: true,
  },
  { label: 'Polygon(Ethereum)', value: 'polygon_ethereum', disabled: true },
];

export const deployData = {
  nodeEthereumUrl: 'mainnet:https://main-light.eth.linkpool.io',
  ethereumUrl: 'https://goerli.infura.io/v3/bab2a1a435b04c07a488d847cf6788f7',
  ethereumNetwork: 'goerli',
  indexerAddress: '',
};

// Make address shorter,
// Example: "0xd5f6e31199220a0d5334cad2b6ecd70c8f1a6b79" => "0xd5f6—1a6b79"
export function shortenAddress(address: string, digit = 6) {
  return address.slice(0, digit) + '—' + address.slice(-digit);
}
