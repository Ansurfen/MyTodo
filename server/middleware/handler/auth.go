package middleware

import (
	userController "MyTodo/controller/user/v1"
	"MyTodo/engine/v1/starter"
)

func AuthJWT(ctx starter.TodoContext) {
	uc, err := userController.Get(ctx)
	if err == nil && uc.User.Exist() {
		return
		// return ctx.Context().JSON()
	}
	ctx.Context().Abort()
}
