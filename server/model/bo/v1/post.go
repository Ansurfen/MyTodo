package bo

import "MyTodo/model/po/v1"

type SnapshotPost struct {
	po.SnapshotPost
	IsFavorite bool `json:"is_favorite" gorm:"column:is_favorite"`
}
