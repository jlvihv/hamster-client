package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	"hamster-client/module/keystorage"
	"strconv"
)

type ServiceImpl struct {
	ctx               context.Context
	db                *gorm.DB
	keyStorageService keystorage.Service
}

func NewServiceImpl(ctx context.Context, db *gorm.DB, keyStorageService keystorage.Service) ServiceImpl {
	return ServiceImpl{ctx, db, keyStorageService}
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
	deploymentData.Initialization.AccountMnemonic = addData.ThegraphIndexer
	deploymentData.Initialization.LeaseTerm = addData.LeaseTerm
	deploymentData.Staking.PledgeAmount = addData.StakingAmount
	deploymentData.Deployment.NodeEthereumUrl = config.EthMainNetwork
	deploymentData.Deployment.EthereumUrl = config.EndpointUrl
	deploymentData.Deployment.EthereumNetwork = config.EthereumRinkebyNetworkName
	deploymentData.Staking.NetworkUrl = config.EndpointUrl
	deploymentData.Staking.Address = config.TheGraphStakingAddress
	jsonData, err := json.Marshal(deploymentData)
	if err != nil {
		applicationVo.Result = false
		return applicationVo, err
	}
	g.keyStorageService.Set("graph_"+strconv.Itoa(int(applyData.ID)), string(jsonData))
	applicationVo.Result = true
	applicationVo.ID = applyData.ID
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
