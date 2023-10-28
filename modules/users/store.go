package users

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

// NewSQLStore declare function NewSQLStore return sqlStore
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db}
}
