package sql

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m *Model) Exist() bool {
	return m.ID != 0
}

type NullString = sql.NullString
