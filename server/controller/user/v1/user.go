package controller

import (
	"MyTodo/dao"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	interfaces "MyTodo/interface"
	"MyTodo/model/bo/v1"
	"MyTodo/model/po/v1"
	"MyTodo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	ctx *gin.Context
	dao.UserDao
	FollowDao dao.Follow
	User      bo.User
}

func Get(ctx starter.TodoContext) (*UserController, error) {
	if con := ctx.Get(interfaces.ControllerTypeUser); con != nil {
		if uc, ok := con.(*UserController); ok {
			return uc, nil
		}
	}

	uc := &UserController{ctx: ctx.Context()}
	jwt := ctx.Context().GetHeader("x-token")

	if len(jwt) == 0 {
		return nil, utils.ErrJWTNotFound
	}

	_, claims, err := utils.ParseToken(jwt)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(claims.Id)
	if err != nil {
		return nil, err
	}

	uc.User, err = uc.FindByID(id)
	if err != nil {
		return nil, err
	}

	ctx.Put(interfaces.ControllerTypeUser, uc)
	return uc, nil
}

func GetBlank(ctx starter.TodoContext) (*UserController, error) {
	return &UserController{ctx: ctx.Context()}, nil
}

func (c *UserController) Follow(follower int) error {
	return c.FollowDao.Create(&po.Follow{
		UID:      c.User.ID,
		Follower: uint(follower),
	})
}

const SQLGetFriend = `
SELECT 
    CASE
        WHEN uid = %d THEN follower
        ELSE uid
    END AS friend
FROM
    follow
WHERE
    uid = %d OR follower = %d;`

func (c *UserController) Friends() (res []int) {
	if c.User.ID > 0 {
		db.SQL().RawExec(SQLGetFriend, c.User.ID, c.User.ID, c.User.ID).Scan(&res)
	}
	return res
}

const SQLGetFriendDetail = `
SELECT u.id, u.is_male, u.name
FROM user u
WHERE u.id IN (
    SELECT 
        CASE
            WHEN uid = %d THEN follower
            ELSE uid
        END AS friend
    FROM
        follow
    WHERE
        uid = %d OR follower = %d
);`

func (c *UserController) DetailedFriends() (res []po.User) {
	if c.User.ID > 0 {
		db.SQL().RawExec(SQLGetFriendDetail, c.User.ID, c.User.ID, c.User.ID).Scan(&res)
	}
	return res
}

func (c *UserController) BatchFindUserByUID(uids []uint) (res []bo.User, err error) {
	err = db.SQL().Find(&res, uids).Error
	return
}
