package routes

import (
	"MyTodo/engine/v1/starter"
	"MyTodo/middleware/handler"
	"MyTodo/service/task/v1"
	"path/filepath"
)

type TaskRoute struct{}

func (r *TaskRoute) InstallTaskRoute(g *starter.MyTodoServerGroup) {
	taskRouter := g.Group("/task")
	{
		taskRouter.POST("/add", middleware.AuthJWT, starter.BindRequest(service.CreateTask))
		taskRouter.GET("/get", middleware.AuthJWT, starter.BindRequest(service.GetTask))
		taskRouter.GET("/info", middleware.AuthJWT, starter.BindRequest(service.InfoTask))
		taskRouter.POST("/commit", middleware.AuthJWT, starter.BindRequest(service.CommitTask))
		taskRouter.GET("/locate/:id", func(c starter.TodoContext) {
			id := c.Context().Param("id")
			if len(id) > 0 {
				c.Context().File(filepath.Join("./assets/task", id+".png"))
			}
		})
		taskRouter.GET("/image/:id", func(c starter.TodoContext) {
			id := c.Context().Param("id")
			if len(id) > 0 {
				c.Context().File(filepath.Join("./assets/task/file", id+".png"))
			}
		})
		taskRouter.POST("/perm_check", starter.BindRequest(service.TaskHasPerm))
	}
}
