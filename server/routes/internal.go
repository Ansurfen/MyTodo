package routes

import (
	"MyTodo/engine/v1/starter"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type InternalRoute struct{}

func (r *InternalRoute) InstallInternalRoute(group *starter.MyTodoServerGroup) {
	internalRouter := group.RouterGroup.Group("/internal")
	{
		internalRouter.GET("/swagger/*any", Swagger())
		internalRouter.GET("/metrics", Metrics())
	}
}

// @Summary Get API Document
// @Description Get API Document
// @Tags System
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /internal/swagger/index.html [get]
func Swagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}

// @Summary Get Metrics
// @Description Get Metrics
// @Tags System
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Router /internal/metrics [get]
func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	}
}
