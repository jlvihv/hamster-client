package keystorage

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	tableName string
	db        *gorm.DB
	Error     error
}

func NewServiceImpl(ctx context.Context, db *gorm.DB) Service {
	return &ServiceImpl{
		db:    db,
		Error: nil,
	}
}

func (self *ServiceImpl) Get(key string) string {
	self.autoMigrate()
	if self == nil || self.Error != nil {
		return ""
	}
	var result KeyStorage
	err := self.db.Table(self.tableName).Where("key = ?", key).First(&result).Error
	if err != nil {
		self.Error = err
		return ""
	}
	return result.Value
}

func (self *ServiceImpl) Set(key, value string) {
	self.autoMigrate()
	if self == nil || self.Error != nil {
		return
	}
	k := KeyStorage{
		Key:   key,
		Value: value,
	}
	self.Get(key)
	if self.Error != nil && errors.Is(self.Error, gorm.ErrRecordNotFound) {
		self.Error = nil
		err := self.db.Table(self.tableName).Create(&k).Error
		if err != nil {
			self.Error = err
			return
		}
		return
	}
	err := self.db.Table(self.tableName).Save(&k).Error
	if err != nil {
		self.Error = err
		return
	}
}

func (self *ServiceImpl) Err() error {
	return self.Error
}

func (self *ServiceImpl) SetTableName(name string) {
	self.tableName = name
}

func (self *ServiceImpl) autoMigrate() {
	if self == nil {
		return
	}
	if self.tableName == "" {
		self.tableName = "key_storage"
	}
	if self.db.Migrator().HasTable(self.tableName) {
		return
	}
	err := self.db.Table(self.tableName).AutoMigrate(&KeyStorage{})
	if err != nil {
		self.Error = err
		return
	}
}
