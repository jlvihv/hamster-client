package chain

import (
	"encoding/json"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	"hamster-client/module/queue"
	"time"
)

type StarkwareDeployParam struct {
	gorm.Model
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Network        string    `json:"network"`       //rinkbey network or mainnet network
	LeaseTerm      int       `json:"leaseTerm"`     //
	ApplicationID  uint      `json:"applicationId"` //application id
	EthereumApiUrl string    `json:"ethereumApiUrl"`
}

type StarkWare struct {
	common     *Common
	DeployType int
}

func (s *StarkWare) DeployJob(appData application.Application) error {
	return s.common.deployJob(appData, s.DeployType)
}

func (s *StarkWare) SaveDeployParam(appInfo application.Application, deployParam DeployParam, db *gorm.DB) error {
	var deployData StarkwareDeployParam
	deployData.LeaseTerm = appInfo.LeaseTerm
	deployData.Network = deployParam.SelectNodeType
	deployData.ApplicationID = appInfo.ID
	deployData.EthereumApiUrl = config.EthereumEndpointMap[appInfo.SelectNodeType]
	if err := db.Create(&deployData).Error; err != nil {
		return err
	}
	return nil
}

func (s *StarkWare) SaveJsonParam(id string, deployParam DeployParam) error {
	addData := deployParam
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
	s.common.helper.KS().Set(string(application.TYPE_StarkWare)+"_"+id, string(jsonData))
	return nil
}

func (s *StarkWare) GetDeployParam(appID int, db *gorm.DB) any {
	var deployData StarkwareDeployParam
	err := db.Table("starkware_deploy_params").
		Where("application_id = ?", appID).
		First(&deployData).
		Error
	if err != nil {
		return nil
	} else {
		return deployData
	}
}

func NewStarkWare(common *Common, deployType int) Chain {
	return &StarkWare{
		common:     common,
		DeployType: deployType,
	}
}

func (s *StarkWare) CreateQueue(appData application.Application) (queue.Queue, error) {
	return s.common.createQueue(appData, s.DeployType)
}
