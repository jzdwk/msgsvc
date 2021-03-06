/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package manager

import (
	"bytes"
	"encoding/json"
	"msgsvc/apigw/kongcli"
	"msgsvc/apigw/model"
)

type ServiceMg struct {
	model       *model.Service
	kongCli     *kongcli.KongClientWrap
	callbackUrl string
}

func NewServiceMg(service *model.Service, kongCli *kongcli.KongClientWrap, callbackUrl string) Manager {
	return &ServiceMg{model: service, kongCli: kongCli, callbackUrl: callbackUrl}
}

//must idempotent
func (sm *ServiceMg) Create() error {
	var err error
	kongResource := KongResourceId{}
	defer func() {
		if err != nil {
			sm.kongCli.DeleteUpstream(kongResource.UpstreamUUID)
			sm.kongCli.DeleteService(kongResource.ServiceUUID)
			sm.failCallback()
		} else {
			sm.successCallback()
		}
	}()
	//1. create kong service
	if kongResource.ServiceUUID, err = sm.kongCli.CreateService(sm.model); err != nil {
		return err
	}
	//2. create kong upstream
	if kongResource.UpstreamUUID, err = sm.kongCli.CreateUpstream(sm.model); err != nil {
		return err
	}
	//3. create kong target
	if kongResource.TargetUUID, err = sm.kongCli.CreateTarget(sm.model); err != nil {
		return err
	}
	//4. auth type
	switch sm.model.AuthType {
	case "apikey":
		_, err = sm.kongCli.CreateAclPlugin(sm.model)
		if err != nil {
			return err
		}
		_, err = sm.kongCli.CreateKeyAuthPlugin(sm.model)
		if err != nil {
			return err
		}
		break
	case "none":
		break
	default:
		break
	}
	return err
}

func (sm *ServiceMg) Update() error {
	panic("implement me")
}

func (sm *ServiceMg) Delete() error {
	panic("implement me")
}

func (sm *ServiceMg) failCallback() {
	body, _ := json.Marshal(sm.model)
	bodyReader := bytes.NewReader(body)
	go callback(sm.callbackUrl+"/fail", bodyReader)
}

func (sm *ServiceMg) successCallback() {
	body, _ := json.Marshal(sm.model)
	bodyReader := bytes.NewReader(body)
	go callback(sm.callbackUrl+"/success", bodyReader)
}
