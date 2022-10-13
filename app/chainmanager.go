package app

import "hamster-client/module/chainmanager"

type ChainManager struct {
	manager chainmanager.Manager
}

func NewChainManagerApp(manager chainmanager.Manager) ChainManager {
	return ChainManager{
		manager: manager,
	}
}

func (c *ChainManager) Start(appID int) error {
	//return c.manager.CreateAndStart(appID)
	//panic("implement me")
	return nil
}

func (c *ChainManager) GetStatusInfo(appID int) (chainmanager.QueueInfo, error) {
	return c.manager.GetQueueInfo(appID)
}

func (c *ChainManager) RetryStart(appID int, runNow bool) error {
	return c.manager.RetryStartQueue(appID, runNow)
}
