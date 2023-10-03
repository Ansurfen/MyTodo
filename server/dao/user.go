package dao

import (
	"MyTodo/engine/v1/db"
	"MyTodo/model/bo/v1"
	"time"
)

type UserDao struct {
	model bo.User
}

func (u *UserDao) FindByID(id int) (user bo.User, err error) {
	err = db.SQL().Model(u.model).Where("id = ?", id).First(&user).Error
	return
}

func (u *UserDao) FindByEmail(email string) (user bo.User, err error) {
	err = db.SQL().Model(u.model).Where("email = ?", email).First(&user).Error
	return
}

func (u *UserDao) CreateUser(user *bo.User) error {
	return db.SQL().Create(user).Error
}

func (u *UserDao) DeleteUser(id uint) error {
	return db.SQL().Update("deleted_at", time.Now()).Where("id = ?", id).Error
}

func (u *UserDao) UpdateUser(id uint, v any) error {
	return db.SQL().Updates(v).Where("id = ?", id).Error
}
