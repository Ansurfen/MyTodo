package dao

import (
	"MyTodo/engine/v1/cli"
	"MyTodo/engine/v1/db"
	interfaces "MyTodo/interface"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type Pagination[T any] struct {
	Index int64 `json:"index"`
	Size  int64 `json:"size"`
	Data  []T   `json:"data"`
}

func (page *Pagination[T]) Offest(offest ...int) int64 {
	skip := 10
	if len(offest) > 0 {
		skip = offest[0]
	}
	if skip <= 0 {
		skip = 10
	}
	return (page.Index - 1) * int64(skip)
}

func (page *Pagination[T]) SQLClause(offest ...int) *gorm.DB {
	return db.SQL().Offset(int(page.Offest(offest...))).Limit(int(page.Size))
}

func (page *Pagination[T]) MongoFindOption(offest ...int) *options.FindOptions {
	return options.Find().
		SetLimit(page.Size).
		SetSkip(page.Offest(offest...))
}

type SQLBaseDao[T any] struct{}

func (dao *SQLBaseDao[T]) Create(v interfaces.Po) error {
	return db.SQL().Create(v).Error
}

func (dao *SQLBaseDao[T]) FindAll(v any) error {
	var m T
	return db.SQL().Model(m).Find(v).Error
}

func (dao *SQLBaseDao[T]) UnscopedFindAll(v any) error {
	var m T
	return db.SQL().Model(m).Unscoped().Find(v).Error
}

func (dao *SQLBaseDao[T]) FindOneByID(id int, v any) error {
	var m T
	return db.SQL().Model(m).Where("id = ?", id).First(v).Error
}

func (dap *SQLBaseDao[T]) UnscopedFindOneByID(id int, v any) error {
	var m T
	return db.SQL().Model(m).Unscoped().Where("id = ?", id).First(v).Error
}

func (dao *SQLBaseDao[T]) FindManyByID(id int, v ...any) error {
	var m T
	return db.SQL().Model(m).Where("id = ?", id).Find(v).Error
}

func (dao *SQLBaseDao[T]) UnscopedFindManyByID(id int, v ...any) error {
	var m T
	return db.SQL().Model(m).Unscoped().Where("id = ?", id).Find(v).Error
}

func (dao *SQLBaseDao[T]) DeleteByID(id int) error {
	var m T
	return db.SQL().Model(m).Where("id = ?", id).Update("deleted_at = ?", time.Now()).Error
}

func (dao *SQLBaseDao[T]) UnscopedDeleteByID(id int) error {
	var m T
	return db.SQL().Model(m).Unscoped().Where("id = ?", id).Update("deleted_at = ?", time.Now()).Error
}

func (dao *SQLBaseDao[T]) UpdateByID(id int, column string, v any) error {
	var m T
	return db.SQL().Model(m).Where("id = ?", id).Update(column, v).Error
}

func (dao *SQLBaseDao[T]) UpdatesByID(id int, v any) error {
	var m T
	return db.SQL().Model(m).Where("id = ?", id).Updates(v).Error
}

type MongoBaseDao[T interfaces.Table] struct {
	col *mongo.Collection
}

func (dao *MongoBaseDao[T]) Collection() *mongo.Collection {
	if dao.col == nil {
		var v T
		dao.col = db.Mongo().Collection(cli.Option.Mongo.Database, v)
	}
	return dao.col
}
