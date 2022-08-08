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

func (q *Queue) GetQueueInfo(key string) ([]queue.StatusInfo, error) {
	info, err := q.service.GetStatusInfo(key)
	if err != nil {
		return []queue.StatusInfo{}, err
	}
	return info, nil
}
