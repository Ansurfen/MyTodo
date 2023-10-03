package dao

import (
	"MyTodo/engine/v1/db"
	"MyTodo/model/po/v1"
	"time"
)

type TopicDao struct {
	model po.Topic
}

func (TopicDao) Create(t *po.Topic) error {
	return db.SQL().Create(t).Error
}

func (t TopicDao) FindAll(creator int) (res []*po.Topic, err error) {
	err = db.SQL().Model(t.model).Where("creator = ?", creator).Find(&res).Error
	return
}

func (TopicDao) FindByInviteCode(code string) (t *po.Topic, err error) {
	err = db.SQL().Where("invite_code = ?", code).Find(&t).Error
	return
}

func (TopicDao) Delete(id uint) error {
	return db.SQL().Where("tt_id = ?", id).Update("deleted_at", time.Now()).Error
}

func (TopicDao) Update(id uint, v any) error {
	return db.SQL().Updates(v).Where("tt_id = ?", id).Error
}
