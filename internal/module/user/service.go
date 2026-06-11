package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/mayron1806/api-template/internal/module/database"
)

type UserService struct {
	db *database.DB
}

func (s *UserService) FindByID(ctx context.Context, id string) (*User, error) {
	user := &User{}
	err := s.db.Run(
		s.db.Builder.Select().From("users").Where(squirrel.Eq{"id": id}).ToSql,
		user,
		ctx,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func newUserService(db *database.DB) *UserService {
	return &UserService{
		db: db,
	}
}
