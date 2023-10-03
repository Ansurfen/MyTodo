package po

import sql "MyTodo/middleware/driver/sql/v1"

type Subscribe struct {
	sql.Model
	UID  uint  `gorm:"column:uid"`
	TTID uint  `gorm:"column:tt_id"`
	Perm uint8 `gorm:"column:perm"`
}

func (Subscribe) TableName() string {
	return "topic_subscribe"
}

type Follow struct {
	sql.Model
	UID      uint `gorm:"column:uid"`
	Follower uint `gorm:"column:follower"`
}

func (Follow) TableName() string {
	return "follow"
}

type Topic struct {
	sql.Model
	Creator    uint   `gorm:"creator"`
	Name       string `gorm:"name"`
	Desc       string `gorm:"desc"`
	InviteCode string `gorm:"invite_code"`
}

func (Topic) TableName() string {
	return "topic"
}
