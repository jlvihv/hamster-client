package v2

import (
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	queue2 "hamster-client/module/queue"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {

	pullJob := PullImageJob{
		ProviderApi: "http://localhost:34002",
	}
	substrateApi, _ := gsrpc.NewSubstrateAPI("ws://183.66.65.207:49944")

	job2, _ := NewWaitResourceJob(substrateApi)
	queue := queue2.NewQueue("1", &pullJob, job2)

	channel := make(chan struct{})
	go queue.Start(channel)
	go func() {
		for {
			time.Sleep(time.Second)
			info, err := queue.GetStatus()
			if err != nil {
				t.Error(err)
			}
			for _, v := range info {
				fmt.Print(v, "; ")
			}
			fmt.Println()
		}
	}()
	// wait
	<-channel
	// view status
	info, _ := queue.GetStatus()

	fmt.Println(info)
}
