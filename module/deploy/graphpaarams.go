package deploy

type DeployParams struct {
	Id              int    `json:"id"`              //application id
	NodeEthereumUrl string `json:"nodeEthereumUrl"` //graph-node eth-url
	EthereumUrl     string `json:"ethereumUrl"`     //indexer-service、indexer-agent eth-url
	EthereumNetwork string `json:"ethereumNetwork"` //eth network
	IndexerAddress  string `json:"indexerAddress"`  //indexer address
	Mnemonic        string `json:"mnemonic"`        // mnemonic
}

type DeployParameter struct {
	Initialization Initialization `json:"initialization"`
	Staking        Staking        `json:"staking"`
	Deployment     Deployment     `json:"deployment"`
}

type ParameterInfo struct {
	Initialization Initialization `json:"initialization"`
	Staking        Staking        `json:"staking"`
	Deployment     Deployment     `json:"deployment"`
}

type Initialization struct {
	LeaseTerm       int    `json:"leaseTerm"`
	PublicKey       string `json:"publicKey"`
	AccountMnemonic string `json:"accountMnemonic"`
}

type Staking struct {
	NetworkUrl   string `json:"networkUrl"`
	Address      string `json:"address"`
	AgentAddress string `json:"agentAddress"`
	PledgeAmount int    `json:"pledgeAmount"`
}

type Deployment struct {
	NodeEthereumUrl string `json:"nodeEthereumUrl"` //graph-node eth-url
	EthereumUrl     string `json:"ethereumUrl"`     //indexer-service、indexer-agent eth-url
	EthereumNetwork string `json:"ethereumNetwork"` //eth network
	IndexerAddress  string `json:"indexerAddress"`  //indexer address
}

type Service interface {
	DeployTheGraph(id int, data string) (bool, error)
	GetDeployInfo(id int) (DeployParameter, error)
	SaveDeployInfo(id int, json string) (bool, error)
	QueryGraphStatus(serviceName ...string) (int, error)
}
