package chain

import (
	"gorm.io/gorm"
	"hamster-client/module/application"
	"hamster-client/module/queue"
)

type Chain interface {
	Deployer
	SaveDeployParam
	GetDeployParam(appID int, db *gorm.DB) any
}

type SaveDeployParam interface {
	SaveDeployParam(appInfo application.Application, deployParam DeployParam, db *gorm.DB) error
	SaveJsonParam(id string, deployParam DeployParam) error
}

type Deployer interface {
	DeployJob(appData application.Application) error
	CreateQueue(appData application.Application) (queue.Queue, error)
}
