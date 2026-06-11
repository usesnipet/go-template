package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/mayron1806/api-template/internal/model"
	"github.com/mayron1806/api-template/internal/module/database"
)

type UserService struct {
	db *database.DB
}

func (s *UserService) FindById(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	query, args, err := s.db.Builder.Select().From("users").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}
	err = s.db.GetContext(ctx, user, query, args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(db *database.DB) *UserService {
	return &UserService{
		db: db,
	}
}
