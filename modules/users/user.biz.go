package users

import (
	"Hotel_BE/modules/bases"
	"context"
)

type UserRepository interface {
	bases.BaseRepository
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

func (biz *UserBiz) ListUser(ctx context.Context) ([]User, error) {
	var users []User
	err := biz.repo.FindAll(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
