/*
@Time : 2022/3/25
@Author : jzd
@Project: msgsvc
*/
package kongcli

import (
	"context"
	gokong "github.com/kong/go-kong/kong"
	"github.com/sirupsen/logrus"
	"msgsvc/apigw/model"
)

func (kong *KongClientWrap) CreateKeyAuthPlugin(service *model.Service) (string, error) {
	pluginName := auth
	keyAuth := &gokong.Plugin{Name: &pluginName, Service: &gokong.Service{Name: &service.Name}}
	auth, err := kong.cli.Plugins.Create(context.Background(), keyAuth)
	if err != nil {
		return "", err
	}
	logrus.Infof("create kong acl plugin success,  uuid [%v]", *auth.ID)
	return *auth.ID, nil
}
