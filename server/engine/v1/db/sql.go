package db

import (
	"MyTodo/engine/v1/cli"
	db "MyTodo/middleware/driver/sql/v1"
	"MyTodo/model/po/v1"
)

var todoDB *db.TodoDB

func SQL() *db.TodoDB {
	return todoDB
}

func init() {
	todoDB = db.New(cli.Option.SQL)
	if err := todoDB.AutoMigrate(
		po.Task{},
		po.TaskCond{},
		po.TaskBound{},
		po.TaskCommit{},
		po.User{},
		po.NotifyText{},
		po.NotifyPub{},
		po.NotifyAction{},
		po.Post{},
		po.PostFavorite{},
		po.PostCommentFavorite{},
		po.Subscribe{},
		po.Follow{},
		po.Topic{},
	); err != nil {
		panic(err)
	}
}
