/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package manager

type Manager interface {
	Create() error
	Update() error
	Delete() error
}

type KongResourceId struct {
	ServiceUUID  string
	UpstreamUUID string
	TargetUUID   []string

	RouteUUID string
}
