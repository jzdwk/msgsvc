/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocraft/work"
	"github.com/sirupsen/logrus"
)

const (
	RedisAddr = "myecs.jzd:65079"

	Apigw          = "apigw"
	ApigwNS        = Apigw + "_ns"
	ApigwCommitJob = "kong_resource_commit"

	Lbserver = "lbserver"
)

var KongHandler Handler

type Handler interface {
	Commit(job *work.Job) error
}

type Context struct {
	types   string
	msgId   string
	handler Handler
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	if _, ok := job.Args["id"]; ok {
		c.msgId = job.ArgString("id")
	}
	if _, ok := job.Args["type"]; ok {
		c.types = job.ArgString("type")
	}
	if c.msgId == "" || c.types == "" {
		jobInfo, _ := json.Marshal(job)
		return errors.New(fmt.Sprintf("msgsvc receives error job, %+v", string(jobInfo)))
	}
	logrus.Infof("msgsvc receives job, id [%v], type [%v]", c.msgId, c.types)
	return next()
}

func (c *Context) CommitHandler(job *work.Job) error {
	// Extract arguments:
	switch c.types {
	case Apigw:
		c.handler = KongHandler
		break
	case Lbserver:
		break
	default:
		break
	}
	return c.handler.Commit(job)
}
