package userController

import (
	dao "MyTodo/dao/v2"
	"MyTodo/model/bo/v1"
	"MyTodo/utils"
	"context"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	rkgrpcctx "github.com/rookie-ninja/rk-grpc/v2/middleware/context"
)

type UserController struct {
	ctx  context.Context
	User bo.User
}

func Get(ctx context.Context) (*UserController, error) {
	token := rkgrpcctx.GetJwtToken(ctx)
	if !token.Valid {
		return nil, utils.ErrJWTNotFound
	}
	uc := &UserController{ctx: ctx}
	id, err := strconv.Atoi(token.Claims.(jwt.MapClaims)["jti"].(string))
	if err != nil {
		return nil, err
	}
	u := dao.User
	po, err := u.WithContext(ctx).Where(u.ID.Eq(uint(id))).First()
	if err != nil {
		return nil, err
	}
	uc.User.User = *po
	return uc, nil
}
