package application

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	Name         string `json:"name"`         //apply name
	Abbreviation string `json:"abbreviation"` //apply abbreviation
	Describe     string `json:"describe"`     //apply describe
	Status       int    `json:"status"`       //apply status 0: not deploy 1:deployed 2:ALL
}

type Service interface {
	AddApplication(application *Application) error
	UpdateApplication(id int, name string, abbreviation string, des string) error
	DeleteApplication(id int) error
	QueryApplicationById(id int) (Application, error)
	ApplicationList(page, pageSize int, name string, status int) (data []Application, err error)
	UpdateApplicationStatus(id, status int) error
}
