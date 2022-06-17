package app

import (
	"context"
	"fmt"
	"hamster-client/module/deploy"
)

type Deploy struct {
	ctx           context.Context
	deployService deploy.Service
}

func NewDeployApp(service deploy.Service) Deploy {
	return Deploy{
		deployService: service,
	}
}

func (d *Deploy) WailsInit(ctx context.Context) error {
	d.ctx = ctx
	return nil
}

// DeployTheGraph deploy the graph
func (d *Deploy) DeployTheGraph(nodeEthereumUrl string, ethereumUrl string, ethereumNetwork string, indexerAddress string, mnemonic string) error {
	fmt.Println(9999999999999)
	var data deploy.DeployParams
	data.Mnemonic = mnemonic
	data.IndexerAddress = indexerAddress
	data.NodeEthereumUrl = nodeEthereumUrl
	data.EthereumUrl = ethereumUrl
	data.EthereumNetwork = ethereumNetwork
	fmt.Println(data.Mnemonic)
	return d.deployService.DeployTheGraph(data)
}
