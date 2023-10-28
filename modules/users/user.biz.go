package users

import "context"

type UserRepository interface {
	Create(ctx context.Context, data *UserCreate) error
	List(ctx context.Context) ([]User, error)
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
	users, err := biz.repo.List(ctx)

	return users, err
}
