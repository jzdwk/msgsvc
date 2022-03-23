/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package main

import (
	"msgsvc/common"
	"msgsvc/kong"
	kongcli "msgsvc/kong/httpcli"
)

var KoHandler common.Handler

func main() {
	//init client
	if kongClient, err := kongcli.NewKongClient(); err == nil {
		KoHandler = &kong.Handler{KongClient: kongClient}
	}
	//init pool
}
