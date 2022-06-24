package app

import (
	"context"
	"fmt"
	"hamster-client/module/application"
	"hamster-client/module/graph"
)

type Application struct {
	ctx                context.Context
	applicationService application.Service
	graphService       graph.Service
}

func NewApplicationApp(service application.Service, graphService graph.Service) Application {
	return Application{
		applicationService: service,
		graphService:       graphService,
	}
}

func (a *Application) WailsInit(ctx context.Context) error {
	a.ctx = ctx
	return nil
}

func (a *Application) AddApplication(application application.Application) error {
	fmt.Println(application.Name)
	return a.applicationService.AddApplication(&application)
}

func (a *Application) UpdateApplication(application application.Application) error {
	return a.applicationService.UpdateApplication(int(application.ID), application.Name, application.Abbreviation, application.Describe)
}

func (a *Application) DeleteApplication(id int) error {
	return a.applicationService.DeleteApplication(id)
}

func (a *Application) ApplicationList(page, pageSize int, name string, status int) (data []application.Application, err error) {
	return a.applicationService.ApplicationList(page, pageSize, name, status)
}

func (a *Application) QueryApplicationById(id int) (application.Application, error) {
	return a.applicationService.QueryApplicationById(id)
}

func (a *Application) DeleteGraphAndParams(applicationId int) error {
	return a.graphService.DeleteGraphAndParams(applicationId)
}
