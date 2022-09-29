package app

import (
	"context"
	"hamster-client/module/account"
	"hamster-client/module/deploy"
	"hamster-client/module/p2p"
)

type Deploy struct {
	ctx            context.Context
	deployService  deploy.Service
	p2pService     p2p.Service
	accountService account.Service
}

func NewDeployApp(service deploy.Service, accountService account.Service, p2p p2p.Service) Deploy {
	return Deploy{
		deployService:  service,
		accountService: accountService,
		p2pService:     p2p,
	}
}

func (d *Deploy) WailsInit(ctx context.Context) error {
	d.ctx = ctx
	return nil
}

// DeployTheGraph deploy the graph
func (d *Deploy) DeployTheGraph(id int, data string) (bool, error) {
	return d.deployService.DeployTheGraph(id, data)
}

func (d *Deploy) GetDeployInfo(id int) (deploy.DeployParameter, error) {
	return d.deployService.GetDeployInfo(id)
}

func (d *Deploy) SaveDeployInfo(id int, json string) (bool, error) {
	return d.deployService.SaveDeployInfo(id, json)
}

func (d *Deploy) QueryGraphStatus(id int, serviceName ...string) (int, error) {
	return d.deployService.QueryGraphStatus(id, serviceName...)
}
