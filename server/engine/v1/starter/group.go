package starter

import "github.com/gin-gonic/gin"

type MyTodoServerGroup struct {
	*gin.RouterGroup
}

func (g *MyTodoServerGroup) GET(relativePath string, handlers ...TodoHandler) {
	g.RouterGroup.GET(relativePath, buildHandlers(handlers...)...)
}

func (g *MyTodoServerGroup) POST(relativePath string, handlers ...TodoHandler) {
	g.RouterGroup.POST(relativePath, buildHandlers(handlers...)...)
}

func (g *MyTodoServerGroup) Group(relativePath string, handlers ...TodoHandler) *MyTodoServerGroup {
	return &MyTodoServerGroup{g.RouterGroup.Group(relativePath, buildHandlers(handlers...)...)}
}

func (s *MyTodoServer) Group(relativePath string, handlers ...TodoHandler) *MyTodoServerGroup {
	return &MyTodoServerGroup{s.Engine.Group(relativePath, buildHandlers(handlers...)...)}
}

func buildHandlers(handlers ...TodoHandler) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, len(handlers))
	for i := 0; i < len(handlers); i++ {
		index := i
		handler := handlers[index]
		ginHandlers[index] = func(fn TodoHandler) gin.HandlerFunc {
			return func(ctx *gin.Context) {
				fn(UpgradeContext(ctx))
			}
		}(handler)
	}
	return ginHandlers
}
