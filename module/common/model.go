package common

import (
	"gorm.io/gorm"
	"time"
)

type EthereumDeployParam struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Network       string    `json:"network"`       //rinkbey network or mainnet network
	LeaseTerm     int       `json:"leaseTerm"`     //
	ApplicationID uint      `json:"applicationId"` //application id
}

type StarkwareDeployParam struct {
	gorm.Model
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Network        string    `json:"network"`       //rinkbey network or mainnet network
	LeaseTerm      int       `json:"leaseTerm"`     //
	ApplicationID  uint      `json:"applicationId"` //application id
	EthereumApiUrl string    `json:"ethereumApiUrl"`
}
