package app

import "hamster-client/module/chain"

type Chain struct {
	chainService chain.Service
}

func NewChainApp(service chain.Service) Chain {
	return Chain{
		chainService: service,
	}
}

func (c *Chain) StartQueue(appID int) error {
	return c.chainService.StartQueue(appID)
}

func (c *Chain) GetQueueInfo(appID int) (chain.QueueInfo, error) {
	return c.chainService.GetQueueInfo(appID)
}
