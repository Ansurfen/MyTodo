package db

import (
	"MyTodo/engine/v1/cli"
	mongo "MyTodo/middleware/driver/mongo/v1"
)

var mongoDB *mongo.TodoMongo

func Mongo() *mongo.TodoMongo {
	return mongoDB
}

func init() {
	mongoDB = mongo.New(cli.Option.Mongo)
}
