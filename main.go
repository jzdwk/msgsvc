/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package main

import (
	"github.com/sirupsen/logrus"
	"msgsvc/apigw"
	"msgsvc/apigw/kongcli"
	"msgsvc/common"
)

func main() {
	//init client
	if kongClient, err := kongcli.NewKongClient(); err == nil {
		common.KongHandler = &apigw.Handler{KongClient: kongClient}
	} else {
		logrus.Errorf("init client err, %v", err.Error())
		return
	}
	//init pool
}
