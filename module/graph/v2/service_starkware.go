package v2

import (
	"encoding/json"
	"fmt"
	"hamster-client/config"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/chainhelper"
	"hamster-client/module/chainjob"
	"hamster-client/module/common"
	"hamster-client/module/deploy"

	"gorm.io/gorm"

	queue2 "hamster-client/module/queue"
)

type StarkWareService struct {
	ServiceImpl
	DeployType int
}

func (s *StarkWareService) saveDeployParam(
	appData application.Application,
	paramData interface{},
	tx *gorm.DB,
) error {
	var deployData common.StarkwareDeployParam
	deployData.LeaseTerm = appData.LeaseTerm
	addData := paramData.(AddParam)
	deployData.Network = addData.SelectNodeType
	deployData.ApplicationID = appData.ID
	deployData.EthereumApiUrl = config.EthereumEndpointMap[appData.SelectNodeType]
	if err := tx.Create(&deployData).Error; err != nil {
		return err
	}
	return nil
}

func (g *StarkWareService) deployJob(addData application.Application) {
	applicationId := int(addData.ID)
	var accountInfo account.Account
	accountInfo, err := g.accountService.GetAccount()
	if err != nil {
		accountInfo.WsUrl = config.DefaultPolkadotNode
	}
	substrateApi, err := g.accountService.GetSubstrateApi()
	if err != nil {
		return
	}
	waitResourceJob, _ := NewWaitResourceJob(
		substrateApi,
		g.accountService,
		g.applicationService,
		g.p2pService,
		applicationId,
		g.walletService,
		g.DeployType,
	)

	helper := chainhelper.NewHelper(
		g.db,
		g.applicationService,
		g.p2pService,
		g.accountService,
		g.walletService,
	)

	pullJob := chainjob.NewPullImageJob(applicationId, helper)

	deployJob := chainjob.NewStartJob(applicationId, helper)

	queue, err := queue2.NewQueue(applicationId, g.db, waitResourceJob, pullJob, deployJob)
	if err != nil {
		fmt.Println("new queue failed,err is: ", err)
	}
	channel := make(chan struct{})
	go queue.Start(channel)
	<-channel
}

func (g *StarkWareService) saveJsonParam(id string, paramData interface{}) error {
	addData := paramData.(AddParam)
	var deploymentData deploy.DeployParameter
	deploymentData.Initialization.AccountMnemonic = addData.ThegraphIndexer
	deploymentData.Initialization.LeaseTerm = addData.LeaseTerm
	deploymentData.Staking.PledgeAmount = addData.StakingAmount
	deploymentData.Deployment.EthereumUrl = config.EthereumEndpointMap[addData.SelectNodeType]
	deploymentData.Deployment.EthereumNetwork = addData.SelectNodeType
	jsonData, err := json.Marshal(deploymentData)
	if err != nil {
		return err
	}
	g.keyStorageService.Set(string(application.TYPE_Thegraph)+"_"+id, string(jsonData))
	return nil
}

func (g *StarkWareService) getDeployParamByAppId(appId int) interface{} {
	var deployData common.StarkwareDeployParam
	err := g.db.Table("starkware_deploy_params").
		Where("application_id = ?", appId).
		First(&deployData).
		Error
	if err != nil {
		return nil
	} else {
		return deployData
	}
}
