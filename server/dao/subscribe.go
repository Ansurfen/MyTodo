package dao

import (
	"MyTodo/engine/v1/db"
	"MyTodo/model/po/v1"
)

type Subscribe struct {
	model po.Subscribe
}

func (Subscribe) Create(s *po.Subscribe) error {
	return db.SQL().Create(s).Error
}

func (s *Subscribe) Delete(uid, ttid uint) error {
	return db.SQL().Model(s.model).Delete("uid = ? and tt_id = ?", uid, ttid).Error
}

func (s *Subscribe) FindByID(ttid uint) (res []po.Subscribe, err error) {
	err = db.SQL().Model(s.model).Where("tt_id = ?", ttid).Find(&res).Error
	return
}

func (s *Subscribe) FindOne(ttid, uid int) (res po.Subscribe, err error) {
	err = db.SQL().Model(s.model).Where("tt_id = ? AND uid = ?", ttid, uid).First(&res).Error
	return
}

func (s *Subscribe) FindHistory() {

}

func (s *Subscribe) UpdatePerm(uid, ttid int, perm uint8) error {
	return db.SQL().Model(s.model).Where("uid = ? and ttid = ?", uid, ttid).Update("perm", perm).Error
}

type Follow struct {
	po.Follow
}

func (Follow) Create(s *po.Follow) error {
	return db.SQL().Create(s).Error
}

func (s *Follow) Delete(uid, follower uint) error {
	return db.SQL().Model(s).Delete("uid = ? and follower = ?", uid, follower).Error
}

type Topic struct {
	po.Topic
}

func (Topic) Create(t *po.Topic) error {
	return db.SQL().Create(t).Error
}
