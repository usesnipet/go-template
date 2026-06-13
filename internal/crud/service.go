package crud

import (
	"context"

	"github.com/mayron1806/api-template/internal/filter"
	"github.com/mayron1806/api-template/internal/logger"
	"github.com/mayron1806/api-template/internal/model"
)

type Service[T model.Model] struct {
	Repository *Repository[T]
	Logger     *logger.Logger
}

func (s *Service[T]) Create(ctx context.Context, model *T) error {
	return s.Repository.Create(ctx, model)
}

func (s *Service[T]) FindByID(ctx context.Context, id string) (T, error) {
	return s.Repository.FindByID(ctx, id)
}

func (s *Service[T]) FindBy(ctx context.Context, options *filter.Options[T]) ([]T, error) {
	s.Logger.Debugf("FindBy: %+v", options)
	return s.Repository.FindBy(ctx, options)
}

func (s *Service[T]) UpdateByID(ctx context.Context, id string, model *T) error {
	return s.Repository.UpdateByID(ctx, id, model)
}

func (s *Service[T]) DeleteByID(ctx context.Context, id string) error {
	return s.Repository.DeleteByID(ctx, id)
}

func NewService[T model.Model](repository *Repository[T], logger *logger.Logger) *Service[T] {
	return &Service[T]{Repository: repository, Logger: logger}
}
