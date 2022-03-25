/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package apigw

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gocraft/work"
	"github.com/sirupsen/logrus"
	"msgsvc/apigw/kongcli"
	"msgsvc/apigw/manager"
	"msgsvc/apigw/model"
	"msgsvc/common"
)

func Start(ctx context.Context) {
	// apigw pool
	redisPool := work.NewWorkerPool(common.Context{}, 10, common.ApigwNS, common.RedisPool)
	// Add middleware that will be executed for each job
	redisPool.Middleware((*common.Context).Log)
	// retry 3 times, commit job
	redisPool.JobWithOptions(common.ApigwCommitJob, work.JobOptions{MaxFails: 3, SkipDead: false}, (*common.Context).CommitHandler)
	// Start processing jobs
	redisPool.Start()
	// Stop the pool
	select {
	case <-ctx.Done():
		logrus.Infof("apigw redis pool exits")
		redisPool.Stop()
	}
}

type Handler struct {
	KongClient *kongcli.KongClientWrap
	Manager    manager.Manager
}

func (h *Handler) Commit(job *work.Job) error {
	contentJson, _ := json.Marshal(job.Args["content"])
	var content model.Content
	if err := json.Unmarshal(contentJson, &content); err != nil {
		return err
	}
	//create kong service resource
	if content.Service != nil {
		h.Manager = manager.NewServiceMg(content.Service, h.KongClient, content.Callback)
	}
	var err error
	switch job.ArgString("operation") {
	case "create":
		err = h.Manager.Create()
	case "update":
		break
	case "delete":
		break
	default:
		err = errors.New("unsupported message operation")
	}
	return err
}

func (h *Handler) Rollback(job *work.Job) error {
	return nil
}
