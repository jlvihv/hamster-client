package app

import (
	graphV2 "hamster-client/module/graph/v2"
)

type Queue struct {
	graphV2Service graphV2.Service
}

func NewQueueApp(service graphV2.Service) Queue {
	return Queue{
		graphV2Service: service,
	}
}

func (q *Queue) GetQueueInfo(id int) (graphV2.QueueInfo, error) {
	info, err := q.graphV2Service.GetQueueInfo(id)
	if err != nil {
		return graphV2.QueueInfo{}, err
	}
	return info, nil
}
