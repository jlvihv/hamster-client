package chain

type Service interface {
	getP2pForwardPort(appID int) (p2pForwardPort int, err error)
	StartQueue(appID int) error
	GetQueueInfo(appID int) (QueueInfo, error)
}
