package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
	"context"
)

type RoomTypeRepository interface {
	bases.BaseRepository
}

type RoomTypeBiz struct {
	repo RoomTypeRepository
}

func NewRoomTypeBiz(repo RoomTypeRepository) *RoomTypeBiz {
	return &RoomTypeBiz{repo: repo}
}

func (biz *RoomTypeBiz) CreateRoomType(ctx context.Context, data *RoomTypeCreate) error {
	err := biz.repo.Create(ctx, data)

	return err
}

func (biz *RoomTypeBiz) UpdateRoomType(ctx context.Context, id int, data *RoomTypeUpdate) error {
	var oldData RoomType
	err := biz.repo.Find(ctx, id, &oldData)

	if err != nil {
		return common.ErrCannotGetEntity(EntityName, err)
	}

	if err := biz.repo.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(EntityName, err)
	}

	return nil
}

func (biz *RoomTypeBiz) GetRoomTypes(ctx context.Context) ([]RoomType, error) {
	var result []RoomType
	err := biz.repo.FindAll(ctx, &result)

	if err != nil {
		return nil, common.ErrCannotListEntity(EntityName, err)
	}

	return result, nil
}
