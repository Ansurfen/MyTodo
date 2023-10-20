package routes

import (
	"MyTodo/engine/v1/starter"
	middleware "MyTodo/middleware/handler"
	service "MyTodo/service/post/v1"

	"github.com/prometheus/client_golang/prometheus"
)

type PostRoute struct{}

func (r *PostRoute) InstallPostRoute(g *starter.MyTodoServerGroup) {
	postRouter := g.Group("/post")
	{
		postRouter.POST("/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "add",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.CreatePost))

		postRouter.GET("/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "get",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.GetPostV2))

		postRouter.POST("/edit",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "edit",
				Help:      "counts view count",
			}))

		postRouter.POST("/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "del",
				Help:      "counts view count",
			}))

		postRouter.GET("/detail/:id",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "detail",
				Help:      "counts view count",
			}), starter.BindRequest(service.PostDetail))

		postRouter.GET("/image/:id",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "image",
				Help:      "counts view count",
			}),
			service.PostImage)

		postRouter.POST("/comment/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_add",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.CreateCommentPost))

		postRouter.POST("/comment/edit",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_edit",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.EditCommentPost))

		postRouter.POST("/comment/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_del",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.DelCommentPost))

		postRouter.POST("/comment/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_get",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.GetCommentPost))

		postRouter.POST("/comment/reply/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_reply_add",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.PostCommentReplyCreate))

		postRouter.POST("/comment/reply/edit",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_reply_edit",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.PostCommentReplyEdit))

		postRouter.POST("/comment/reply/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_reply_del",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.PostCommentReplyDelete))

		postRouter.POST("/comment/favorite",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_favorite",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.PostCommentFavorite))

		postRouter.POST("/comment/unfavorite",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_unfavorite",
				Help:      "counts view count",
			}),
			middleware.AuthJWT,
			starter.BindRequest(service.PostCommentUnfavorite))

		postRouter.POST("/comment/favoriteCount",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "comment_favorite_count",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.PostCommentFavoriteCount))

		postRouter.POST("/favorite/get",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "favorite_get",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.GetPostFavorite))

		postRouter.POST("/favorite/add",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "favorite_add",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.FavoritePost))

		postRouter.POST("/favorite/del",
			middleware.PromCount(prometheus.CounterOpts{
				Namespace: "post",
				Name:      "favorite_del",
				Help:      "counts view count",
			}),
			starter.BindRequest(service.UnfavoritePost))
	}
}
