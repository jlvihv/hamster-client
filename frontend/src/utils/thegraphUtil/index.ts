export const pluginConfigs = [
  { plugin: 'The Graph(Görli)' },
  { plugin: 'The Graph(Ethereum)', disabled: true },
  { plugin: 'Polygon(Ethereum)', disabled: true },
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
  return address.slice(0, digit) + '-' + address.slice(-digit);
}
