package po

import (
	sql "MyTodo/middleware/driver/sql/v1"
	"database/sql/driver"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	sql.Model
	UID     uint   `json:"uid" gorm:"column:uid"`
	Content string `json:"content" gorm:"column:content"`
	Image   Images `json:"image" gorm:"column:image;type:json"`
}

func (Post) TableName() string {
	return "post"
}

type Image struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Images []Image

func (t *Images) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Images) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type PostComment struct {
	sql.Model
	PID     uint   `gorm:"column:pid"`
	UID     uint   `gorm:"column:uid"`
	Content string `gorm:"column:content"`
	Image   Images `gorm:"column:image;type:json"`
}

func (PostComment) TableName() string {
	return "post_comment"
}

type PostCommentFavorite struct {
	sql.Model
	CID string `gorm:"uniqueIndex:unique_cid_uid;column:cid;type:varchar(50)"`
	UID uint   `gorm:"uniqueIndex:unique_cid_uid;column:uid"`
}

func (PostCommentFavorite) TableName() string {
	return "post_comment_favorite"
}

type PostCommentV2 struct {
	// mongo.Model
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at"`
	UID       uint               `json:"uid" bson:"uid"`
	PID       uint               `json:"pid" bson:"pid"`
	Reply     uint               `json:"reply" bson:"reply"`
	Content   []string           `json:"content" bson:"content"`
	Replies   []PostCommentV2    `json:"replies" bson:"replies,omitempty"`
}

func (PostCommentV2) TableName() string {
	return "post_comment"
}

type PostFavorite struct {
	sql.Model
	PID uint `gorm:"uniqueIndex:unique_pid_uid;column:pid"`
	UID uint `gorm:"uniqueIndex:unique_pid_uid;column:uid"`
}

func (PostFavorite) TableName() string {
	return "post_favorite"
}

type SnapshotPost struct {
	Id          uint      `json:"id"`
	Uid         uint      `json:"uid"`
	Username    string    `json:"username"`
	IsMale      bool      `json:"isMale"`
	CreateAt    time.Time `json:"created_at" gorm:"column:created_at"`
	Content     string    `json:"content"`
	Image       Images    `json:"image"`
	FavoriteCnt int       `json:"fc" gorm:"column:favoriteCnt"`
	CommentCnt  int       `json:"cc" gorm:"column:commentCnt"`
}

type DetailedPostComment struct {
	PostCommentV2
	Username    string `json:"username"`
	ReplyName   string `json:"reply_name"`
	Favorite    int64  `json:"favorite"`
	YouFavorite bool   `json:"you_favorite"`
}

type DetailedPostCommentFavorite struct {
	CID   string `json:"cid" gorm:"column:cid"`
	Count int64  `json:"count" gorm:"column:count"`
}

type DetailedPost struct {
	Id         uint   `json:"id" gorm:"column:id"`
	Username   string `json:"username" gorm:"column:username"`
	Favorite   int64  `json:"favorite" gorm:"column:favorite_count"`
	IsFavorite bool   `json:"is_favorite" gorm:"column:is_favorite"`
	UID        uint   `json:"uid" gorm:"column:uid"`
	Content    string `json:"content" gorm:"column:content"`
	IsMale     bool   `json:"isMale" gorm:"column:is_male"`
	Images     Images `json:"images" gorm:"column:image"`
}
