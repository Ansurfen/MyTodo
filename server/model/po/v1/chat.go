package po

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	// mongo.Model
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at"`
	From      uint               `json:"from" bson:"from"`
	To        uint               `json:"to" bson:"to"`
	Reply     string             `json:"reply" bson:"reply" default:""`
	Content   []string           `json:"content" bson:"content"`
}

func (Chat) TableName() string {
	return "chat"
}
