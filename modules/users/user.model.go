package users

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
)

const EntityName = "User"

type User struct {
	bases.BaseModel
	Username string `json:"username" gorm:"column:username;unique"`
	Password string `json:"password" gorm:"column:password"`
	Name     string `json:"name" gorm:"column:name"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask() {
	u.FakeId(common.DbUser)
}
