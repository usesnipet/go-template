package crud

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/usesnipet/go-template/internal/api"
	"github.com/usesnipet/go-template/internal/filter"
	"github.com/usesnipet/go-template/internal/logger"
	"github.com/usesnipet/go-template/internal/model"
)

type Handler[T model.Model] struct {
	service *Service[T]
	logger  *logger.Logger
}

func (h *Handler[T]) FindByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	h.logger.Verbosef("%s %s FindByID: %s", r.Method, r.URL.Path, id)
	model, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		return err
	}
	return api.WriteJSON(w, http.StatusOK, model)
}

func (h *Handler[T]) FindBy(w http.ResponseWriter, r *http.Request) error {
	h.logger.Verbosef("%s %s FindBy", r.Method, r.URL.Path)
	options, err := filter.FromRequest[T](r)
	if err != nil {
		return err
	}
	models, err := h.service.FindBy(r.Context(), options)
	if err != nil {
		return err
	}
	return api.WriteJSON(w, http.StatusOK, models)
}

func (h *Handler[T]) DeleteByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	h.logger.Verbosef("%s %s DeleteByID: %s", r.Method, r.URL.Path, id)
	return h.service.DeleteByID(r.Context(), id)
}

func NewHandler[T model.Model](service *Service[T], logger *logger.Logger) *Handler[T] {
	return &Handler[T]{service: service, logger: logger}
}
