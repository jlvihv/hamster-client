package app

import (
	"context"
	"fmt"
	"hamster-client/module/account"
	"hamster-client/module/deploy"
	"hamster-client/module/p2p"
)

type Deploy struct {
	ctx            context.Context
	deployService  deploy.Service
	p2pServer      p2p.Service
	accountService account.Service
}

func NewDeployApp(service deploy.Service, accountService account.Service, p2p p2p.Service) Deploy {
	return Deploy{
		deployService:  service,
		accountService: accountService,
		p2pServer:      p2p,
	}
}

func (d *Deploy) WailsInit(ctx context.Context) error {
	d.ctx = ctx
	return nil
}

// DeployTheGraph deploy the graph
func (d *Deploy) DeployTheGraph(nodeEthereumUrl string, ethereumUrl string, ethereumNetwork string, indexerAddress string, mnemonic string) error {
	var data deploy.DeployParams
	data.Mnemonic = mnemonic
	data.IndexerAddress = indexerAddress
	data.NodeEthereumUrl = nodeEthereumUrl
	data.EthereumUrl = ethereumUrl
	data.EthereumNetwork = ethereumNetwork
	fmt.Println("p2p start")
	info, err := d.accountService.GetAccount()
	if err != nil {
		return nil
	}
	fmt.Println(info.PeerId)
	proErr := d.p2pServer.ProLink(info.PeerId)
	if proErr != nil {
		return proErr
	}
	fmt.Println("p2p end")
	return d.deployService.DeployTheGraph(data)
}
