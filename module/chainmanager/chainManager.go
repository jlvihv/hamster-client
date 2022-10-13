package chainmanager

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/chain"
	"hamster-client/module/chainhelper"
	"hamster-client/module/keystorage"
	"hamster-client/module/p2p"
	"hamster-client/module/pallet"
	"hamster-client/module/queue"
	"hamster-client/module/wallet"
	"strconv"
)

type VersionVo struct {
	Version string `json:"version"`
}

type ChainManager struct {
	db      *gorm.DB
	app     application.Service
	p2p     p2p.Service
	account account.Service
	wallet  wallet.Service
	queue   queue.Service
	ks      keystorage.Service
}

func NewManager(
	db *gorm.DB,
	app application.Service,
	p2p p2p.Service,
	account account.Service,
	wallet wallet.Service,
	queue queue.Service,
	ks keystorage.Service,
) Manager {
	return &ChainManager{
		db:      db,
		app:     app,
		p2p:     p2p,
		account: account,
		wallet:  wallet,
		queue:   queue,
		ks:      ks,
	}
}

func (c *ChainManager) getChain(serviceType string) (chain.Chain, error) {
	helper := chainhelper.NewHelper(c.db, c.app, c.p2p, c.account, c.wallet, c.queue, c.ks)
	chainCommon := chain.NewCommon(helper)
	var ch chain.Chain
	switch serviceType {
	case application.TYPE_Thegraph:
	case application.TYPE_Aptos:
	case application.TYPE_Sui:
	case application.TYPE_Ethereum:
		ch = chain.NewEthereum(chainCommon, application.VALUE_Ethereum)
	case application.TYPE_BSC:
	case application.TYPE_Polygon:
	case application.TYPE_Avalanche:
	case application.TYPE_Optimism:
	case application.TYPE_StarkWare:
		ch = chain.NewStarkWare(chainCommon, application.VALUE_StarkWare)
	case application.TYPE_Near:
	case application.TYPE_Cess:
	default:
		return nil, fmt.Errorf("unknown service type: %s", serviceType)
	}
	return ch, nil
}

func (c *ChainManager) CreateAndStart(param chain.DeployParam) (chain.DeployResult, error) {
	helper := chainhelper.NewHelper(c.db, c.app, c.p2p, c.account, c.wallet, c.queue, c.ks)
	ch, err := c.getChain(param.ServiceType)
	if err != nil {
		return chain.DeployResult{}, err
	}
	var appInfo application.Application
	var deployResult chain.DeployResult
	err = helper.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("name=?", param.Name).First(&appInfo).Error; err != gorm.ErrRecordNotFound {
			return fmt.Errorf("application:%s already exists", param.Name)
		}
		appInfo.Name = param.Name
		appInfo.ServiceType = param.ServiceType
		appInfo.SelectNodeType = param.SelectNodeType
		appInfo.LeaseTerm = param.LeaseTerm
		appInfo.Status = application.Deploying
		if err := tx.Create(&appInfo).Error; err != nil {
			return err
		}
		if err := ch.SaveDeployParam(appInfo, param, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return deployResult, err
	}
	err = ch.SaveJsonParam(strconv.Itoa(int(appInfo.ID)), param)
	if err != nil {
		return deployResult, err
	}
	go ch.DeployJob(appInfo)
	deployResult.Result = true
	deployResult.ID = appInfo.ID
	return deployResult, nil
}

func (c *ChainManager) GetQueueInfo(appID int) (QueueInfo, error) {
	q, err := queue.GetQueue(appID)

	if err != nil {
		return QueueInfo{}, fmt.Errorf("queue %d not found", appID)
	}
	info, err := q.GetStatus()
	if err != nil {
		return QueueInfo{}, err
	}
	return QueueInfo{
		Info: info,
	}, nil
}

func (c *ChainManager) RetryStartQueue(appID int, runNow bool) error {
	var q queue.Queue
	var err error
	q, err = queue.GetQueue(appID)
	if err != nil {
		log.Infof("queue %d not found, create", appID)
		// query application data
		var appInfo application.Application
		result := c.db.Where("id=?", appID).First(&appInfo)
		if result.Error != nil {
			log.Errorf("query application %d error: %s", appID, result.Error)
			return result.Error
		}
		ch, err := c.getChain(appInfo.ServiceType)
		if err != nil {
			log.Errorf("get chain error: %s", err)
			return err
		}
		q, err = ch.CreateQueue(appInfo)
		if err != nil {
			log.Errorf("create queue error: %s", err)
			return err
		}
	}
	if runNow {
		done := make(chan struct{})
		go q.Start(done)
		<-done
	} else {
		statusInfo, err := q.GetStatus()
		if err != nil {
			log.Errorf("get queue status error: %s", err)
			return err
		}
		for _, job := range statusInfo {
			if job.Status == queue.Running {
				log.Infof("job %s status is running, set to failed", job.Name)
				q.SetJobStatus(job.Name, queue.StatusInfo{
					Name:   job.Name,
					Status: queue.Failed,
					Error:  "Abnormal exit",
				})
			}
		}
	}
	return nil
}

func (c *ChainManager) Delete(appID int) (bool, error) {
	helper := chainhelper.NewHelper(c.db, c.app, c.p2p, c.account, c.wallet, c.queue, c.ks)
	err := helper.Queue().StopQueue(appID)
	if err != nil {
		return false, err
	}
	app, err := helper.App().QueryApplicationById(appID)
	if err != nil {
		return false, err
	}
	err = helper.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Where("id = ?", appID).Delete(&application.Application{}).Error; err != nil {
			return err
		}
		if err := tx.Debug().Where("application_id = ?", appID).Delete(chain.GraphDeployParameter{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	keypair, err := helper.Wallet().GetWalletKeypair()
	if err != nil {
		return false, err
	}
	if api, err := helper.Account().GetSubstrateApi(); err != nil {
		if meta, err := api.RPC.State.GetMetadataLatest(); err != nil {
			_ = pallet.CancelOrder(api, meta, keypair, app.OrderIndex)
		}
	}
	return true, nil
}
