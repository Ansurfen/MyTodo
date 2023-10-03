package bo

import (
	"MyTodo/model/po/v1"
	"MyTodo/utils"
)

type User struct {
	po.User
}

func (u *User) IsLogout() bool {
	return u.DeletedAt.Valid
}

func (u *User) New(email, password string) User {
	return User{
		User: po.User{
			Name:     utils.RandString(8),
			Email:    email,
			Password: utils.MD5(password),
		}}
}

func (u *User) Login(password string) bool {
	return utils.MD5(password) == u.Password
}
