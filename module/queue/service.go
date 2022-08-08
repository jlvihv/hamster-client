package queue

type Service interface {
	GetStatusInfo(key string) ([]StatusInfo, error)
}

type ServiceImpl struct{}

func NewServiceImpl() Service {
	return &ServiceImpl{}
}

func (s *ServiceImpl) GetStatusInfo(key string) ([]StatusInfo, error) {
	q, err := GetQueue(key)
	if err != nil {
		return nil, err
	}
	return q.(Queue).GetStatus()
}
