package dao

import (
	"MyTodo/engine/v1/db"
	"MyTodo/model/po/v1"
	"time"
)

type TextNotify struct {
	SQLBaseDao[po.NotifyText]
}

func (TextNotify) Delete(uid, nid int) error {
	return db.SQL().Model(po.NotifyText{}).Where("id = ? AND uid = ?", nid, uid).Update("deleted_at", time.Now()).Error
}

func (TextNotify) FindByCreator(creator int) (res []po.NotifyText, err error) {
	err = db.SQL().Where("uid = ?", creator).Find(&res).Error
	return
}

func (TextNotify) FindOneByID(id int) (res po.NotifyText, err error) {
	err = db.SQL().Unscoped().Where("id = ?", id).First(&res).Error
	return
}

type ActionNotify struct{}

func (ActionNotify) Create(n *po.NotifyAction) error {
	return db.SQL().Create(n).Error
}

func (ActionNotify) Delete(uid, id int) error {
	return db.SQL().Model(po.NotifyAction{}).Where("id = ? AND sender = ?", id, uid).Update("deleted_at", time.Now()).Error
}

func (ActionNotify) FindAllByUID(uid int) (res []po.NotifyAction, err error) {
	err = db.SQL().Where("receiver = ?", uid).Find(&res).Error
	return
}

func (ActionNotify) FindOneByID(id int) (res po.NotifyAction, err error) {
	err = db.SQL().Where("id = ?", id).First(&res).Error
	return
}

func (ActionNotify) Commit(id, uid, state int) error {
	return db.SQL().Model(po.NotifyAction{}).Where("id = ? AND receiver = ?", id, uid).Update("status", state).Error
}

type NotifyPub struct{}

func (NotifyPub) Create(n *po.NotifyPub) error {
	return db.SQL().Create(n).Error
}

func (NotifyPub) Delete(id int) error {
	return db.SQL().Model(po.NotifyPub{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (NotifyPub) FindByID(id int) (res po.NotifyPub, err error) {
	err = db.SQL().Where("id = ?", id).First(&res).Error
	return
}

func (NotifyPub) FindByCreator(id int) (res []po.NotifyPub, err error) {
	err = db.SQL().Where("uid = ?", id).Find(&res).Error
	return
}
