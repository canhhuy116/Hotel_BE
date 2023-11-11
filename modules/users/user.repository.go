package users

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
	"context"
)

func (s *UserStore) ListUser(ctx context.Context, paging *common.Paging, moreKeys ...string) ([]User, error) {
	db := s.db

	var result []User

	db = db.Table(User{}.TableName()).Find(&result)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if paging.FakeCursor != "" {
		db = db.Where("id < ?", bases.GetIdFromFakeId(paging.FakeCursor))
	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
