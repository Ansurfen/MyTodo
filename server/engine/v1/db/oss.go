package db

import (
	"MyTodo/engine/v1/cli"
	oss "MyTodo/middleware/driver/oss/v1"
)

var ossClient *oss.TodoMinio

func init() {
	ossClient = oss.New(cli.Option.Minio)
}

func OSS() *oss.TodoMinio {
	return ossClient
}