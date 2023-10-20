package routes

import (
	"MyTodo/engine/v1/starter"
	middleware "MyTodo/middleware/handler"
	service "MyTodo/service/chat/v1"

	"github.com/prometheus/client_golang/prometheus"
)

type ChatRoute struct{}

func (r *ChatRoute) InstallChatRoutes(g *starter.MyTodoServerGroup) {
	chatRouter := g.Group("/chat")
	{
		chatRouter.POST("/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "chat",
				Name:      "add",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.AddChat))

		chatRouter.POST("/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "chat",
				Name:      "del",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.DelChat))

		chatRouter.POST("/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "chat",
				Name:      "get",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.GetChat))

		chatRouter.GET("/friend",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "chat",
				Name:      "friend",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.Friend))
	}
}
