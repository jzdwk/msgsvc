/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package model

type Content struct {
	Service *Service
	Api     *Api
}

type Service struct {
	Name         string   `json:"name"`
	UpstreamName string   `json:"upstream_name"`
	Schema       string   `json:"schema"`
	EndPoints    []string `json:"end_points"`
}

type Api struct {
	Name   string `json:"name"`
	Path   string
	Method string
}
