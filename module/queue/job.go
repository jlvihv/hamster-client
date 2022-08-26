package queue

type Job interface {
	InitStatus()
	Run(si chan StatusInfo) (StatusInfo, error)
	Status() StatusInfo
}
