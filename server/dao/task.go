package dao

import (
	"MyTodo/engine/v1/db"
	"MyTodo/model/po/v1"
	"time"
)

type TaskCondDao struct {
	po.TaskCond
}

func (tc *TaskCondDao) FindAll() (res []po.TaskCond, err error) {
	err = db.SQL().Model(tc).Find(&res).Error
	return
}

func (tc *TaskCondDao) Create(cond po.TaskCond) error {
	return db.SQL().Create(&cond).Error
}

func (tc *TaskCondDao) Delete(id uint) error {
	return db.SQL().Model(tc).Update("deletead_at", time.Now()).Where("id = ?", id).Error
}

type TaskDao struct {
	po.Task
}

func (t *TaskDao) Create(task *po.Task) error {
	return db.SQL().Create(task).Error
}

func (t *TaskDao) Delete(name string) error {
	return db.SQL().Update("deleted_at", time.Now()).Where("name = ?", name).Error
}

func (t *TaskDao) Update(name string, v any) error {
	return db.SQL().Updates(v).Where("name = ?", name).Error
}

const PageSize = 10

func (t *TaskDao) Find(page, limit int) (ret []po.Task, err error) {
	err = db.SQL().Model(TaskDao{}).Offset((page - 1) * PageSize).Limit(limit).Order("created_at desc").Find(&ret).Error
	return
}

func (t *TaskDao) FindById(id int) (ret TaskDao, err error) {
	ret.ID = uint(id)
	err = db.SQL().Find(&ret).Error
	return
}

type TaskCommitDao struct {
	po.TaskCommit
}

func (t *TaskCommitDao) Create(commit po.TaskCommit) error {
	return db.SQL().Create(&commit).Error
}

type TaskBoundDao struct {
	po.TaskBound
}

func (t *TaskBoundDao) Create(bound po.TaskBound) error {
	return db.SQL().Create(&bound).Error
}

func (t *TaskBoundDao) FindOne(tid, _type int) (ret po.TaskBound, err error) {
	err = db.SQL().Model(t).Find(&ret).Where(po.TaskBound{
		TID:  uint(tid),
		TCID: uint(_type),
	}).Error
	return
}

func (t *TaskBoundDao) FindMany(query po.TaskBound) (ret []po.TaskBound, err error) {
	err = db.SQL().Model(t).Find(&ret).Where(query).Error
	return
}

func (t *TaskBoundDao) Update() {}
