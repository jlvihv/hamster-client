package state

import "errors"

type Service interface {
	GetStateInfo(key string) (Info, error)
}

type ServiceImpl struct{}

func NewServiceImpl() Service {
	return &ServiceImpl{}
}

func (s *ServiceImpl) GetStateInfo(key string) (Info, error) {
	task, ok := tasks.Load(key)
	if !ok {
		return Info{}, errors.New("task not found")
	}
	return task.(Task).State(), nil
}
