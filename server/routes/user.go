package routes

import (
	"MyTodo/engine/v1/starter"
	middleware "MyTodo/middleware/handler"
	"MyTodo/service/user/v1"

	"github.com/prometheus/client_golang/prometheus"
)

type UserRoute struct{}

func (r *UserRoute) InstallUserRoute(g *starter.MyTodoServerGroup) {
	userRouter := g.Group("/user")
	{
		userRouter.POST("/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "user",
				Name:      "get",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.UserGet))

		userRouter.POST("/sign",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "user",
				Name:      "sign",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.UserSign))

		userRouter.POST("/edit",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "user",
				Name:      "edit",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.UserEdit))

		userRouter.GET("/profile/:id",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "user",
				Name:      "profile",
				Help:      "counts view count",
			}),
			service.UserProfile)

		userRouter.GET("/info/:id",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "user",
				Name:      "info",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.UserInfo))
	}
}
