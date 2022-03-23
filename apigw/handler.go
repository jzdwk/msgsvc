/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package apigw

import (
	"encoding/json"
	"github.com/gocraft/work"
	"msgsvc/apigw/kongcli"
	"msgsvc/apigw/manager"
	"msgsvc/apigw/model"
	"msgsvc/common"
	"os"
	"os/signal"
)

func Receive() {
	// apigw pool
	redisPool := work.NewWorkerPool(common.Context{}, 10, common.ApigwNS, common.RedisPool)
	// Add middleware that will be executed for each job
	redisPool.Middleware((*common.Context).Log)
	// retry 3 times, commit job
	redisPool.JobWithOptions(common.ApigwCommitJob, work.JobOptions{MaxFails: 3, SkipDead: false}, (*common.Context).CommitHandler)
	//rollback job
	redisPool.JobWithOptions(common.ApigwRollbackJob, work.JobOptions{MaxFails: 3, SkipDead: false}, (*common.Context).CommitHandler)
	// Start processing jobs
	redisPool.Start()
	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
	// Stop the pool
	redisPool.Stop()
}

type Handler struct {
	KongClient *kongcli.KongClientWrap
	Manager    manager.Manager
}

func (h *Handler) Commit(job *work.Job) error {
	contentJson := job.ArgString("content")
	var content model.Content
	if err := json.Unmarshal([]byte(contentJson), &content); err != nil {
		return err
	}
	//create kong service resource
	if content.Service != nil {
		h.Manager = manager.NewServiceMg(content.Service, h.KongClient)
	}
	return h.Manager.Create()
}

func (h *Handler) Rollback(job *work.Job) error {
	return nil
}
