/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package kong

import (
	"encoding/json"
	"fmt"
	"github.com/gocraft/work"
	"msgsvc/common"
	"os"
	"os/signal"
)



func Receive() {
	// apigw poll
	redisPool := work.NewWorkerPool(common.Context{}, 10, common.ApigwNS, common.RedisPool)

	// Add middleware that will be executed for each job
	redisPool.Middleware((*common.Context).Log)
	// retry 3 times
	redisPool.JobWithOptions(common.ApigwCommitJob, work.JobOptions{MaxFails: 3, SkipDead: false}, (*common.Context).Handler)
	// Start processing jobs
	redisPool.Start()
	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
	// Stop the pool
	redisPool.Stop()
}


func Handler(job *work.Job) error {
	content := job.Args["content"]
}

