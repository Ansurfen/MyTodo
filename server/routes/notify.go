package routes

import (
	"MyTodo/engine/v1/starter"
	middleware "MyTodo/middleware/handler"
	service "MyTodo/service/notify/v1"

	"github.com/prometheus/client_golang/prometheus"
)

type NotificationRoute struct{}

func (r *NotificationRoute) InstallNotificationRoute(g *starter.MyTodoServerGroup) {
	notifyRouter := g.Group("/notify")
	{
		notifyRouter.POST("/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "add",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyAdd))

		notifyRouter.POST("/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "del",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyDel))

		notifyRouter.GET("/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "get",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyGet))

		notifyRouter.POST("/detail",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "detail",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.NotifyGetDetial))

		notifyRouter.POST("/pub/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "pub_add",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyPubAdd))

		notifyRouter.POST("/pub/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "pub_del",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotfiyPubDel))

		notifyRouter.GET("/pub/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "pub_get",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyPubGet))

		notifyRouter.POST("/action/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "action_add",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyActionAdd))

		notifyRouter.POST("/action/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "action_del",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyActionDel))

		notifyRouter.POST("/action/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "action_get",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyActionGet))

		notifyRouter.POST("/action/commit",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "action_commit",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyActionCommit))

		notifyRouter.GET("/all",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "notify",
				Name:      "all",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.NotifyAll))
	}
}
