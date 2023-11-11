package bases

import "time"

type BaseEntity interface {
	TableName() string
}

type BaseModel struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:current_timestamp ON update current_timestamp"`
}
