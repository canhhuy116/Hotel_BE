package users

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Username  string    `json:"username" gorm:"column:username;unique"`
	Password  string    `json:"password" gorm:"column:password"`
	Name      string    `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:current_timestamp ON update current_timestamp"`
}

func (User) TableName() string {
	return "users"
}
