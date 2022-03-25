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
	var err error
	kongResource := KongResourceId{}
	//create kong service
	if kongResource.ServiceUUID, err = s.kongCli.CreateService(s.model); err != nil {
		return err
	}
	if kongResource.UpstreamUUID, err = s.kongCli.CreateUpstream(s.model); err != nil {
		_ = s.kongCli.DeleteService(kongResource.ServiceUUID)
		return err
	}
	if kongResource.TargetUUID, err = s.kongCli.CreateTarget(s.model); err != nil {
		_ = s.kongCli.DeleteUpstream(kongResource.UpstreamUUID)
		_ = s.kongCli.DeleteService(kongResource.ServiceUUID)
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
