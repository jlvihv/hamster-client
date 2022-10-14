package chain

import (
	log "github.com/sirupsen/logrus"
	"hamster-client/module/application"
	"hamster-client/module/chainhelper"
	"hamster-client/module/chainjob"
	"hamster-client/module/queue"
)

type Common struct {
	helper chainhelper.Helper
}

func NewCommon(helper chainhelper.Helper) *Common {
	return &Common{
		helper: helper,
	}
}

func (c *Common) deployJob(appData application.Application, deployType int) error {
	q, err := c.createQueue(appData, deployType)
	if err != nil {
		return err
	}
	done := make(chan struct{})
	go q.Start(done)
	<-done
	return nil
}

func (c *Common) deployJobWithQueue(q queue.Queue) error {
	done := make(chan struct{})
	go q.Start(done)
	<-done
	return nil
}

func (c *Common) createQueue(appData application.Application, deployType int) (queue.Queue, error) {
	appID := int(appData.ID)
	waitJob := chainjob.NewWaitResourceJob(appID, c.helper, deployType)
	log.Info("common deployParam: ")
	pullJob := chainjob.NewPullImageJob(appID, c.helper, deployType, nil)
	startJob := chainjob.NewStartJob(appID, c.helper, nil)
	q, err := queue.NewQueue(appID, c.helper.DB(), waitJob, pullJob, startJob)
	if err != nil {
		log.Errorf("new queue failed: %v", err)
		return nil, err
	}
	return q, nil
}
