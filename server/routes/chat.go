package routes

import (
	"MyTodo/engine/v1/starter"
	"MyTodo/service/chat/v1"
)

type ChatRoute struct{}

func (r *ChatRoute) InstallChatRoutes(g *starter.MyTodoServerGroup) {
	chatRouter := g.Group("/chat")
	{
		chatRouter.POST("/add", starter.BindRequest(service.AddChat))
		chatRouter.POST("/del", starter.BindRequest(service.DelChat))
		chatRouter.POST("/get", starter.BindRequest(service.GetChat))
		chatRouter.GET("/friend", starter.BindRequest(service.Friend))
	}
}
