/*
@Time : 20-9-15
@Author : jzd
@Project: apigw
*/
package kongcli

import (
	"msgsvc/apigw/model"
	"testing"
)

func TestServiceCreate(t *testing.T) {
	cli, _ := newKongClient()
	client := KongClientWrap{cli: cli}
	endpoints := []string{"localhost:8000", "localhost:8001"}
	service := model.Service{Name: "test_svc", UpstreamName: "test_upstream", Schema: "http", EndPoints: endpoints}
	if _, err := client.CreateService(&service); err != nil {
		t.Fatalf("failed to create kong service: %v", err)
	}
	t.Log("create kong service success")
}
