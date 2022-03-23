/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package kongcli

import (
	"context"
	gokong "github.com/hbagdi/go-kong/kong"
	"github.com/sirupsen/logrus"
	"msgsvc/apigw/model"
)

func (kong *KongClientWrap) CreateTarget(service *model.Service) error {
	for _, ep := range service.EndPoints {
		endpoint := ep
		kongTarget := &gokong.Target{Target: &endpoint}
		tg, err := kong.cli.Targets.Create(context.Background(), &service.UpstreamName, kongTarget)
		if err != nil {
			return err
		}
		logrus.Infof("crate kong target success, uuid %v", *tg.ID)
	}
	return nil
}
