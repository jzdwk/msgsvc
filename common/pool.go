/*
@Time : 2022/3/23
@Author : jzd
@Project: msgsvc
*/
package common

import "github.com/gomodule/redigo/redis"

var RedisPool *redis.Pool

func init() {
	RedisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", RedisAddr)
		},
	}
}
