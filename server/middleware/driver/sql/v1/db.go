package sql

import (
	"MyTodo/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoDB struct {
	*gorm.DB
}

func New(opt conf.SQLOption) *TodoDB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &TodoDB{db}
}

func (db *TodoDB) RawExec(sql string, a ...any) (tx *gorm.DB) {
	return db.Raw(fmt.Sprintf(sql, a...))
}
