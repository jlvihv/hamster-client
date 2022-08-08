package app

import "hamster-client/module/queue"

type Queue struct {
	service queue.Service
}

func NewQueueApp(service queue.Service) Queue {
	return Queue{
		service: service,
	}
}

type QueueInfo struct {
	Info []queue.StatusInfo `json:"info"`
}

func (q *Queue) GetQueueInfo(key string) (QueueInfo, error) {
	info, err := q.service.GetStatusInfo(key)
	if err != nil {
		return QueueInfo{}, err
	}
	return QueueInfo{Info: info}, nil
}
