package user

import (
	"github.com/mayron1806/api-template/internal/crud"
	"github.com/mayron1806/api-template/internal/logger"
	"github.com/mayron1806/api-template/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	*crud.Repository[model.User]
}

func NewUserRepository(db *gorm.DB, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		Repository: crud.NewRepository[model.User](db, logger),
	}
}
