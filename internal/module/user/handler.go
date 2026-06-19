package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/usesnipet/go-template/internal/api"
	"github.com/usesnipet/go-template/internal/crud"
	"github.com/usesnipet/go-template/internal/logger"
	"github.com/usesnipet/go-template/internal/model"
	"github.com/usesnipet/go-template/internal/module/app"
)

type UserHandler struct {
	*crud.Handler[model.User]
	service *UserService
	logger  *logger.Logger
}

func (h *UserHandler) RegisterRoutes(r chi.Router, serve func(app.HandlerFunc) http.HandlerFunc) {
	r.Get("/{id}", serve(h.FindByID))
	r.Get("/", serve(h.FindAll))
	r.Post("/", serve(h.Create))
	r.Put("/{id}", serve(h.UpdateByID))
	r.Delete("/{id}", serve(h.DeleteByID))
}

// FindByID godoc
//
//	@Summary		Find user by ID
//	@Description Return a user by ID.
//	@Tags			users
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	model.User
//	@Failure		500	{object}	api.ErrorResponse
//	@Router			/users/{id} [GET]
func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) error {
	return h.Handler.FindByID(w, r)
}

// FindAll godoc
//
//	@Summary		Find all users
//	@Description	Return a list of users with pagination and optional filters.
//	@Tags			users
//	@Produce		json
//	@Param			take	query		int	false	"Maximum number of records"	default(2000)
//	@Param			skip	query		int	false	"Number of records to skip"	default(0)
//	@Success		200		{array}		model.User
//	@Failure		400		{object}	api.ErrorResponse
//	@Failure		500		{object}	api.ErrorResponse
//	@Router			/users [GET]
func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) error {
	return h.Handler.FindBy(w, r)
}

// Create godoc
//
//	@Summary		Create user
//	@Description	Create a new user.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		CreateUserDTO	true	"User data"
//	@Success		200
//	@Failure		400	{object}	api.ErrorResponse
//	@Failure		500	{object}	api.ErrorResponse
//	@Router			/users [POST]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) error {
	h.logger.Verbosef("%s %s Create", r.Method, r.URL.Path)
	dto := &CreateUserDTO{}
	if err := api.DecodeJSON(r, dto); err != nil {
		return err
	}
	return h.service.Create(r.Context(), dto)
}

// UpdateByID godoc
//
//	@Summary		Update user by ID
//	@Description	Update a user by ID.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Param			user	body		CreateUserDTO	true	"User data"
//	@Success		200
//	@Failure		400	{object}	api.ErrorResponse
//	@Failure		500	{object}	api.ErrorResponse
//	@Router			/users/{id} [PUT]
func (h *UserHandler) UpdateByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	h.logger.Verbosef("%s %s UpdateByID: %s", r.Method, r.URL.Path, id)
	dto := &CreateUserDTO{}
	if err := api.DecodeJSON(r, dto); err != nil {
		return err
	}
	return h.service.UpdateByID(r.Context(), id, dto)
}

// DeleteByID godoc
//
//	@Summary		Delete user by ID
//	@Description	Delete a user by ID.
//	@Tags			users
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200
//	@Failure		400	{object}	api.ErrorResponse
//	@Failure		500	{object}	api.ErrorResponse
//	@Router			/users/{id} [DELETE]
func (h *UserHandler) DeleteByID(w http.ResponseWriter, r *http.Request) error {
	return h.Handler.DeleteByID(w, r)
}

func NewUserHandler(service *UserService, logger *logger.Logger) *UserHandler {
	return &UserHandler{
		Handler: crud.NewHandler(service.Service, logger),
		service: service,
		logger:  logger,
	}
}
