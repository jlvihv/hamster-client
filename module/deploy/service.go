package deploy

import (
	"context"
	"encoding/json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/application"
	"hamster-client/module/graph"
	"hamster-client/module/keystorage"
	"hamster-client/utils"
	"strconv"
)

type ServiceImpl struct {
	ctx               context.Context
	httpUtil          *utils.HttpUtil
	db                *gorm.DB
	graphService      graph.Service
	keyStorageService keystorage.Service
}

func NewServiceImpl(ctx context.Context, httpUtil *utils.HttpUtil, db *gorm.DB, graphService graph.Service, keyStorageService *keystorage.Service) ServiceImpl {
	return ServiceImpl{ctx, httpUtil, db, graphService, *keyStorageService}
}

func (s *ServiceImpl) DeployTheGraph(data DeployParams) (bool, error) {
	runtime.LogInfo(s.ctx, "start Deploy the graph")
	res, err := s.httpUtil.NewRequest().SetBody(data).Post(config.Httpprovider)
	if err != nil {
		runtime.LogError(s.ctx, "DeployTheGraph http error:"+err.Error())
		return false, err
	}
	if !res.IsSuccess() {
		runtime.LogError(s.ctx, "DeployTheGraph Response error: "+res.Status())
	}
	// save graph config params
	var applyData application.Application
	result := s.db.Where("id = ? ", data.Id).First(&applyData)
	if result.Error != nil {
		return false, result.Error
	}
	var graphData graph.GraphParameter
	graphData.Application = applyData
	graphData.EthereumNetwork = data.EthereumNetwork
	graphData.NodeEthereumUrl = data.NodeEthereumUrl
	graphData.ApplicationId = uint(data.Id)
	graphData.EthereumUrl = data.EthereumUrl
	graphData.Mnemonic = data.Mnemonic
	graphData.IndexerAddress = data.IndexerAddress
	return s.graphService.SaveGraphParameter(graphData)
}

func (s *ServiceImpl) GetDeployInfo(id int) (DeployParameter, error) {
	data := s.keyStorageService.Get("graph_" + strconv.Itoa(id))
	var param DeployParameter
	if err := json.Unmarshal([]byte(data), &param); err == nil {
		return param, err
	}
	return param, nil
}

func (s *ServiceImpl) SaveDeployInfo(id int, json string) (bool, error) {
	s.keyStorageService.Set("graph_"+strconv.Itoa(id), json)
	err := s.keyStorageService.Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
