package routes

import (
	"MyTodo/engine/v1/starter"
	middleware "MyTodo/middleware/handler"
	service "MyTodo/service/task/v1"

	"github.com/prometheus/client_golang/prometheus"
)

type TaskRoute struct{}

func (r *TaskRoute) InstallTaskRoute(g *starter.MyTodoServerGroup) {
	taskRouter := g.Group("/task")
	{
		taskRouter.POST("/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "task",
				Name:      "add",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.CreateTask))

		taskRouter.GET("/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "task",
				Name:      "get",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.GetTask))

		taskRouter.GET("/info",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "task",
				Name:      "info",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.InfoTask))

		taskRouter.POST("/commit",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "task",
				Name:      "commit",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.CommitTask))

		taskRouter.GET("/locate/:id",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "task",
				Name:      "locate",
				Help:      "counts view count",
			}),
			service.TaskLocateImage)

		taskRouter.GET("/image/:id",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "task",
				Name:      "image",
				Help:      "counts view count",
			}),
			service.TaskImage)

		taskRouter.POST("/perm_check",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "task",
				Name:      "perm_check",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.TaskHasPerm))
	}
}
