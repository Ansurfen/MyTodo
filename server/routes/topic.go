package routes

import (
	"MyTodo/engine/v1/starter"
	"MyTodo/middleware/handler"
	"MyTodo/service/topic/v1"
)

type TopicRoute struct{}

func (r *TopicRoute) InstallTopicRoute(g *starter.MyTodoServerGroup) {
	topicRouter := g.Group("/topic")
	{
		topicRouter.POST("/add", middleware.AuthJWT, starter.BindRequest(service.CreateTopic))
		topicRouter.POST("/del", middleware.AuthJWT, starter.BindRequest(service.DeleteTopic))
		topicRouter.GET("/get", middleware.AuthJWT, starter.BindRequest(service.GetSubscribedTopic))
		topicRouter.POST("/sub", middleware.AuthJWT, starter.BindRequest(service.SubscirbeTopic))
		topicRouter.GET("/member/:id", starter.BindRequest(service.GetTopicMember))
		topicRouter.POST("/invite")
		topicRouter.POST("/chat")
		topicRouter.POST("/history", middleware.AuthJWT, starter.BindRequest(service.HistoryTopic))
		topicRouter.POST("/grant", middleware.AuthJWT, starter.BindRequest(service.GrantAdminTopic))
		topicRouter.POST("/ungrant", middleware.AuthJWT, starter.BindRequest(service.UngrantAdminTopic))
		topicRouter.POST("/migrate", middleware.AuthJWT, starter.BindRequest(service.MigrateTopicAdmin))
		topicRouter.POST("/member/rm", middleware.AuthJWT, starter.BindRequest(service.RemoveTopicMember))
		topicRouter.POST("/exit", middleware.AuthJWT, starter.BindRequest(service.ExitTopic))
	}
}
