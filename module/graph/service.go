package graph

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hamster-client/module/application"
)

type ServiceImpl struct {
	ctx context.Context
	db  *gorm.DB
}

func NewServiceImpl(ctx context.Context, db *gorm.DB) ServiceImpl {
	return ServiceImpl{ctx, db}
}

func (g *ServiceImpl) SaveGraphParameter(data GraphParameter) error {
	var graphParams GraphParameter
	err := g.db.Preload("Application").Where("application_id = ? ", data.ApplicationId).First(&graphParams).Error
	if err != gorm.ErrRecordNotFound {
		g.db.Create(&data)
		return nil
	}
	return errors.New(fmt.Sprintf("graph param -> application :%s already exists", data.Application.Name))
}

func (g *ServiceImpl) QueryParamByApplyId(applicationId int) (GraphParameter, error) {
	var data GraphParameter
	err := g.db.Preload("Application").Where("application_id = ? ", applicationId).First(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (g *ServiceImpl) DeleteGraphAndParams(applicationId int) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Application").Where("application_id = ? ", applicationId).First(&GraphParameter{}).Error; err != nil {
			return err
		}
		if err := tx.Preload("Application").Where("application_id = ? ", applicationId).Delete(&GraphParameter{}).Error; err != nil {
			return err
		}
		if err := tx.Where("id = ? ", applicationId).First(&application.Application{}).Error; err != nil {
			return err
		}
		if err := tx.Debug().Where("id = ?", applicationId).Delete(&application.Application{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
