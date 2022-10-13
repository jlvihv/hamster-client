package chain

import (
	"encoding/json"
	"hamster-client/config"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	"hamster-client/module/queue"
	"time"

	"gorm.io/gorm"
)

type Ethereum struct {
	common     *Common
	DeployType int
}

type EthereumDeployParam struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Network       string    `json:"network"`       // rinkbey network or mainnet network
	LeaseTerm     int       `json:"leaseTerm"`     //
	ApplicationID uint      `json:"applicationId"` // application id
}

func NewEthereum(common *Common, deployType int) Chain {
	return &Ethereum{
		common:     common,
		DeployType: deployType,
	}
}

func (e *Ethereum) DeployJob(appData application.Application) error {
	return e.common.deployJob(appData, e.DeployType)
}

func (e *Ethereum) CreateQueue(appData application.Application) (queue.Queue, error) {
	return e.common.createQueue(appData, e.DeployType)
}

func (e *Ethereum) GetDeployParam(appID int, db *gorm.DB) any {
	var deployData EthereumDeployParam
	err := db.Table("ethereum_deploy_params").
		Where("application_id = ?", appID).
		First(&deployData).
		Error
	if err != nil {
		return nil
	} else {
		return deployData
	}
}

func (e *Ethereum) SaveDeployParam(appInfo application.Application, deployParam DeployParam, db *gorm.DB) error {
	//deployData := paramData.(EthereumDeployParam)
	var deployData EthereumDeployParam
	deployData.ApplicationID = appInfo.ID
	//deployData.Network = paramData.(DeployParam).Network
	deployData.LeaseTerm = deployParam.LeaseTerm
	err := db.Table("ethereum_deploy_params").Create(&deployData).Error
	return err
}

func (e *Ethereum) SaveJsonParam(id string, deployParam DeployParam) error {
	var deployInfo deploy.DeployParameter
	pluginDeployInfo := config.PluginMap[deployParam.SelectNodeType]
	deployInfo.Initialization.AccountMnemonic = deployParam.ThegraphIndexer
	deployInfo.Initialization.LeaseTerm = deployParam.LeaseTerm
	deployInfo.Staking.PledgeAmount = deployParam.StakingAmount
	deployInfo.Deployment.NodeEthereumUrl = pluginDeployInfo.EthNetwork
	deployInfo.Deployment.EthereumNetwork = pluginDeployInfo.EthereumNetworkName
	deployInfo.Staking.NetworkUrl = pluginDeployInfo.EndpointUrl
	deployInfo.Staking.Address = pluginDeployInfo.TheGraphStakingAddress
	jsonData, err := json.Marshal(deployInfo)
	if err != nil {
		return err
	}
	e.common.helper.KS().Set(string(application.TYPE_Thegraph)+"_"+id, string(jsonData))
	return nil
}
