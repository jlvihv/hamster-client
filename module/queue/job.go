package queue

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

var queues sync.Map

type StatusCode = int

const (
	None      StatusCode = iota // 0 无状态
	Running                     // 1 运行中
	Succeeded                   // 2 已成功
	Failed                      // 3 已失败
)

type Job interface {
	Run(sc chan StatusCode) (StatusCode, error)
	Name() string
	Error() error
}

type Queue interface {
	GetStatus() (info []StatusInfo, err error)
	Start(done chan struct{})
}

type queue struct {
	Status sync.Map
	jobs   []Job
	index  int
	mu     sync.Mutex
}

type StatusInfo struct {
	Name   string
	Status StatusCode
	Error  error
}

func NewQueue(jobs ...Job) (q Queue, key string) {
	q = &queue{
		jobs:  jobs,
		index: 0,
	}
	key = uuid.New().String()
	queues.Store(key, q)
	return
}

func GetQueue(key string) (q Queue, err error) {
	if v, ok := queues.Load(key); ok {
		return v.(Queue), nil
	}
	return nil, errors.New("queue not found")
}

func (q *queue) Start(done chan struct{}) {
	// index 等于 0 说明之前没有运行过，所以需要初始化，否则认为是失败重试，不需要初始化
	if q.index == 0 {
		q.init()
	}
	state := make(chan StatusCode)
	go func() {
		for {
			select {
			case s := <-state:
				q.Status.Store(q.jobs[q.index].Name(), s)
			}
		}
	}()
	for i, j := range q.jobs {
		q.mu.Lock()
		if i < q.index {
			q.mu.Unlock()
			continue
		}
		q.index = i
		q.mu.Unlock()
		statusCode, err := j.Run(state)
		if err != nil {
			fmt.Println(err)
			break
		}
		q.Status.Store(j.Name(), statusCode)
		if statusCode != Succeeded {
			break
		}
	}
	done <- struct{}{}
}

// init 方法用来初始化结构体中的Status字段，这是一个sync.Map结构，用来保存任务的名称和当前状态
func (q *queue) init() {
	for _, j := range q.jobs {
		name := j.Name()
		q.Status.Store(name, None)
	}
}

func (q *queue) GetStatus() (info []StatusInfo, err error) {
	time.Sleep(time.Second * 1)
	for _, j := range q.jobs {
		name := j.Name()
		status, ok := q.Status.Load(name)
		if !ok {
			err = fmt.Errorf("job %s not found", name)
			return
		}
		statusCode, ok := status.(StatusCode)
		if !ok {
			err = fmt.Errorf("job %s status not found", name)
			return
		}
		info = append(info, StatusInfo{Name: name, Status: statusCode, Error: j.Error()})
	}
	return
}
