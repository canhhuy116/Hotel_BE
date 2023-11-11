package bases

import (
	"Hotel_BE/common"
	"context"
)

func (s *SqlStore) Create(ctx context.Context, data interface{}) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *SqlStore) Find(ctx context.Context, id int, entity interface{}) error {
	if err := s.db.Table(s.Entity.TableName()).First(entity, id).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *SqlStore) FindAll(ctx context.Context, entities interface{}) error {
	if err := s.db.Table(s.Entity.TableName()).Find(entities).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *SqlStore) Update(ctx context.Context, id int, data interface{}) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *SqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(s.Entity.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
