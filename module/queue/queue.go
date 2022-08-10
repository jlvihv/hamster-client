package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hamster-client/module/keystorage"
	"sync"
)

var queues sync.Map

type StatusCode = int

const (
	None      StatusCode = iota // 0
	Running                     // 1
	Succeeded                   // 2
	Failed                      // 3
)

type StatusInfo struct {
	Name   string     `json:"name,omitempty"`
	Status StatusCode `json:"status,omitempty"`
	Error  string     `json:"error,omitempty"`
}

type Queue interface {
	ID() string
	GetStatus() (info []StatusInfo, err error)
	Start(done chan struct{})
	Reset()
	SaveStatus(db *gorm.DB) error
	LoadStatus(db *gorm.DB) error
}

type queue struct {
	id            string
	statusInfoMap sync.Map
	jobs          []Job
	index         int
	mu            sync.Mutex
}

func NewQueue(id string, jobs ...Job) (q Queue, err error) {
	for i := range jobs {
		jobs[i].InitStatus()
	}
	if isJobDuplicate(jobs...) {
		err = errors.New("job duplicate")
		return
	}
	q = &queue{
		id:    id,
		jobs:  jobs,
		index: 0,
	}
	queues.Store(id, q)
	return
}

func GetQueue(id string) (q Queue, err error) {
	if v, ok := queues.Load(id); ok {
		return v.(Queue), nil
	}
	return nil, errors.New("queue not found")
}

func (q *queue) ID() string {
	return q.id
}

func (q *queue) Start(done chan struct{}) {
	if q.index == 0 {
		q.init()
	}
	statusInfo := make(chan StatusInfo)

	go func() {
		for {
			select {
			case si := <-statusInfo:
				name := q.jobs[q.index].Status().Name
				q.statusInfoMap.Store(name, si)
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
		si, err := j.Run(statusInfo)
		if err != nil {
			log.Errorf("job %s run failed: %s", j.Status().Name, err)
			break
		}
		q.statusInfoMap.Store(j.Status().Name, si)
		if si.Status != Succeeded {
			break
		}
	}
	done <- struct{}{}
}

func (q *queue) init() {
	log.Debugf("init queue %s", q.id)
	for _, j := range q.jobs {
		q.statusInfoMap.Store(j.Status().Name, j.Status())
	}
}

func (q *queue) GetStatus() (info []StatusInfo, err error) {
	for _, j := range q.jobs {
		statusInfo, ok := q.statusInfoMap.Load(j.Status().Name)
		if !ok {
			log.Errorf("status info not found for job %s", j.Status().Name)
			err = fmt.Errorf("job %s not found", j.Status().Name)
			return
		}
		si, ok := statusInfo.(StatusInfo)
		if !ok {
			log.Errorf("job %s status info not found", j.Status().Name)
			return nil, fmt.Errorf("job %s status info not found", j.Status().Name)
		}
		if si.Name == "" {
			si.Name = j.Status().Name
		}
		info = append(info, si)
	}
	return
}

type StatusStorage struct {
	Index      int          `json:"index,omitempty"`
	JobsStatus []StatusInfo `json:"jobsStatus,omitempty"`
}

func (q *queue) SaveStatus(db *gorm.DB) error {
	log.Debugf("save status to db for queue %s", q.id)
	info, err := q.GetStatus()
	if err != nil {
		log.Errorf("get status failed: %s", err)
		return err
	}
	statusStorage := StatusStorage{
		Index:      q.index,
		JobsStatus: info,
	}
	statusStorageJson, err := json.Marshal(statusStorage)
	if err != nil {
		log.Errorf("marshal status storage failed: %s", err)
		return err
	}
	keystorage.NewServiceImpl(context.Background(), db).Set(q.id, string(statusStorageJson))
	return nil
}

func (q *queue) LoadStatus(db *gorm.DB) error {
	log.Debugf("load status from db for queue %s", q.id)
	statusStorageJson := keystorage.NewServiceImpl(context.Background(), db).Get(q.id)
	if statusStorageJson == "" {
		log.Debugf("status storage not found for queue %s, skip", q.id)
		return nil
	}
	log.Debugf("load status info from db: %s", statusStorageJson)
	statusStorage := new(StatusStorage)
	err := json.Unmarshal([]byte(statusStorageJson), statusStorage)
	if err != nil {
		log.Errorf("unmarshal status storage failed: %s", err)
		return err
	}
	q.index = statusStorage.Index
	for _, si := range statusStorage.JobsStatus {
		q.statusInfoMap.Store(si.Name, si)
	}
	return nil
}

func (q *queue) Reset() {
	q.index = 0
	for _, j := range q.jobs {
		j.InitStatus()
	}
	q.init()
}

func isJobDuplicate(jobs ...Job) bool {
	names := make([]string, 0, len(jobs))
	for _, j := range jobs {
		names = append(names, j.Status().Name)
	}
	return duplicateInArray(names)
}

func duplicateInArray(arr []string) bool {
	visited := make(map[string]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}
