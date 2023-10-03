package db

import (
	"MyTodo/engine/v1/cli"
	rdb "MyTodo/middleware/driver/redis/v1"
)

var redisDB *rdb.RedisDB

func Redis() *rdb.RedisDB {
	return redisDB
}

func init() {
	redisDB = rdb.New(cli.Option.Redis)
}
