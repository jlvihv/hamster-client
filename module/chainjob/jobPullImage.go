package chainjob

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"hamster-client/module/chainhelper"
	"hamster-client/module/queue"
	"hamster-client/utils"
	"time"
)

// PullImageJob pull image job
type PullImageJob struct {
	appID      int
	si         queue.StatusInfo
	helper     chainhelper.Helper
	deployType int
}

func NewPullImageJob(appID int, helper chainhelper.Helper, deployType int) queue.Job {
	return &PullImageJob{
		appID:      appID,
		helper:     helper,
		deployType: deployType,
	}
}

func (j *PullImageJob) InitStatus() {
	j.si.Name = "Pull Image"
	j.si.Status = queue.None
}

func (j *PullImageJob) Run(si chan queue.StatusInfo) (queue.StatusInfo, error) {
	j.si.Status = queue.Running
	si <- j.si

	var err error

	// get p2p forward port
	p2pForwardPort, err := j.helper.GetP2pForwardPort(j.appID)
	if err != nil {
		log.Errorf("get p2p forward port error: %v", err)
		j.si.Error = err.Error()
		j.si.Status = queue.Failed
		si <- j.si
		return j.si, err
	}

	// send pull image request
	url := fmt.Sprintf("http://localhost:%d/api/v1/chains/pullImage", p2pForwardPort)

	for i := 0; i < 3; i++ {

		// 为什么需要传递post参数呢？忘记了，先注释掉
		//jsonStr, err := json.Marshal(data)
		//if err != nil {
		//	continue
		//}
		//fmt.Println("param: ", string(jsonStr))

		req := utils.NewHttp().NewRequest()
		//req.SetBody(data)
		response, err := req.Post(url)
		if err != nil {
			log.Errorf("send pull image request error: %v", err)
			log.Errorf("response: %v", string(response.Body()))
			j.si.Error = err.Error()
			fmt.Println(string(response.Body()))
			continue
		}
		if response.IsSuccess() {
			j.si.Status = queue.Succeeded
			j.si.Error = ""
			si <- j.si
			fmt.Println(
				"========================== chain pull image success ==========================",
			)
			return j.si, nil
		} else {
			log.Infof("pull image failed, response: %v", string(response.Body()))
			time.Sleep(3 * time.Second)
			continue
		}
	}

	j.si.Status = queue.Failed
	j.si.Error = "chain pull image failed: " + err.Error()
	si <- j.si
	return j.si, fmt.Errorf(j.si.Error)
}

func (j *PullImageJob) Status() queue.StatusInfo {
	return j.si
}
