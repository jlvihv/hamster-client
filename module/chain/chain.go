package chain

import (
	"fmt"
	"hamster-client/module/queue"
	"hamster-client/utils"
	"time"
)

// pull image job
type ChainPullImageJob struct {
	si queue.StatusInfo
}

func (c *ChainPullImageJob) InitStatus() {
	c.si.Name = "ChainPullImage"
	c.si.Status = queue.None
}

func (c *ChainPullImageJob) Run(si chan queue.StatusInfo) (queue.StatusInfo, error) {
	c.si.Status = queue.Running
	si <- c.si

	// send pull image request
	url := fmt.Sprintf("http://localhost:%d/api/v1/chains/pullImage", 35003)

	// 有 3 次重试机会
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
			fmt.Println("========================== chain pull image success ==========================")
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

func (c *ChainPullImageJob) Status() queue.StatusInfo {
	return c.si
}

// container start job
type ChainStartJob struct {
	si queue.StatusInfo
}

func (c *ChainStartJob) InitStatus() {
	c.si.Name = "ChainStart"
	c.si.Status = queue.None
}

func (c *ChainStartJob) Run(si chan queue.StatusInfo) (queue.StatusInfo, error) {
	c.si.Status = queue.Running
	si <- c.si

	// send pull image request
	url := fmt.Sprintf("http://localhost:%d/api/v1/chains/start", 35003)

	// 有 3 次重试机会
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

func (c *ChainStartJob) Status() queue.StatusInfo {
	return c.si
}
