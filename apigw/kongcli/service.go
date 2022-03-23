/*
@Time : 20-7-15
@Author : jzd
@Project: apigw
*/
package kongcli

import (
	"context"
	gokong "github.com/hbagdi/go-kong/kong"
	"github.com/sirupsen/logrus"
	"msgsvc/apigw/model"
)

func (kong *KongClientWrap) CreateService(service *model.Service) error {
	kongService := &gokong.Service{Name: &service.Name, Host: &service.UpstreamName, Protocol: &service.Schema}
	svc, err := kong.cli.Services.Create(context.Background(), kongService)
	if err != nil {
		return err
	}
	logrus.Infof("crate kong service success, uuid %v", *svc.ID)
	return nil
}
