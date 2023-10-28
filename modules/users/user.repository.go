package users

import "context"

func (s *sqlStore) Create(ctx context.Context, data *UserCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) List(ctx context.Context) ([]User, error) {
	db := s.db

	var users []User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
