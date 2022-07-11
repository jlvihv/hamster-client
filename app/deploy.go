package app

import (
	"context"
	"encoding/json"
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
func (d *Deploy) DeployTheGraph(params string) (bool, error) {
	var param deploy.DeployParameter
	if err := json.Unmarshal([]byte(params), &param); err == nil {
		return false, err
	}
	var data deploy.DeployParams
	data.Mnemonic = param.Data.Initialization.AccountMnemonic
	data.IndexerAddress = param.Data.Deployment.IndexerAddress
	data.NodeEthereumUrl = param.Data.Deployment.NodeEthereumUrl
	data.EthereumUrl = param.Data.Deployment.EthereumUrl
	data.EthereumNetwork = param.Data.Deployment.EthereumNetwork
	data.Id = param.Id
	fmt.Println("p2p start")
	info, err := d.accountService.GetAccount()
	if err != nil {
		return false, err
	}
	fmt.Println(info.PeerId)
	proErr := d.p2pServer.ProLink(info.PeerId)
	if proErr != nil {
		return false, proErr
	}
	fmt.Println("p2p end")
	return d.deployService.DeployTheGraph(data, params)
}

func (d *Deploy) GetDeployInfo(id int) (deploy.DeployParameter, error) {
	return d.deployService.GetDeployInfo(id)
}
