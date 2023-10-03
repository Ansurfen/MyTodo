package routes

import (
	"MyTodo/engine/v1/starter"
	"MyTodo/middleware/handler"
	"MyTodo/service/post/v1"
)

type PostRoute struct{}

func (r *PostRoute) InstallPostRoute(g *starter.MyTodoServerGroup) {
	postRouter := g.Group("/post")
	{
		postRouter.POST("/add", middleware.AuthJWT, starter.BindRequest(service.CreatePost))
		postRouter.GET("/get", starter.BindRequest(service.GetPostV2))
		postRouter.POST("/edit")
		postRouter.POST("/del")
		postRouter.GET("/detail/:id", starter.BindRequest(service.PostDetail))
		postRouter.GET("/image/:id", service.PostImage)
		postRouter.POST("/comment/add", middleware.AuthJWT, starter.BindRequest(service.CreateCommentPost))
		postRouter.POST("/comment/edit", middleware.AuthJWT, starter.BindRequest(service.EditCommentPost))
		postRouter.POST("/comment/del", middleware.AuthJWT, starter.BindRequest(service.DelCommentPost))
		postRouter.POST("/comment/get", starter.BindRequest(service.GetCommentPost))
		postRouter.POST("/comment/reply/add", starter.BindRequest(service.PostCommentReplyCreate))
		postRouter.POST("/comment/reply/edit", starter.BindRequest(service.PostCommentReplyEdit))
		postRouter.POST("/comment/reply/del", starter.BindRequest(service.PostCommentReplyDelete))
		postRouter.POST("/comment/favorite", middleware.AuthJWT, starter.BindRequest(service.PostCommentFavorite))
		postRouter.POST("/comment/unfavorite", middleware.AuthJWT, starter.BindRequest(service.PostCommentUnfavorite))
		postRouter.POST("/comment/favoriteCount", starter.BindRequest(service.PostCommentFavoriteCount))
		postRouter.POST("/favorite/get", starter.BindRequest(service.GetPostFavorite))
		postRouter.POST("/favorite/add", starter.BindRequest(service.FavoritePost))
		postRouter.POST("/favorite/del", starter.BindRequest(service.UnfavoritePost))
	}
}
