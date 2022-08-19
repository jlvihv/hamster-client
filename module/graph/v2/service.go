package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	"hamster-client/module/keystorage"
	"hamster-client/module/p2p"
	queue2 "hamster-client/module/queue"
	"hamster-client/module/wallet"
	"hamster-client/utils"
	"strconv"
)

type ServiceImpl struct {
	ctx                context.Context
	db                 *gorm.DB
	keyStorageService  keystorage.Service
	accountService     account.Service
	applicationService application.Service
	p2pServer          p2p.Service
	deployService      deploy.Service
	walletService      wallet.Service
	queueService       queue2.Service
}

func NewServiceImpl(ctx context.Context, db *gorm.DB, keyStorageService keystorage.Service, accountService account.Service, applicationService application.Service, p2pServer p2p.Service, deployService deploy.Service, walletService wallet.Service, queueService queue2.Service) ServiceImpl {
	return ServiceImpl{ctx, db, keyStorageService, accountService, applicationService, p2pServer, deployService, walletService, queueService}
}

func (g *ServiceImpl) SaveGraphDeployParameterAndApply(addData AddParam) (AddApplicationVo, error) {
	var applyData application.Application
	var deployData GraphDeployParameter
	var applicationVo AddApplicationVo
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("name=?", addData.Name).First(&applyData).Error; err != gorm.ErrRecordNotFound {
			return errors.New(fmt.Sprintf("application:%s already exists", addData.Name))
		}
		applyData.Name = addData.Name
		applyData.SelectNodeType = addData.SelectNodeType
		applyData.LeaseTerm = addData.LeaseTerm
		if err := tx.Create(&applyData).Error; err != nil {
			return err
		}
		deployData.Application = applyData
		deployData.LeaseTerm = addData.LeaseTerm
		deployData.ThegraphIndexer = addData.ThegraphIndexer
		deployData.StakingAmount = addData.StakingAmount
		deployData.ApplicationID = applyData.ID
		if err := tx.Create(&deployData).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		applicationVo.Result = false
		return applicationVo, err
	}
	var deploymentData deploy.DeployParameter
	pluginDeployInfo := config.PluginMap[addData.SelectNodeType]
	deploymentData.Initialization.AccountMnemonic = addData.ThegraphIndexer
	deploymentData.Initialization.LeaseTerm = addData.LeaseTerm
	deploymentData.Staking.PledgeAmount = addData.StakingAmount
	deploymentData.Deployment.NodeEthereumUrl = pluginDeployInfo.EthNetwork
	deploymentData.Deployment.EthereumUrl = pluginDeployInfo.EndpointUrl
	deploymentData.Deployment.EthereumNetwork = pluginDeployInfo.EthereumNetworkName
	deploymentData.Staking.NetworkUrl = pluginDeployInfo.EndpointUrl
	deploymentData.Staking.Address = pluginDeployInfo.TheGraphStakingAddress
	jsonData, err := json.Marshal(deploymentData)
	if err != nil {
		applicationVo.Result = false
		return applicationVo, err
	}
	g.keyStorageService.Set("graph_"+strconv.Itoa(int(applyData.ID)), string(jsonData))
	applicationVo.Result = true
	applicationVo.ID = applyData.ID
	go g.deployGraphJob(int(applyData.ID), pluginDeployInfo.EndpointUrl)
	return applicationVo, nil
}

func (g *ServiceImpl) DeleteGraphDeployParameterAndApply(id int) (bool, error) {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Where("id = ?", id).Delete(&application.Application{}).Error; err != nil {
			return err
		}
		if err := tx.Debug().Where("application_id = ?", id).Delete(GraphDeployParameter{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	//delete key storage
	//stop docker
	return true, nil
}

func (g *ServiceImpl) deployGraphJob(applicationId int, networkUrl string) {
	stakingJob := NewGraphStakingJob(g.keyStorageService, applicationId, networkUrl)
	var accountInfo account.Account
	accountInfo, err := g.accountService.GetAccount()
	if err != nil {
		accountInfo.WsUrl = config.DefaultPolkadotNode
	}
	substrateApi, _ := gsrpc.NewSubstrateAPI(accountInfo.WsUrl)
	waitResourceJob, _ := NewWaitResourceJob(substrateApi, g.accountService, g.applicationService, g.p2pServer, applicationId, g.walletService)

	pullJob := NewPullImageJob(g.applicationService, applicationId, g.p2pServer, g.accountService, g.walletService)

	deployJob := NewServiceDeployJob(g.keyStorageService, g.deployService, applicationId, g.p2pServer, g.accountService, g.applicationService, g.walletService)

	queue, err := queue2.NewQueue(applicationId, &stakingJob, waitResourceJob, &pullJob, &deployJob)
	if err != nil {
		fmt.Println("new queue failed,err is: ", err)
	}
	channel := make(chan struct{})
	defer func() {
		err = queue.SaveStatus(g.db)
		if err != nil {
			fmt.Println("save status failed,err is: ", err)
		}
	}()
	go queue.Start(channel)
	<-channel
}

func (g *ServiceImpl) RetryDeployGraphJob(applicationId int, runNow bool) error {
	var queue queue2.Queue
	queue, err := queue2.GetQueue(applicationId)
	if err != nil {
		var data application.Application
		result := g.db.Where("id = ? ", applicationId).First(&data)
		if result.Error != nil {
			fmt.Println("query application failed,error is: ", result.Error)
			return result.Error
		}
		pluginDeployInfo := config.PluginMap[data.SelectNodeType]
		stakingJob := NewGraphStakingJob(g.keyStorageService, applicationId, pluginDeployInfo.EndpointUrl)
		var accountInfo account.Account
		accountInfo, err := g.accountService.GetAccount()
		if err != nil {
			accountInfo.WsUrl = config.DefaultPolkadotNode
		}
		substrateApi, _ := gsrpc.NewSubstrateAPI(accountInfo.WsUrl)
		waitResourceJob, _ := NewWaitResourceJob(substrateApi, g.accountService, g.applicationService, g.p2pServer, applicationId, g.walletService)

		pullJob := NewPullImageJob(g.applicationService, applicationId, g.p2pServer, g.accountService, g.walletService)

		deployJob := NewServiceDeployJob(g.keyStorageService, g.deployService, applicationId, g.p2pServer, g.accountService, g.applicationService, g.walletService)

		queue, err = queue2.NewQueue(applicationId, &stakingJob, waitResourceJob, &pullJob, &deployJob)
		if err != nil {
			fmt.Println("new queue failed, err is: ", err)
			return err
		}
	}

	err = queue.LoadStatus(g.db)
	if err != nil {
		fmt.Println("queue LoadStatus error, init queue")
		queue.InitStatus()
	}
	channel := make(chan struct{})
	defer func() {
		err = queue.SaveStatus(g.db)
		if err != nil {
			fmt.Println("save status failed,err is: ", err)
		}
	}()

	if runNow {
		go queue.Start(channel)
		<-channel
	} else {
		statusInfo, err := queue.GetStatus()
		if err != nil {
			log.Errorf("get status failed, error: %v", err)
			return nil
		}
		for _, job := range statusInfo {
			if job.Status == queue2.Running {
				log.Infof("job %s status is running, set to failed", job.Name)
				queue.SetJobStatus(job.Name, queue2.StatusInfo{
					Name:   job.Name,
					Status: queue2.Failed,
					Error:  "Abnormal exit",
				})
			}
		}
	}

	return nil
}

// FakeQueue 创建一个虚假的队列，用来在重新启动后，展示给前端页面
func (g *ServiceImpl) FakeQueue(applicationId int) {

}

func (g *ServiceImpl) GetQueueInfo(applicationId int) (QueueInfo, error) {
	var data application.Application
	result := g.db.Where("id = ? ", applicationId).First(&data)
	if result.Error != nil {
		fmt.Println("query application failed,error is: ", result.Error)
		return QueueInfo{}, result.Error
	}
	pluginDeployInfo := config.PluginMap[data.SelectNodeType]
	stakingJob := NewGraphStakingJob(g.keyStorageService, applicationId, pluginDeployInfo.EndpointUrl)
	var accountInfo account.Account
	accountInfo, err := g.accountService.GetAccount()
	if err != nil {
		accountInfo.WsUrl = config.DefaultPolkadotNode
	}
	substrateApi, _ := gsrpc.NewSubstrateAPI(accountInfo.WsUrl)
	waitResourceJob, _ := NewWaitResourceJob(substrateApi, g.accountService, g.applicationService, g.p2pServer, applicationId, g.walletService)

	pullJob := NewPullImageJob(g.applicationService, applicationId, g.p2pServer, g.accountService, g.walletService)

	deployJob := NewServiceDeployJob(g.keyStorageService, g.deployService, applicationId, g.p2pServer, g.accountService, g.applicationService, g.walletService)

	queue, err := queue2.NewQueue(applicationId, &stakingJob, waitResourceJob, &pullJob, &deployJob)
	if err != nil {
		fmt.Println("new queue failed,err is: ", err)
		return QueueInfo{}, err
	}
	err = queue.LoadStatus(g.db)
	if err != nil {
		return QueueInfo{}, err
	}
	info, err := g.queueService.GetStatusInfo(applicationId)
	if err != nil {
		return QueueInfo{Info: info}, err
	}
	return QueueInfo{Info: info}, nil
}

func (g *ServiceImpl) GraphStart(appID int, deploymentID string) error {
	port, peerId, err := g.getP2pPort(appID)
	_ = g.p2pServer.LinkByProtocol("/x/provider", port, peerId)
	if err != nil {
		return err
	}
	keyringPair, err := g.walletService.GetWalletKeypair()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://localhost:%d/api/v1/thegraph/start", port)
	req := utils.NewHttp().NewRequest()
	req.SetHeader("SS58AuthData", utils.GetSS58AuthDataWithKeyringPair(keyringPair))
	resp, err := req.Get(fmt.Sprintf("%s?deploymentID=%s", url, deploymentID))
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		return nil
	}
	return parseResponseError(resp)
}

func (g *ServiceImpl) GraphRules(appID int) ([]GraphRule, error) {
	port, peerId, err := g.getP2pPort(appID)
	_ = g.p2pServer.LinkByProtocol("/x/provider", port, peerId)

	fmt.Println("#### p2p port : ", port)
	if err != nil {
		return nil, err
	}
	keyringPair, err := g.walletService.GetWalletKeypair()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("http://localhost:%d/api/v1/thegraph/rules", port)
	req := utils.NewHttp().NewRequest()
	req.SetHeader("SS58AuthData", utils.GetSS58AuthDataWithKeyringPair(keyringPair))
	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, parseResponseError(resp)
	}
	var result Result
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}

func parseResponseError(resp *resty.Response) error {
	var result Result
	err := json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return fmt.Errorf("%s", resp.Body())
	}
	return fmt.Errorf("%s", result.Message)
}

type Result struct {
	Code    uint64      `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Result  []GraphRule `json:"result"`
}

type GraphRule struct {
	AllocationAmount        string `json:"allocationAmount"`
	AllocationLifetime      string `json:"allocationLifetime"`
	AutoRenewal             bool   `json:"autoRenewal"`
	DecisionBasis           string `json:"decisionBasis"`
	Identifier              string `json:"identifier"`
	IdentifierType          string `json:"identifierType"`
	MaxAllocationPercentage string `json:"maxAllocationPercentage"`
	MaxSignal               string `json:"maxSignal"`
	MinAverageQueryFees     string `json:"minAverageQueryFees"`
	MinSignal               string `json:"minSignal"`
	MinStake                string `json:"minStake"`
	ParallelAllocations     string `json:"parallelAllocations"`
	RequireSupported        bool   `json:"requireSupported"`
}

func (g *ServiceImpl) getP2pPort(appID int) (int, string, error) {
	vo, err := g.applicationService.QueryApplicationById(appID)
	if err != nil {
		return 0, "", err
	}
	return vo.P2pForwardPort, vo.PeerId, nil
}
