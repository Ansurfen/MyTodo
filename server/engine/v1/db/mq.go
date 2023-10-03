package db

import (
	"MyTodo/engine/v1/cli"
	mq "MyTodo/middleware/driver/mq/v1"
)

var mqClient *mq.TodoRabbitMQ

func init() {
	mqClient = mq.New(cli.Option.MQ)
}

func MQ() *mq.TodoRabbitMQ {
	return mqClient
}
