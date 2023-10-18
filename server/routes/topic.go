package routes

import (
	"MyTodo/engine/v1/starter"
	middleware "MyTodo/middleware/handler"
	service "MyTodo/service/topic/v1"

	"github.com/prometheus/client_golang/prometheus"
)

type TopicRoute struct{}

func (r *TopicRoute) InstallTopicRoute(g *starter.MyTodoServerGroup) {
	topicRouter := g.Group("/topic")
	{
		topicRouter.POST("/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "add",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.CreateTopic))

		topicRouter.POST("/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "del",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.DeleteTopic))

		topicRouter.GET("/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "get",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.GetSubscribedTopic))

		topicRouter.POST("/sub",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "sub",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.SubscirbeTopic))

		topicRouter.GET("/member/:id",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "member",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.GetTopicMember))

		topicRouter.POST("/invite")
		topicRouter.POST("/chat")

		topicRouter.POST("/history",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "history",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.HistoryTopic))

		topicRouter.POST("/grant",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "grant",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.GrantAdminTopic))

		topicRouter.POST("/ungrant",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "ungrant",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.UngrantAdminTopic))

		topicRouter.POST("/migrate",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "migrate",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.MigrateTopicAdmin))

		topicRouter.POST("/member/rm",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "remove",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.RemoveTopicMember))

		topicRouter.POST("/exit",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "topic",
				Name:      "exit",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.ExitTopic))
	}
}
