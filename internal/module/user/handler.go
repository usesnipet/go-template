package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mayron1806/api-template/internal/filter"
	"github.com/mayron1806/api-template/internal/model"
)

type UserHandler struct {
	service *UserService
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/:id", h.FindByID)
	router.Get("/", h.FindAll)
	router.Post("/", h.Create)
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
func (h *UserHandler) FindByID(c fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
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
func (h *UserHandler) FindAll(c fiber.Ctx) error {
	options, err := filter.FromFiber[model.User](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	users, err := h.service.FindBy(c.Context(), options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
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
func (h *UserHandler) Create(c fiber.Ctx) error {
	user := &CreateUserDTO{}
	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return h.service.Create(c.Context(), user)
}

func newUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
