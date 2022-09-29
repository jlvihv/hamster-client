package application

import (
	"gorm.io/gorm"
	"time"
)

type Application struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	Name           string         `json:"name"` //apply name
	ServiceType    string         `json:"serviceType"`
	SelectNodeType string         `json:"selectNodeType"` //apply plugin
	Status         int            `json:"status" gorm:"default 2"`
	P2pForwardPort int            `json:"p2pForwardPort"`
	CliForwardPort int            `json:"cliForwardPort"`
	GrtIncome      float64        `json:"grtIncome"`
	LeaseTerm      int            `json:"leaseTerm"`
	PeerId         string         `json:"peerId"`
	OrderIndex     int            `json:"orderIndex"`
	ResourceIndex  int            `json:"resourceIndex"`
	ThinkingTime   int            `json:"thinkingTime"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type ApplicationStatus = int

const DB_KEY_PREFIX = "queue_"

const (
	All              ApplicationStatus = iota // 0
	Running                                   //1
	Deploying                                 //2
	DeploymentFailed                          //3
	Offline                                   //4
)

const (
	TYPE_THEGRAPH string = "thegraph"
	TYPE_ETHEREUM string = "ethereum"
)

type AddApplicationParam struct {
	Name           string `json:"name"`           //apply name
	SelectNodeType string `json:"selectNodeType"` //apply plugin
}

type UpdateApplicationParam struct {
	ID             uint   `json:"id"`             //application ID
	Name           string `json:"name"`           //apply name
	SelectNodeType string `json:"selectNodeType"` //apply plugin
}

type PageApplicationVo struct {
	Items []ListVo `json:"items"`
	Total int64    `json:"total"`
}

type ApplyVo struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	Name           string    `json:"name"`           //apply name
	SelectNodeType string    `json:"selectNodeType"` //apply plugin
	Status         int       `json:"status"`
	LeaseTerm      int       `json:"leaseTerm"`
	P2pForwardPort int       `json:"p2pForwardPort"`
	CliForwardPort int       `json:"cliForwardPort"`
	PeerId         string    `json:"peerId"`
	OrderIndex     int       `json:"orderIndex"`
	ThinkingTime   int       `json:"thinkingTime"`
}

type ListVo struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`           //apply name
	SelectNodeType string  `json:"selectNodeType"` //apply plugin
	Status         int     `json:"status"`
	GrtIncome      float64 `json:"grtIncome"`
}

type Service interface {
	AddApplication(application *AddApplicationParam) (bool, error)
	UpdateApplication(id int, name string, plugin string) (bool, error)
	DeleteApplication(id int) (bool, error)
	QueryApplicationById(id int) (ApplyVo, error)
	ApplicationList(page, pageSize int, name string, status int) (PageApplicationVo, error)
	UpdateApplicationStatus(id, status int) error
	UpdateApplicationP2pForwardPort(id, port int) error
	QueryNextP2pPort() int
	QueryCliP2pPort(id int) (int, error)
	QueryNextCliP2pPort() int
	UpdateApplicationCliForwardPort(id, port int) error
	UpdatePeerIdAndOrderIndex(id, orderIndex, resourceIndex int, peerId string) error
	UpdateApplicationIncome(id int, income float64) (bool, error)
	UpdateThinkingTime(id, time int) (bool, error)
}
