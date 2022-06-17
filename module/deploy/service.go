package deploy

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hamster-client/config"
	"hamster-client/utils"
)

type ServiceImpl struct {
	ctx      context.Context
	httpUtil *utils.HttpUtil
}

func NewServiceImpl(ctx context.Context, httpUtil *utils.HttpUtil) ServiceImpl {
	return ServiceImpl{ctx, httpUtil}
}

func (s *ServiceImpl) DeployTheGraph(data DeployParams) error {
	runtime.LogInfo(s.ctx, "start Deploy the graph")
	res, err := s.httpUtil.NewRequest().SetBody(data).Post(config.Httpprovider)
	if err != nil {
		runtime.LogError(s.ctx, "DeployTheGraph http error:"+err.Error())
	}
	if !res.IsSuccess() {
		runtime.LogError(s.ctx, "DeployTheGraph Response error: "+res.Status())
	}
	return err
}
