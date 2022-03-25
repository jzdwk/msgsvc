/*
@Time : 2022/3/23
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

func (kong *KongClientWrap) CreateTarget(service *model.Service) ([]string, error) {
	var targetId []string
	for _, ep := range service.EndPoints {
		endpoint := ep
		kongTarget := &gokong.Target{Target: &endpoint}
		tg, err := kong.cli.Targets.Create(context.Background(), &service.UpstreamName, kongTarget)
		if err != nil {
			return nil, err
		}
		logrus.Infof("crate kong target success, uuid %v", *tg.ID)
		targetId = append(targetId, *tg.ID)
	}
	return targetId, nil
}

func (kong *KongClientWrap) DeleteTargets(upstream string, uuids []string) error {
	for _, uuid := range uuids {
		id := uuid
		if err := kong.cli.Targets.Delete(context.Background(), &upstream, &id); err != nil {
			return err
		}
		logrus.Infof("crate kong target success, uuid %v", id)
	}
	return nil
}
