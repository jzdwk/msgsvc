/*
@Time : 20-7-15
@Author : jzd
@Project: apigw
*/
package kongcli

import (
	"context"
	gokong "github.com/kong/go-kong/kong"
	"github.com/sirupsen/logrus"
	"msgsvc/apigw/model"
)

//must idempotent
func (kong *KongClientWrap) CreateService(service *model.Service) (string, error) {
	kongService := &gokong.Service{Name: &service.Name, Host: &service.UpstreamName, Protocol: &service.Schema}
	svc, err := kong.cli.Services.Create(context.Background(), kongService)
	if err != nil {
		return "", err
	}
	logrus.Infof("create kong service success, name [%v], uuid [%v]", *svc.Name, *svc.ID)
	return *svc.ID, nil
}

func (kong *KongClientWrap) DeleteService(uuid string) error {
	if err := kong.cli.Services.Delete(context.Background(), &uuid); err != nil {
		return err
	}
	logrus.Infof("delete kong service success,  uuid [%v]", uuid)
	return nil
}
