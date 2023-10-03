package routes

import (
	"MyTodo/engine/v1/starter"
	"MyTodo/middleware/handler"
	"MyTodo/service/notify/v1"
)

type NotificationRoute struct{}

func (r *NotificationRoute) InstallNotificationRoute(g *starter.MyTodoServerGroup) {
	notifyRouter := g.Group("/notify")
	{
		notifyRouter.POST("/add", middleware.AuthJWT, starter.BindRequest(service.NotifyAdd))
		notifyRouter.POST("/del", middleware.AuthJWT, starter.BindRequest(service.NotifyDel))
		notifyRouter.GET("/get", middleware.AuthJWT, starter.BindRequest(service.NotifyGet))
		notifyRouter.POST("/detail", starter.BindRequest(service.NotifyGetDetial))
		notifyRouter.POST("/pub/add", middleware.AuthJWT, starter.BindRequest(service.NotifyPubAdd))
		notifyRouter.POST("/pub/del", middleware.AuthJWT, starter.BindRequest(service.NotfiyPubDel))
		notifyRouter.GET("/pub/get", middleware.AuthJWT, starter.BindRequest(service.NotifyPubGet))
		notifyRouter.POST("/action/add", middleware.AuthJWT, starter.BindRequest(service.NotifyActionAdd))
		notifyRouter.POST("/action/del", middleware.AuthJWT, starter.BindRequest(service.NotifyActionDel))
		notifyRouter.POST("/action/get", middleware.AuthJWT, starter.BindRequest(service.NotifyActionGet))
		notifyRouter.POST("/action/commit", middleware.AuthJWT, starter.BindRequest(service.NotifyActionCommit))
		notifyRouter.GET("/all", middleware.AuthJWT, starter.BindRequest(service.NotifyAll))
	}
}
