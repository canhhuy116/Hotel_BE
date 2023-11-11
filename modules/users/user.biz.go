package users

import (
	"Hotel_BE/common"
	"Hotel_BE/modules/bases"
	"context"
)

type UserRepository interface {
	bases.BaseRepository
	ListUser(ctx context.Context, paging *common.Paging, moreKeys ...string) ([]User, error)
}

type UserBiz struct {
	repo UserRepository
}

func NewUserBiz(repo UserRepository) *UserBiz {
	return &UserBiz{repo: repo}
}

func (biz *UserBiz) CreateUser(ctx context.Context, data *UserCreate) error {
	err := biz.repo.Create(ctx, data)

	return err
}

func (biz *UserBiz) ListUser(ctx context.Context, paging *common.Paging) ([]User,
	error) {
	//var users []User
	//err := biz.repo.FindAll(ctx, &users)
	users, err := biz.repo.ListUser(ctx, paging)

	return users, err
}
