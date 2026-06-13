package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mayron1806/api-template/internal/filter"
)

type UserHandler struct {
	service *UserService
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/:id", h.FindByID)
	router.Get("/", h.FindAll)
	router.Post("/", h.Create)
}

func (h *UserHandler) FindByID(c fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *UserHandler) FindAll(c fiber.Ctx) error {
	options, err := filter.FromFiber[User](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	users, err := h.service.FindBy(c.Context(), options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

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
