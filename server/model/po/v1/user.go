package po

import "MyTodo/middleware/driver/sql/v1"

type User struct {
	sql.Model
	IsMale    bool           `json:"isMale" gorm:"column:is_male"`
	Email     string         `json:"email" gorm:"column:email;type:varchar(50);unique;"`
	Telephone sql.NullString `json:"telephone" gorm:"column:telephone;type:varchar(11)"`
	Name      string         `json:"name" gorm:"column:name;type:varchar(25);"`
	Password  string         `json:"password" gorm:"column:password;type:text;"`
}

func (User) TableName() string {
	return "user"
}
