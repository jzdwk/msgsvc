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

func (kong *KongClientWrap) CreateAclPlugin(service *model.Service) (string, error) {
	pluginName := acl
	config := make(map[string]interface{})
	config["whitelist"] = service.Name
	kongAcl := &gokong.Plugin{Name: &pluginName, Service: &gokong.Service{Name: &service.Name}, Config: config}
	acl, err := kong.cli.Plugins.Create(context.Background(), kongAcl)
	if err != nil {
		return "", err
	}
	logrus.Infof("create kong acl plugin success,  uuid [%v]", *acl.ID)
	return *acl.ID, nil
}
