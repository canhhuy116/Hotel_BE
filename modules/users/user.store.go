package users

import (
	"Hotel_BE/modules/bases"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
	*bases.SqlStore
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db, bases.NewSQLStore(db, &User{})}
}
