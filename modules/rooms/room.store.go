package rooms

import (
	"Hotel_BE/modules/bases"
	"gorm.io/gorm"
)

type RoomStore struct {
	db *gorm.DB
	*bases.SqlStore
}

func NewRoomStore(db *gorm.DB) *RoomStore {
	return &RoomStore{db: db, SqlStore: bases.NewSQLStore(db, &Room{})}
}
