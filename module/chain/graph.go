package chain

import (
	"hamster-client/module/application"
	"hamster-client/module/queue"
	"time"
)

type GraphDeployParameter struct {
	ID              uint                    `json:"id"`
	CreatedAt       time.Time               `json:"createdAt"`
	UpdatedAt       time.Time               `json:"updatedAt"`
	Network         string                  `json:"network"`         // rinkbey network or mainnet network
	LeaseTerm       int                     `json:"leaseTerm"`       // indexer-service、indexer-agent eth-url
	ThegraphIndexer string                  `json:"thegraphIndexer"` // mnemonic
	StakingAmount   int                     `json:"stakingAmount"`
	Application     application.Application `json:"application"`   // application entity
	ApplicationID   uint                    `json:"applicationId"` // application id
}

type QueueInfo struct {
	Info []queue.StatusInfo `json:"info"`
}
