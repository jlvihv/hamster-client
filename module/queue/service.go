package queue

type Service interface {
	GetStatusInfo(id int) ([]StatusInfo, error)
	StopQueue(id int) error
}

type ServiceImpl struct{}

func NewServiceImpl() Service {
	return &ServiceImpl{}
}

func (s *ServiceImpl) GetStatusInfo(id int) ([]StatusInfo, error) {
	q, err := GetQueue(id)
	if err != nil {
		return nil, err
	}
	return q.(Queue).GetStatus()
}

func (s *ServiceImpl) StopQueue(id int) error {
	q, err := GetQueue(id)
	if err != nil {
		return nil
	}

	err = q.Stop()
	if err != nil {
		return err
	}
	queues.Delete(id)
	return nil
}
