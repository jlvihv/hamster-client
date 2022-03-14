package resource

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/utils"
)

type HttpService struct {
	ctx      context.Context
	db       *gorm.DB
	httpUtil *utils.HttpUtil
}

func NewServiceImpl(ctx context.Context, db *gorm.DB, httpUtil *utils.HttpUtil) HttpService {
	return HttpService{ctx, db, httpUtil}
}

// GetResourceList query my resource list
func (r *HttpService) GetResourceList(publicKey string) ([]Resource, error) {
	runtime.LogInfo(r.ctx, "start GetResourceList")
	var resources []Resource
	// get my resource list via http request
	res, err := r.httpUtil.NewRequest().
		SetQueryParam("user", publicKey).
		SetResult(&resources).
		Get(config.HttpGetResource)
	if err != nil {
		runtime.LogError(r.ctx, "GetResourceList http error:"+err.Error())
	}
	if !res.IsSuccess() {
		runtime.LogError(r.ctx, "GetResourceList Response error: "+res.Status())
	}
	return resources, err
}
