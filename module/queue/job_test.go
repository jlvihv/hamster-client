package queue

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 需要实现Job接口
type helloJob struct {
	err error
}

func (h *helloJob) Run(sc chan StatusCode) (StatusCode, error) {
	// 首先要运行
	sc <- Running
	// 等待一段时间
	time.Sleep(time.Second * 5)
	// 假如运行成功
	sc <- Succeeded
	return Succeeded, nil
}

func (h *helloJob) Name() string {
	return "hello"
}

func (h *helloJob) Error() error {
	return h.err
}

type hiJob struct {
	err error
}

func (h *hiJob) Run(sc chan StatusCode) (StatusCode, error) {
	sc <- Running
	time.Sleep(time.Second * 5)
	sc <- Failed
	h.err = errors.New("hi job error")
	return Failed, nil
}

func (h *hiJob) Name() string {
	return "hi"
}

func (h *hiJob) Error() error {
	return h.err
}

func TestAll(t *testing.T) {
	hello := helloJob{}
	hi := hiJob{}
	q, _ := NewQueue([]Job{&hello, &hi}...)
	done := make(chan struct{})
	// 启动队列，开始执行任务，在一个新的goroutine中
	go q.Start(done)
	// 可以在这里查看状态
	go func() {
		for {
			time.Sleep(time.Second)
			info, err := q.GetStatus()
			if err != nil {
				t.Error(err)
			}
			for _, v := range info {
				fmt.Print(v, "; ")
			}
			fmt.Println()
		}
	}()
	// 等待任务执行完成
	<-done
	// 任务全部完成以后，再查看一次状态
	info, err := q.GetStatus()
	if err != nil {
		t.Error(err)
	}
	for _, v := range info {
		if v.Error != nil {
			fmt.Printf("%s job error: %s\n", v.Name, v.Error)
		} else {
			fmt.Printf("%s job succeed\n", v.Name)
		}
	}
}
