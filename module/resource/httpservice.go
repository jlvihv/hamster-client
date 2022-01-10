package resource

import (
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
	"gorm.io/gorm"
	"link-server/config"
	"link-server/utils"
)

type HttpService struct {
	log      *wails.CustomLogger
	db       *gorm.DB
	httpUtil *utils.HttpUtil
}

func NewServiceImpl(db *gorm.DB, httpUtil *utils.HttpUtil) HttpService {
	log := logger.NewCustomLogger("Module_P2P")
	return HttpService{log, db, httpUtil}
}

// GetResourceList query my resource list
func (r *HttpService) GetResourceList(publicKey string) ([]Resource, error) {
	r.log.Info("start GetResourceList")
	var resources []Resource
	// get my resource list via http request
	res, err := r.httpUtil.NewRequest().
		SetQueryParam("user", publicKey).
		SetResult(&resources).
		Get(config.HttpGetResource)
	if err != nil {
		r.log.Errorf("GetResourceList http error: %s\n", err)
	}
	if !res.IsSuccess() {
		r.log.Errorf("GetResourceList Response error: %s\n", res.Status())
	}
	return resources, err
}
