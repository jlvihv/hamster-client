package queue

type Service interface {
	GetStatusInfo(id int) ([]StatusInfo, error)
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
