package bases

import "gorm.io/gorm"

type SqlStore struct {
	db     *gorm.DB
	Entity BaseEntity
}

// NewSQLStore declare function NewSQLStore return sqlStore
func NewSQLStore(db *gorm.DB, Entity BaseEntity) *SqlStore {
	return &SqlStore{db, Entity}
}
