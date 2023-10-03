package notifyController

import (
	"MyTodo/dao"
	"MyTodo/engine/v1/starter"

	"github.com/gin-gonic/gin"
)

type NotifyController struct {
	ctx             *gin.Context
	NotifyDao       dao.TextNotify
	NotifyPubDao    dao.NotifyPub
	NotifyActionDao dao.ActionNotify
}

func Get(ctx starter.TodoContext) *NotifyController {
	return &NotifyController{ctx: ctx.Context()}
}
