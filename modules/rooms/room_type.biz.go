package rooms

import (
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
