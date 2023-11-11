package bases

import (
	"github.com/sqids/sqids-go"
	"time"
)

type BaseEntity interface {
	TableName() string
}

type BaseModel struct {
	Id        int       `json:"-" gorm:"primaryKey;column:id;autoIncrement"`
	FakeID    string    `json:"id" gorm:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:current_timestamp ON update current_timestamp"`
}

func (m *BaseModel) FakeId(dbType int) {
	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})
	uid, _ := s.Encode([]uint64{uint64(m.Id), uint64(dbType)})
	m.FakeID = uid
}

func GetIdFromFakeId(fakeId string) int {
	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})
	ids := s.Decode(fakeId)
	return int(ids[0])
}
