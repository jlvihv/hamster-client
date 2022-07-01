package application

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	ctx context.Context
	db  *gorm.DB
}

func NewServiceImpl(ctx context.Context, db *gorm.DB) ServiceImpl {
	return ServiceImpl{ctx, db}
}

// AddApplication add application data
func (a *ServiceImpl) AddApplication(application *Application) error {
	var applyData Application
	err := a.db.Where("name=?", application.Name).First(&applyData).Error
	if err == gorm.ErrRecordNotFound {
		a.db.Create(&application)
		return nil
	}
	return errors.New(fmt.Sprintf("application:%s already exists", application.Name))
}

// UpdateApplication update application field
func (a *ServiceImpl) UpdateApplication(id int, name string, abbreviation string, des string) error {
	var applyData Application
	result := a.db.Model(applyData).Where("id = ?", id).Updates(Application{Name: name, Abbreviation: abbreviation, Describe: des})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteApplication delete application by id
func (a *ServiceImpl) DeleteApplication(id int) error {
	result := a.db.Debug().Where("id = ?", id).Delete(&Application{})
	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// QueryApplicationById query application data by id
func (a *ServiceImpl) QueryApplicationById(id int) (Application, error) {
	var data Application
	result := a.db.Where("id = ? ", id).First(&data)
	if result.Error != nil {
		return Application{}, result.Error
	}
	return data, nil
}

// ApplicationList Paging query application data
func (a *ServiceImpl) ApplicationList(page, pageSize int, name string, status int) (data []Application, count int64, err error) {
	var total int64
	var list []Application
	tx := a.db.Model(Application{})
	if name != "" {
		tx = tx.Where("name like ? ", "%"+name+"%")
	}
	if status != 2 {
		tx = tx.Where("status = ?", status)
	}
	result := tx.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Count(&total)
	if result.Error != nil {
		return nil, total, result.Error
	}
	return list, total, nil
}

// UpdateApplicationStatus update deploy status
func (a *ServiceImpl) UpdateApplicationStatus(id, status int) error {
	var applyData Application
	result := a.db.Model(applyData).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
