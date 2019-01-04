package main

// 实例二

import (
	"hytx_sync/pkg/setting"
	"hytx_sync/models"
	"hytx_sync/pkg/logging"
	"hytx_sync/pkg/gredis"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var redisConn *redis.Pool

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	redisConn = gredis.RedisConn
}

func main() {
	// mysql model操作
	UserInterPushInfo := models.User{}
	last_id, err := UserInterPushInfo.GetLastId() // 查询匹配列表最后个id


	fmt.Println(last_id, err)

	// redis 操作
	conn := redisConn.Get()
	defer conn.Close()
	conn.Do("FLUSHALL")

	// 下面业务逻辑整理....
	// ...
	// ...
	// ...
}




