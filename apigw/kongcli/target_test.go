/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package kongcli

import (
	"msgsvc/apigw/model"
	"testing"
)

func TestTargetCreate(t *testing.T) {
	cli, _ := newKongClient()
	client := KongClientWrap{cli: cli}
	endpoints := []string{"localhost:8000", "localhost:8001"}
	service := model.Service{Name: "test_svc", UpstreamName: "test_upstream", Schema: "http", EndPoints: endpoints}
	if err := client.CreateTarget(&service); err != nil {
		t.Fatalf("failed to create kong service: %v", err)
	}
	t.Log("create kong target success")
}
