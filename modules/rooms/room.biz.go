package rooms

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
	"context"
)

type RoomRepository interface {
	bases.BaseRepository
	ListRoom(ctx context.Context, paging *common.Paging, moreKeys ...string) ([]Room, error)
	FindRoom(ctx context.Context, id int, moreKeys ...string) (*Room, error)
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

func (biz *RoomBiz) UpdateRoom(ctx context.Context, id int, data *RoomUpdate) error {
	var oldData Room

	if err := biz.repo.Find(ctx, id, &oldData); err != nil {
		return common.ErrCannotGetEntity(RoomEntityName, err)
	}

	if err := biz.repo.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(RoomEntityName, err)
	}

	return nil
}

func (biz *RoomBiz) GetRooms(ctx context.Context, paging *common.Paging) ([]Room, error) {
	var rooms []Room

	rooms, err := biz.repo.ListRoom(ctx, paging, "RoomType")

	if err != nil {
		return nil, common.ErrCannotGetEntity(RoomEntityName, err)
	}

	return rooms, nil
}

func (biz *RoomBiz) GetRoom(ctx context.Context, id int) (*Room, error) {
	result, err := biz.repo.FindRoom(ctx, id, "RoomType")

	if err != nil {
		return nil, common.ErrCannotGetEntity(RoomEntityName, err)
	}

	return result, nil
}
