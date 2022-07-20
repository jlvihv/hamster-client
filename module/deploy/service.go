package deploy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/graph"
	"hamster-client/module/keystorage"
	"hamster-client/module/p2p"
	"hamster-client/utils"
	"net/url"
	"strconv"
)

type ServiceImpl struct {
	ctx               context.Context
	httpUtil          *utils.HttpUtil
	db                *gorm.DB
	graphService      graph.Service
	keyStorageService keystorage.Service
	accountService    account.Service
	p2pServer         p2p.Service
}

func NewServiceImpl(ctx context.Context, httpUtil *utils.HttpUtil, db *gorm.DB, graphService graph.Service, keyStorageService *keystorage.Service, accountService account.Service, p2pServer p2p.Service) ServiceImpl {
	return ServiceImpl{ctx, httpUtil, db, graphService, *keyStorageService, accountService, p2pServer}
}

func (s *ServiceImpl) DeployTheGraph(id int, jsonData string) (bool, error) {
	//Judge whether the account has peerId
	info, err := s.accountService.GetAccount()
	if err != nil {
		return false, err
	}
	if info.PeerId == "" {
		//Modify the status of the application to wait for resources
		result := s.db.Model(application.Application{}).Where("id = ?", id).Update("status", config.WAIT_RESOURCE).Error
		if result != nil {
			return false, result
		}
		return true, nil
	}
	//Determine whether to initialize configuration
	_, resultErr := s.p2pServer.GetSetting()
	if resultErr != nil {
		res := s.p2pServer.InitSetting()
		if res != nil {
			return false, err
		}
	}
	fmt.Println("p2p start")
	fmt.Println(info.PeerId)
	proErr := s.p2pServer.ProLink(info.PeerId)
	if proErr != nil {
		return false, proErr
	}
	fmt.Println("p2p end")
	var param DeployParameter
	jsonParam := s.keyStorageService.Get("graph_" + strconv.Itoa(id))
	if err := json.Unmarshal([]byte(jsonParam), &param); err != nil {
		return false, err
	}
	var sendData DeployParams
	sendData.Mnemonic = param.Initialization.AccountMnemonic
	sendData.Id = id
	sendData.EthereumUrl = param.Deployment.EthereumUrl
	sendData.IndexerAddress = param.Deployment.IndexerAddress
	sendData.NodeEthereumUrl = param.Deployment.NodeEthereumUrl
	sendData.EthereumNetwork = param.Deployment.EthereumNetwork
	runtime.LogInfo(s.ctx, "start Deploy the graph")
	res, err := s.httpUtil.NewRequest().SetBody(sendData).Post(config.Httpprovider)
	if err != nil {
		runtime.LogError(s.ctx, "DeployTheGraph http error:"+err.Error())
		return false, err
	}
	if !res.IsSuccess() {
		runtime.LogError(s.ctx, "DeployTheGraph Response error: "+res.Status())
		return false, errors.New(res.Status())
	}
	//Modification status is in deployment
	result := s.db.Model(application.Application{}).Where("id = ?", id).Update("status", config.IN_DEPLOYMENT).Error
	if result != nil {
		return false, result
	}
	go s.queryDeployStatus()
	return true, nil
}

func (s *ServiceImpl) GetDeployInfo(id int) (DeployParameter, error) {
	data := s.keyStorageService.Get("graph_" + strconv.Itoa(id))
	var param DeployParameter
	if err := json.Unmarshal([]byte(data), &param); err != nil {
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

func (g *ServiceImpl) QueryGraphStatus(serviceName ...string) (int, error) {
	var status int
	res, err := g.httpUtil.NewRequest().
		SetQueryParamsFromValues(url.Values{"serviceName": serviceName}).
		SetResult(&status).
		Get(config.HttpGraphStatus)
	if err != nil {
		runtime.LogError(g.ctx, "DeployTheGraph http error:"+err.Error())
		return 0, err
	}
	if !res.IsSuccess() {
		runtime.LogError(g.ctx, "DeployTheGraph Response error: "+res.Status())
		return 0, errors.New(fmt.Sprintf("Query status request failed. The request status is:%s", res.Status()))
	}
	return status, nil
}

// query deploy graph status
func (s *ServiceImpl) queryDeployStatus() {
	containerIds := []string{"graph-node", "postgres", "index-service", "index-agent", "index-cli"}
	for {
		res, _ := s.QueryGraphStatus(containerIds...)
		if res == 1 {
			result := s.db.Model(application.Application{}).Where("status = ?", config.IN_DEPLOYMENT).Update("status", config.DEPLOYED).Error
			if result == nil {
				return
			}
		} else {
			s.db.Model(application.Application{}).Where("status = ?", config.IN_DEPLOYMENT).Update("status", config.DEPLOY_FAILED)
			return
		}
	}
}
