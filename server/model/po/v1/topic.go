package po

import "time"

type SubscribedTopic struct {
	Id         uint      `json:"id" gorm:"column:id"`
	CreateAt   time.Time `json:"created_at" gorm:"column:created_at"`
	DeleteAt   time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Name       string    `json:"name" gorm:"column:name"`
	Desc       string    `json:"desc" gorm:"column:desc"`
	Creator    uint      `json:"creator" gorm:"column:creator"`
	InviteCode string    `json:"invite_code" gorm:"column:invite_code"`
}

type SubscribedMemeber struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Perm uint8  `json:"perm"`
}

const (
	PermMember = iota
	PermManager
	PermAdmin
)
