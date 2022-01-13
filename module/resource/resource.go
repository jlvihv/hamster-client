package resource

import (
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	gorm.Model
	PeerId      string    `json:"peerId"`
	Cpu         string    `json:"cpu"`
	Memory      string    `json:"memory"`
	SystemImage string    `json:"systemImage"`
	VmType      string    `json:"vmType"`
	Creator     string    `json:"creator"`
	ExpireTime  time.Time `json:"expireTime"`
	User        string    `json:"user"`
	Status      int       `json:"status"`
}

type Service interface {
	GetResourceList(publicKey string) ([]Resource, error)
}
