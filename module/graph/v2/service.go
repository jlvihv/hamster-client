package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	"hamster-client/module/keystorage"
	"hamster-client/module/p2p"
	queue2 "hamster-client/module/queue"
	"hamster-client/module/wallet"
	"strconv"
)

type ServiceImpl struct {
	ctx                context.Context
	db                 *gorm.DB
	keyStorageService  keystorage.Service
	accountService     account.Service
	applicationService application.Service
	p2pServer          p2p.Service
	deployService      deploy.Service
	walletService      wallet.Service
}

func NewServiceImpl(ctx context.Context, db *gorm.DB, keyStorageService keystorage.Service, accountService account.Service, applicationService application.Service, p2pServer p2p.Service, deployService deploy.Service, walletService wallet.Service) ServiceImpl {
	return ServiceImpl{ctx, db, keyStorageService, accountService, applicationService, p2pServer, deployService, walletService}
}

func (g *ServiceImpl) SaveGraphDeployParameterAndApply(addData AddParam) (AddApplicationVo, error) {
	var applyData application.Application
	var deployData GraphDeployParameter
	var applicationVo AddApplicationVo
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("name=?", addData.Name).First(&applyData).Error; err != gorm.ErrRecordNotFound {
			return errors.New(fmt.Sprintf("application:%s already exists", addData.Name))
		}
		applyData.Name = addData.Name
		applyData.SelectNodeType = addData.SelectNodeType
		applyData.LeaseTerm = addData.LeaseTerm
		if err := tx.Create(&applyData).Error; err != nil {
			return err
		}
		deployData.Application = applyData
		deployData.LeaseTerm = addData.LeaseTerm
		deployData.ThegraphIndexer = addData.ThegraphIndexer
		deployData.StakingAmount = addData.StakingAmount
		deployData.ApplicationID = applyData.ID
		if err := tx.Create(&deployData).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		applicationVo.Result = false
		return applicationVo, err
	}
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
		applicationVo.Result = false
		return applicationVo, err
	}
	g.keyStorageService.Set("graph_"+strconv.Itoa(int(applyData.ID)), string(jsonData))
	applicationVo.Result = true
	applicationVo.ID = applyData.ID
	go g.deployGraphJob(int(applyData.ID), pluginDeployInfo.EndpointUrl)
	return applicationVo, nil
}

func (g *ServiceImpl) DeleteGraphDeployParameterAndApply(id int) (bool, error) {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Where("id = ?", id).Delete(&application.Application{}).Error; err != nil {
			return err
		}
		if err := tx.Debug().Where("applicationId = ?", id).Delete(GraphDeployParameter{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	//delete key storage
	//stop docker
	return true, nil
}

func (g *ServiceImpl) deployGraphJob(applicationId int, networkUrl string) {
	stakingJob := NewGraphStakingJob(g.keyStorageService, applicationId, networkUrl)
	var accountInfo account.Account
	accountInfo, err := g.accountService.GetAccount()
	if err != nil {
		accountInfo.WsUrl = config.DefaultPolkadotNode
	}
	substrateApi, _ := gsrpc.NewSubstrateAPI(accountInfo.WsUrl)
	waitResourceJob, _ := NewWaitResourceJob(substrateApi, g.accountService, g.applicationService, g.p2pServer, applicationId, g.walletService)

	pullJob := NewPullImageJob(g.applicationService, applicationId, g.walletService)

	deployJob := NewServiceDeployJob(g.keyStorageService, g.deployService, applicationId, g.walletService)

	queue, err := queue2.NewQueue(applicationId, &stakingJob, waitResourceJob, &pullJob, &deployJob)
	if err != nil {
		fmt.Println("new queue failed,err is: ", err)
	}
	channel := make(chan struct{})
	go queue.Start(channel)
}

func (g *ServiceImpl) DeployGraphJob(applicationId int) error {
	var data application.Application
	result := g.db.Where("id = ? ", applicationId).First(&data)
	if result.Error != nil {
		fmt.Println("query application failed,error is: ", result.Error)
		return result.Error
	}
	pluginDeployInfo := config.PluginMap[data.SelectNodeType]
	stakingJob := NewGraphStakingJob(g.keyStorageService, applicationId, pluginDeployInfo.EndpointUrl)
	var accountInfo account.Account
	accountInfo, err := g.accountService.GetAccount()
	if err != nil {
		accountInfo.WsUrl = config.DefaultPolkadotNode
	}
	substrateApi, _ := gsrpc.NewSubstrateAPI(accountInfo.WsUrl)
	waitResourceJob, _ := NewWaitResourceJob(substrateApi, g.accountService, g.applicationService, g.p2pServer, applicationId, g.walletService)

	pullJob := NewPullImageJob(g.applicationService, applicationId, g.walletService)

	deployJob := NewServiceDeployJob(g.keyStorageService, g.deployService, applicationId, g.walletService)

	queue, err := queue2.NewQueue(applicationId, &stakingJob, waitResourceJob, &pullJob, &deployJob)
	if err != nil {
		return err
		fmt.Println("new queue failed,err is: ", err)
	}
	channel := make(chan struct{})
	go queue.Start(channel)
	return nil
}
