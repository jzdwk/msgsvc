/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"msgsvc/apigw"
	"msgsvc/apigw/kongcli"
	"msgsvc/common"
	"os"
	"os/signal"
)

func main() {
	//init client
	if kongClient, err := kongcli.NewKongClient(); err == nil {
		common.KongHandler = &apigw.Handler{KongClient: kongClient}
	} else {
		logrus.Errorf("init kong client err, %v", err.Error())
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	go apigw.Start(ctx)
	// Wait for a signal to quit
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
	cancel()
}
