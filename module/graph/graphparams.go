package graph

import (
	"time"
)

type GraphParameterVo struct {
	NodeEthereumUrl string    `json:"nodeEthereumUrl"` //graph-node eth-url
	EthereumUrl     string    `json:"ethereumUrl"`     //indexer-service、indexer-agent eth-url
	EthereumNetwork string    `json:"ethereumNetwork"` //eth network
	IndexerAddress  string    `json:"indexerAddress"`  //indexer address
	Mnemonic        string    `json:"mnemonic"`        // mnemonic
	ApplicationId   uint      `json:"applicationId"`   //application id
	Name            string    `json:"name"`            //apply name
	Plugin          string    `json:"plugin"`          //apply plugin
	Status          int       `json:"status"`          //apply status 0: not deploy 1:deployed 2:ALL
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
