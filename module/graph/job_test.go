package graph

import (
	"fmt"
	queue2 "hamster-client/module/queue"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {

	pullJob := PullImageJob{
		ProviderApi: "http://localhost:34002",
	}
	queue, _ := queue2.NewQueue(&pullJob)
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
