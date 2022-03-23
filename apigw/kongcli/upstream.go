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

func (kong *KongClientWrap) CreateUpstream(service *model.Service) error {
	kongUpstream := &gokong.Upstream{Name: &service.UpstreamName}
	up, err := kong.cli.Upstreams.Create(context.Background(), kongUpstream)
	if err != nil {
		return err
	}
	logrus.Infof("crate kong upstream success, uuid %v", *up.ID)
	return nil
}
