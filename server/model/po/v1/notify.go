package po

import (
	interfaces "MyTodo/interface"
	"MyTodo/middleware/driver/sql/v1"
	"time"
)

var (
	_ interfaces.Table = (*NotifyText)(nil)
	_ interfaces.Table = (*NotifyPub)(nil)
	_ interfaces.Table = (*NotifyAction)(nil)

	_ interfaces.Po = (*NotifyText)(nil)
	_ interfaces.Po = (*NotifyPub)(nil)
	_ interfaces.Po = (*NotifyAction)(nil)
)

type NotifyText struct {
	sql.Model
	Creator int    `json:"creator" gorm:"column:uid"`
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
}

func (NotifyText) TableName() string {
	return "notify_text"
}

const (
	A_Unknown = iota
	A_NotifyAddFriend
	A_NotifyInviteFriend
	A_NotifyText
)

const (
	A_StateUnknown = iota
	A_StateWait
	A_StateConfirm
	A_StateRefuse
)

type NotifyPub struct {
	sql.Model
	NID uint `gorm:"column:nid"`
	UID uint `gorm:"column:uid"`
}

func (NotifyPub) TableName() string {
	return "notify_pub"
}

type NotifyAction struct {
	sql.Model
	Type     uint8  `gorm:"column:type"`
	Sender   uint   `gorm:"column:sender"`
	Receiver uint   `gorm:"column:receiver"`
	Status   uint8  `gorm:"column:status"`
	Param    string `gorm:"column:param"`
}

func (NotifyAction) TableName() string {
	return "notify_action"
}

type Notify struct {
	Id       int       `json:"id"`
	Type     uint8     `json:"type"`
	Status   uint8     `json:"status"`
	CreateAt time.Time `json:"created_at"`
	Param    string    `json:"param"`
}
