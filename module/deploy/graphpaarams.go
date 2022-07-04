package deploy

type DeployParams struct {
	Id              int    `json:"id"`              //application id
	NodeEthereumUrl string `json:"nodeEthereumUrl"` //graph-node eth-url
	EthereumUrl     string `json:"ethereumUrl"`     //indexer-service、indexer-agent eth-url
	EthereumNetwork string `json:"ethereumNetwork"` //eth network
	IndexerAddress  string `json:"indexerAddress"`  //indexer address
	Mnemonic        string `json:"mnemonic"`        // mnemonic
}

type Service interface {
	DeployTheGraph(data DeployParams) (bool, error)
}
