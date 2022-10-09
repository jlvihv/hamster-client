package chain

import (
	"fmt"
	"hamster-client/module/queue"
	"hamster-client/utils"
	"time"
)

// StartJob container start job
type StartJob struct {
	appID int
	tools Service
	si    queue.StatusInfo
}

func NewStartJob(appID int, tools Service) *StartJob {
	return &StartJob{
		appID: appID,
		tools: tools,
	}
}

func (c *StartJob) InitStatus() {
	c.si.Name = "ChainStart"
	c.si.Status = queue.None
}

func (c *StartJob) Run(si chan queue.StatusInfo) (queue.StatusInfo, error) {
	c.si.Status = queue.Running
	si <- c.si

	// get p2p forward port
	p2pForwardPort, err := c.tools.getP2pForwardPort(c.appID)
	if err != nil {
		c.si.Error = err.Error()
		c.si.Status = queue.Failed
		si <- c.si
		return c.si, err
	}

	// send pull image request
	url := fmt.Sprintf("http://localhost:%d/api/v1/chains/start", p2pForwardPort)

	for i := 0; i < 3; i++ {
		req := utils.NewHttp().NewRequest()
		response, err := req.Get(url)
		if err != nil {
			c.si.Error = err.Error()
			fmt.Println(string(response.Body()))
			continue
		}
		if response.IsSuccess() {
			c.si.Status = queue.Succeeded
			c.si.Error = ""
			si <- c.si
			fmt.Println("========================== chain start success ==========================")
			return c.si, nil
		} else {
			time.Sleep(3 * time.Second)
			continue
		}
	}

	c.si.Status = queue.Failed
	c.si.Error = "chain start failed"
	si <- c.si
	return c.si, fmt.Errorf(c.si.Error)
}

func (c *StartJob) Status() queue.StatusInfo {
	return c.si
}
