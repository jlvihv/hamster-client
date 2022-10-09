package chain

import (
	"encoding/json"
	"fmt"
	"hamster-client/module/queue"
	"hamster-client/utils"
	"time"
)

// PullImageJob pull image job
type PullImageJob struct {
	appID int
	si    queue.StatusInfo
	tools Service
}

func NewPullImageJob(appID int, tools Service) *PullImageJob {
	return &PullImageJob{
		appID: appID,
		tools: tools,
	}
}

func (c *PullImageJob) InitStatus() {
	c.si.Name = "ChainPullImage"
	c.si.Status = queue.None
}

func (c *PullImageJob) Run(si chan queue.StatusInfo) (queue.StatusInfo, error) {
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
	url := fmt.Sprintf("http://localhost:%d/api/v1/chains/pullImage", p2pForwardPort)

	for i := 0; i < 3; i++ {
		req := utils.NewHttp().NewRequest()
		data := c.tools.GetDeployParam(c.appID)
		json, err := json.Marshal(data)
		if err != nil {
			continue
		}
		fmt.Println("param: ", string(json))
		req.SetBody(c.tools.GetDeployParam(c.appID))
		response, err := req.Post(url)
		if err != nil {
			c.si.Error = err.Error()
			fmt.Println(string(response.Body()))
			continue
		}
		if response.IsSuccess() {
			c.si.Status = queue.Succeeded
			c.si.Error = ""
			si <- c.si
			fmt.Println(
				"========================== chain pull image success ==========================",
			)
			return c.si, nil
		} else {
			time.Sleep(3 * time.Second)
			continue
		}
	}

	c.si.Status = queue.Failed
	c.si.Error = "chain pull image failed"
	si <- c.si
	return c.si, fmt.Errorf(c.si.Error)
}

func (c *PullImageJob) Status() queue.StatusInfo {
	return c.si
}
