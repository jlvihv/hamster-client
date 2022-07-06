package keystorage

import (
	"time"
)

type Service interface {
	Get(key string) string
	Set(key, value string)
	Delete(key string)
	Err() error
	SetTableName(string)
}

type KeyStorage struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Key   string `gorm:"uniqueIndex" json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
