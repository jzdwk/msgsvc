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

	Apigw            = "apigw"
	ApigwNS          = Apigw + "_ns"
	ApigwCommitJob   = "kong_resource_commit"
	ApigwRollbackJob = "kong_resource_rollback"

	Lbserver = "lbserver"
)

var KongHandler Handler

type Handler interface {
	Commit(job *work.Job) error
	Rollback(job *work.Job) error
}

type Context struct {
	types   string
	msgId   string
	handler Handler
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	if _, ok := job.Args["message_id"]; ok {
		c.msgId = job.ArgString("message_id")
	}
	if _, ok := job.Args["message_type"]; ok {
		c.types = job.ArgString("message_type")
	}
	if c.msgId == "" || c.types == "" {
		jobInfo, _ := json.Marshal(job)
		return errors.New(fmt.Sprintf("msgsvc receives error job, %+v", string(jobInfo)))
	}
	logrus.Infof("msgsvc receives job, id [%v], type [%v]", c.msgId, c.types)
	return next()
}

func (c *Context) CommitHandler(job *work.Job, next work.NextMiddlewareFunc) error {
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

func (c *Context) RollbackHandler(job *work.Job, next work.NextMiddlewareFunc) error {
	// Extract arguments:
	switch c.types {
	case Apigw:
		c.handler = KongHandler
		return c.handler.Rollback(job)
	case Lbserver:
		break
	default:
		break
	}
	return nil
}
