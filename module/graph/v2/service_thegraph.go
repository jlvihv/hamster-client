package v2

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	queue2 "hamster-client/module/queue"
)

type ThegraphDeploySaveServiceImpl struct {
	ServiceImpl
}

func (s *ThegraphDeploySaveServiceImpl) saveDeployParam(appData application.Application, paramData interface{}, tx *gorm.DB) error {
	var deployData GraphDeployParameter
	deployData.Application = appData
	deployData.LeaseTerm = appData.LeaseTerm
	addData := paramData.(AddParam)
	deployData.ThegraphIndexer = addData.ThegraphIndexer
	deployData.StakingAmount = addData.StakingAmount
	deployData.ApplicationID = appData.ID
	if err := tx.Create(&deployData).Error; err != nil {
		return err
	}
	return nil
}
func (g *ThegraphDeploySaveServiceImpl) deployJob(addData application.Application) {
	applicationId := int(addData.ID)
	pluginDeployInfo := config.PluginMap[addData.SelectNodeType]
	stakingJob := NewGraphStakingJob(g.keyStorageService, applicationId, pluginDeployInfo.EndpointUrl, pluginDeployInfo.ChainId)
	var accountInfo account.Account
	accountInfo, err := g.accountService.GetAccount()
	if err != nil {
		accountInfo.WsUrl = config.DefaultPolkadotNode
	}
	substrateApi, err := g.accountService.GetSubstrateApi()
	if err != nil {
		return
	}
	const deployType = 0
	waitResourceJob, _ := NewWaitResourceJob(substrateApi, g.accountService, g.applicationService, g.p2pService, applicationId, g.walletService, deployType)

	pullJob := NewPullImageJob(g.applicationService, applicationId, g.p2pService, g.accountService, g.walletService)

	deployJob := NewServiceDeployJob(g.keyStorageService, g.deployService, applicationId, g.p2pService, g.accountService, g.applicationService, g.walletService)

	queue, err := queue2.NewQueue(applicationId, g.db, &stakingJob, waitResourceJob, &pullJob, &deployJob)
	if err != nil {
		fmt.Println("new queue failed,err is: ", err)
	}
	channel := make(chan struct{})
	go queue.Start(channel)
	<-channel
}

func (g *ThegraphDeploySaveServiceImpl) saveJsonParam(id string, paramData interface{}) error {
	addData := paramData.(AddParam)
	var deploymentData deploy.DeployParameter
	pluginDeployInfo := config.PluginMap[addData.SelectNodeType]
	deploymentData.Initialization.AccountMnemonic = addData.ThegraphIndexer
	deploymentData.Initialization.LeaseTerm = addData.LeaseTerm
	deploymentData.Staking.PledgeAmount = addData.StakingAmount
	deploymentData.Deployment.NodeEthereumUrl = pluginDeployInfo.EthNetwork
	deploymentData.Deployment.EthereumUrl = pluginDeployInfo.EndpointUrl
	deploymentData.Deployment.EthereumNetwork = pluginDeployInfo.EthereumNetworkName
	deploymentData.Staking.NetworkUrl = pluginDeployInfo.EndpointUrl
	deploymentData.Staking.Address = pluginDeployInfo.TheGraphStakingAddress
	jsonData, err := json.Marshal(deploymentData)
	if err != nil {
		return err
	}
	g.keyStorageService.Set(string(application.TYPE_THEGRAPH)+"_"+id, string(jsonData))
	return nil
}
