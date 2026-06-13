package user

import (
	"context"

	"github.com/mayron1806/api-template/internal/crud"
	"github.com/mayron1806/api-template/internal/logger"
)

type UserService struct {
	*crud.Service[User]
}

func (s *UserService) Create(ctx context.Context, model *CreateUserDTO) error {
	return s.Service.Create(
		ctx,
		model.ToModel(),
	)
}

func NewUserService(repository *UserRepository, logger *logger.Logger) *UserService {
	return &UserService{
		Service: crud.NewService(repository.Repository, logger),
	}
}
