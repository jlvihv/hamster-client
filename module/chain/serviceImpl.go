package chain

import (
	"fmt"
	"hamster-client/config"
	"hamster-client/module/application"
	"hamster-client/module/p2p"
	"hamster-client/module/queue"

	"gorm.io/gorm"
)

type VersionVo struct {
	Version string `json:"version"`
}

type ServiceImpl struct {
	db             *gorm.DB
	app            application.Service
	p2p            p2p.Service
	q              queue.Queue
	getDeployParam func(appId int, db *gorm.DB) interface{}
}

func NewServiceImpl(db *gorm.DB, app application.Service, p2p p2p.Service, getDeployParam func(appId int, db *gorm.DB) interface{}) *ServiceImpl {
	return &ServiceImpl{
		db:             db,
		app:            app,
		p2p:            p2p,
		getDeployParam: getDeployParam,
	}
}

func (c *ServiceImpl) getP2pForwardPort(appID int) (int, error) {
	vo, err := c.app.QueryApplicationById(appID)
	if err != nil {
		fmt.Println("query application by id failed, err: ", err)
		return 0, err
	}

	fmt.Println("pull before: reForwardLink:", vo.PeerId)
	if _, err = c.p2p.GetSetting(); err != nil {
		_ = c.p2p.InitSetting()
	}

	err = reForwardLink(c.p2p, vo.P2pForwardPort, vo.PeerId)
	if err != nil {
		fmt.Println("reForwardLink failed, err: ", err)
		return 0, err
	}
	return vo.P2pForwardPort, nil
}

func reForwardLink(p2pService p2p.Service, port int, peerID string) error {
	protocol := config.ProviderProtocol
	err := p2pService.LinkByProtocol(protocol, port, peerID)
	return err
}

func (c *ServiceImpl) StartQueue(appID int) error {
	pullImageJob := NewPullImageJob(appID, c)
	startJob := NewStartJob(appID, c)
	q, err := queue.NewQueue(appID, c.db, pullImageJob, startJob)
	if err != nil {
		return err
	}
	c.q = q

	// start queue, in a new goroutine
	done := make(chan struct{})
	go q.Start(done)

	// wait
	<-done

	return nil
}

func (c *ServiceImpl) GetQueueInfo(appID int) (QueueInfo, error) {
	info, err := c.q.GetStatus()
	if err != nil {
		return QueueInfo{}, err
	}
	return QueueInfo{
		Info: info,
	}, nil
}

func (c *ServiceImpl) GetDeployParam(appId int) interface{} {

	if c.getDeployParam != nil {
		return c.getDeployParam(appId, c.db)
	}

	return nil

}
