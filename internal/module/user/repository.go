package user

import (
	"github.com/mayron1806/api-template/internal/crud"
	"github.com/mayron1806/api-template/internal/logger"
	"gorm.io/gorm"
)

type UserRepository struct {
	*crud.Repository[User]
}

func NewUserRepository(db *gorm.DB, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		Repository: crud.NewRepository[User](db, logger),
	}
}
