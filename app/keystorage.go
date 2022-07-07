package app

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hamster-client/module/keystorage"
)

type KeyStorage struct {
	ctx     context.Context
	service keystorage.Service
}

func NewKeyStorageApp(service *keystorage.Service) KeyStorage {
	return KeyStorage{
		ctx:     context.Background(),
		service: *service,
	}
}

func (self *KeyStorage) Get(key string) string {
	v := self.service.Get(key)
	if self.service.Err() != nil {
		runtime.LogErrorf(self.ctx, "kv storage get error: %s", self.service.Err())
		return ""
	}
	return v
}

func (self *KeyStorage) Set(key, value string) {
	self.service.Set(key, value)
	if self.service.Err() != nil {
		runtime.LogErrorf(self.ctx, "kv storage set error: %s", self.service.Err())
	}
}

func (self *KeyStorage) Delete(key string) {
	self.service.Delete(key)
	if self.service.Err() != nil {
		runtime.LogErrorf(self.ctx, "kv storage delete error: %s", self.service.Err())
	}
}
