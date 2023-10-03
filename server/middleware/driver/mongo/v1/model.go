package mongo

import (
	"time"
)

type Model struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}
