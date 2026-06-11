package user

import "github.com/gofiber/fiber/v3"

type UserHandler struct {
	service *UserService
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/:id", h.FindByID)
}

func (h *UserHandler) FindByID(c fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func newUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
