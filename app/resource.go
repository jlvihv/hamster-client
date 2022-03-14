package app

import (
	"context"
	"hamster-client/module/account"
	"hamster-client/module/resource"
)

type Resource struct {
	ctx             context.Context
	resourceService resource.Service
	accountService  account.Service
}

func NewResourceApp(resourceService resource.Service, accountService account.Service) Resource {
	return Resource{
		resourceService: resourceService,
		accountService:  accountService,
	}
}

func (s *Resource) WailsInit(ctx context.Context) error {
	s.ctx = ctx
	return nil
}

func (s *Resource) GetResources() ([]resource.Resource, error) {
	accountInfo, err := s.accountService.GetAccount()
	if err != nil {
		return nil, err
	}
	list, err := s.resourceService.GetResourceList(accountInfo.PublicKey)
	return list, err
}
