package chainmanager

import (
	"hamster-client/module/chain"
	"hamster-client/module/queue"
)

type QueueInfo struct {
	Info []queue.StatusInfo `json:"info"`
}

type Manager interface {
	CreateAndStart(appInfo chain.DeployParam) (chain.DeployResult, error)
	RetryStartQueue(appID int, runNow bool) error
	GetQueueInfo(appID int) (QueueInfo, error)
}
