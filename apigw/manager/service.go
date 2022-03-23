/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package manager

import (
	"msgsvc/apigw/kongcli"
	"msgsvc/apigw/model"
)

type ServiceMg struct {
	model   *model.Service
	kongCli *kongcli.KongClientWrap
}

func NewServiceMg(service *model.Service, kongCli *kongcli.KongClientWrap) Manager {
	return &ServiceMg{model: service, kongCli: kongCli}
}

func (s *ServiceMg) Create() error {
	//create kong service
	if err := s.kongCli.CreateService(s.model); err != nil {
		return err
	}
	if err := s.kongCli.CreateUpstream(s.model); err != nil {
		return err
	}
	if err := s.kongCli.CreateTarget(s.model); err != nil {
		return err
	}
	return nil
}

func (s *ServiceMg) Update() error {
	panic("implement me")
}

func (s *ServiceMg) Delete() error {
	panic("implement me")
}
