package deploy

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/application"
	"hamster-client/module/graph"
	"hamster-client/utils"
)

type ServiceImpl struct {
	ctx          context.Context
	httpUtil     *utils.HttpUtil
	db           *gorm.DB
	graphService graph.Service
}

func NewServiceImpl(ctx context.Context, httpUtil *utils.HttpUtil, db *gorm.DB, graphService graph.Service) ServiceImpl {
	return ServiceImpl{ctx, httpUtil, db, graphService}
}

func (s *ServiceImpl) DeployTheGraph(data DeployParams) error {
	runtime.LogInfo(s.ctx, "start Deploy the graph")
	res, err := s.httpUtil.NewRequest().SetBody(data).Post(config.Httpprovider)
	if err != nil {
		runtime.LogError(s.ctx, "DeployTheGraph http error:"+err.Error())
		return err
	}
	if !res.IsSuccess() {
		runtime.LogError(s.ctx, "DeployTheGraph Response error: "+res.Status())
	}
	// save graph config params
	var applyData application.Application
	result := s.db.Where("id = ? ", data.Id).First(&applyData)
	if result.Error != nil {
		return result.Error
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
