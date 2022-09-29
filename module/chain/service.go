package chain

import "hamster-client/module/queue"

type QueueInfo struct {
	Info []queue.StatusInfo `json:"info"`
}

type Service interface {
	getP2pForwardPort(appID int) (int, error)
	StartQueue(appID int) error
	GetQueueInfo(appID int) (QueueInfo, error)
}
