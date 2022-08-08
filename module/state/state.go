package state

import (
	"errors"
	"github.com/google/uuid"
	"sync"
)

var tasks sync.Map

type StatusCode = int

const (
	None      StatusCode = iota // 0 无状态
	Running                     // 1 运行中
	Succeeded                   // 2 已成功
	Failed                      // 3 已失败
	Retrying                    // 4 重试中
)

type TaskItem interface {
	SetRunning()
	SetSucceeded()
	SetFailed()
	SetRetrying()
	State() StatusCode
}

type Task interface {
	State() Info
	Current() TaskItem
	Next() error
	Reset()
}

type Info struct {
	Pledge StatusCode `json:"pledge"`
	Match  StatusCode `json:"match"`
	Pull   StatusCode `json:"pull"`
	Deploy StatusCode `json:"deploy"`
}

type task struct {
	index int
	queue [4]TaskItem
}

func NewTask() (newTask Task, key string) {
	newTask = &task{
		queue: [4]TaskItem{
			&baseTask{},
			&baseTask{},
			&baseTask{},
			&baseTask{},
		},
	}
	key = uuid.New().String()
	tasks.Store(key, newTask)
	return
}

func GetTask(key string) (Task, error) {
	if v, ok := tasks.Load(key); ok {
		return v.(Task), nil
	}
	return nil, errors.New("task not found")
}

func (t *task) State() Info {
	return Info{
		Pledge: t.queue[0].State(),
		Match:  t.queue[1].State(),
		Pull:   t.queue[2].State(),
		Deploy: t.queue[3].State(),
	}
}

func (t *task) Current() TaskItem {
	return t.queue[t.index]
}

func (t *task) Next() error {
	if t.Current().State() == Succeeded {
		if t.index < 3 {
			t.index++
		}
		return nil
	} else {
		return errors.New("current task is not succeeded")
	}
}

func (t *task) Reset() {
	t.index = 0
	t.queue = [4]TaskItem{
		&baseTask{},
		&baseTask{},
		&baseTask{},
		&baseTask{},
	}
}

type baseTask struct {
	StatusCode StatusCode
}

func (b *baseTask) SetRunning() {
	b.StatusCode = Running
}

func (b *baseTask) SetSucceeded() {
	b.StatusCode = Succeeded
}

func (b *baseTask) SetFailed() {
	b.StatusCode = Failed
}

func (b *baseTask) SetRetrying() {
	b.StatusCode = Retrying
}

func (b *baseTask) State() StatusCode {
	return b.StatusCode
}
