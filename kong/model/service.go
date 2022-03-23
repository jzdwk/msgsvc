/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package model

type ServiceConfig struct {
	Name         string   `json:"name"`
	UpstreamName string   `json:"upstream_name"`
	Schema       string   `json:"schema"`
	EndPoints    []string `json:"end_points"`
}
