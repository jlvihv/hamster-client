package graph

import (
	"gorm.io/gorm"
	"hamster-client/module/application"
)

type GraphParameter struct {
	gorm.Model
	NodeEthereumUrl string                  `json:"nodeEthereumUrl"` //graph-node eth-url
	EthereumUrl     string                  `json:"ethereumUrl"`     //indexer-service、indexer-agent eth-url
	EthereumNetwork string                  `json:"ethereumNetwork"` //eth network
	IndexerAddress  string                  `json:"indexerAddress"`  //indexer address
	Mnemonic        string                  `json:"mnemonic"`        // mnemonic
	Application     application.Application `json:"application"`     //application entity
	ApplicationId   uint                    `json:"applicationId"`   //application id
}

type Service interface {
	SaveGraphParameter(data GraphParameter) error
	QueryParamByApplyId(applicationId int) (GraphParameter, error)
	DeleteGraphAndParams(applicationId int) error
}
