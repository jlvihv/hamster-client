package queue

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type helloJob struct {
	err error
}

func (h *helloJob) Run(sc chan StatusCode) (StatusCode, error) {
	sc <- Running
	time.Sleep(time.Second * 5)
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
	q, _ := NewQueue(&hello, &hi)
	done := make(chan struct{})
	// start queue, in a new goroutine
	go q.Start(done)
	// view status, in a new goroutine
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
	// wait
	<-done
	// view status
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
