package controller

import (
	"MyTodo/dao"
	"MyTodo/engine/v1/starter"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	ctx     *gin.Context
	ChatDao dao.Chat
}

func Get(ctx starter.TodoContext) *ChatController {
	return &ChatController{
		ctx: ctx.Context(),
	}
}
