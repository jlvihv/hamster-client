package application

import (
	"gorm.io/gorm"
	"time"
)

type Application struct {
	ID           uint   `gorm:"primarykey"`
	Name         string `json:"name"`         //apply name
	Abbreviation string `json:"abbreviation"` //apply abbreviation
	Describe     string `json:"describe"`     //apply describe
	Status       int    `json:"status"`       //apply status 0: not deploy 1:deployed 2:ALL
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type AddApplicationParam struct {
	Name         string //apply name
	Abbreviation string //apply abbreviation
	Describe     string //apply describe
}

type UpdateApplicationParam struct {
	ID           uint   //application ID
	Name         string //apply name
	Abbreviation string //apply abbreviation
	Describe     string //apply describe
}

type PageApplicationVo struct {
	Item  []Application
	Total int64
}

type ApplyVo struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string //apply name
	Abbreviation string //apply abbreviation
	Describe     string //apply describe
}

type Service interface {
	AddApplication(application *AddApplicationParam) (bool, error)
	UpdateApplication(id int, name string, abbreviation string, des string) (bool, error)
	DeleteApplication(id int) (bool, error)
	QueryApplicationById(id int) (ApplyVo, error)
	ApplicationList(page, pageSize int, name string, status int) (PageApplicationVo, error)
	UpdateApplicationStatus(id, status int) error
}
