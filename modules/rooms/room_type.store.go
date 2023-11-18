package rooms

import (
	"Hotel_BE/modules/bases"
	"gorm.io/gorm"
)

type RoomTypeStore struct {
	db *gorm.DB
	*bases.SqlStore
}

func NewRoomTypeStore(db *gorm.DB) *RoomTypeStore {
	return &RoomTypeStore{db: db, SqlStore: bases.NewSQLStore(db, &RoomType{})}
}
