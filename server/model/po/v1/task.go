package po

import (
	sql "MyTodo/middleware/driver/sql/v1"
	"time"
)

type TaskTopic struct {
	sql.Model
	Creator uint
	Name    string
	Desc    string
}

func (TaskTopic) TableName() string {
	return "task_topic"
}

type TopicJoin struct {
	sql.Model
	UID  uint
	TTID uint
}

func (TopicJoin) TableName() string {
	return "topic_join"
}

type Task struct {
	sql.Model
	TTID      uint      `json:"tt_id" gorm:"column:tt_id"`
	User      uint      `json:"user" gorm:"column:user;"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(50);"`
	Desc      string    `json:"desc" gorm:"column:desc;type:text;"`
	Departure time.Time `json:"departure" gorm:"column:departure;"`
	Arrival   time.Time `json:"arrival" gorm:"column:arrival;"`
}

func (Task) TableName() string {
	return "task"
}

const (
	COND_HAND = iota
	COND_TIMER
	COND_LOCATE
	COND_FILE
	COND_IMAGE
	COND_CONTENT
	COND_QR
)

type TaskCond struct {
	sql.Model
	Name string `json:"name"`
	Desc string `json:"desc"`
	TCID uint   `json:"tcid" gorm:"column:tc_id"`
}

func (TaskCond) TableName() string {
	return "task_cond"
}

type ITaskCond interface {
	Auth(p string) bool
}

type TaskCondQR struct {
	Param string
}

func (c *TaskCondQR) Auth(p string) bool {
	return true
}

type TaskCondHand struct{}

func (c *TaskCondHand) Auth(p string) bool {
	return true
}

type TaskCondLocate struct {
	Param string
}

func (c *TaskCondLocate) Auth(p string) bool {
	return true
}

type TaskBound struct {
	sql.Model
	TID   uint   `json:"tid"`
	TCID  uint   `json:"tcid"`
	Param string `json:"param"`
}

func (TaskBound) TableName() string {
	return "task_bound"
}

type TaskCommit struct {
	sql.Model
	UID   uint   `json:"uid" gorm:"column:uid;"`
	TID   uint   `json:"tid"`
	TCID  uint   `json:"tcid"`
	Param string `json:"param"`
}

func (TaskCommit) TableName() string {
	return "task_commit"
}

type DetailedTask struct {
	Id uint `json:"id" gorm:"column:task_id"`
	// Creator   uint      `json:"creator" gorm:"column:user"`
	Topic     string    `json:"topic" gorm:"column:topic_name"`
	Name      string    `json:"name" gorm:"column:task_name"`
	Desc      string    `json:"desc" gorm:"column:task_desc"`
	Departure time.Time `json:"departure" gorm:"column:departure"`
	Arrival   time.Time `json:"arrival" gorm:"column:arrival"`
	CondType  uint      `json:"-" gorm:"column:tc_id"`
	Conds     []uint    `json:"conds" gorm:"-"`
}
