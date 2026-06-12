package user

import (
	"context"

	"github.com/mayron1806/api-template/internal/module/database"
)

type UserService struct {
	db *database.DB
}

func (s *UserService) Create(ctx context.Context, user *User) error {
	return nil
}

func (s *UserService) FindByID(ctx context.Context, id string) (*User, error) {
	user := &User{}
	return user, nil
}

func newUserService(db *database.DB) *UserService {
	return &UserService{
		db: db,
	}
}
