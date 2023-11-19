package rooms

import (
	"Hotel_BE/modules/bases"
	"context"
)

type RoomRepository interface {
	bases.BaseRepository
}

type RoomBiz struct {
	repo RoomRepository
}

func NewRoomBiz(repo RoomRepository) *RoomBiz {
	return &RoomBiz{repo: repo}
}

func (biz *RoomBiz) CreateRoom(ctx context.Context, data *RoomCreate) error {
	err := biz.repo.Create(ctx, data)

	return err
}
