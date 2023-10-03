package starter

import (
	"MyTodo/utils"
	interfaces "MyTodo/interface"
	"github.com/gin-gonic/gin"
)

type TodoHandler func(c TodoContext)

type TodoServerHandlerWithRequest[T any] func(c TodoContext, req T) (interfaces.Response, error)

func BindRequest[T any](fn TodoServerHandlerWithRequest[T]) TodoHandler {
	return func(c TodoContext) {
		var req T
		if err := c.Context().Bind(&req); err != nil {
			c.Context().JSON(200, gin.H{
				"code": 500,
				"msg":  "bad request",
			})
			return
		}
		res, err := fn(c, req)
		if err != nil {
			switch err {
			case utils.ErrSkip:
			case utils.ErrComposeRes:
			default:
				c.Context().JSON(200, gin.H{
					"code": 500,
					"data": res,
					"msg":  err.Error(),
				})
			}
			return
		}
		c.Context().JSON(200, gin.H{
			"code": 200,
			"data": res,
		})
	}
}
