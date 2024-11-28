package auth

import (
	auth_models "medods-test/internal/auth/models"
	"medods-test/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

type handlers struct {
	usecase Usecase
	logger  *logger.Logger
}

func NewHandlers(usecase Usecase, logger *logger.Logger) *handlers {
	return &handlers{
		usecase: usecase,
		logger:  logger,
	}
}

func (h *handlers) MapRoutes(router fiber.Router) {
	router.Get("/signin", h.Signin())
	router.Get("/refresh", h.Refresh())
}

func (h *handlers) Signin() fiber.Handler {
	return func(c fiber.Ctx) error {
		ip := c.Get("X-Forwarded-For", "")
		guid := c.Query("GUID", "")
		if ip == "" {
			return c.Status(fiber.StatusNonAuthoritativeInformation).SendString("Invalid ip")
		}

		if guid == "" {
			return c.Status(fiber.StatusNonAuthoritativeInformation).SendString("Invalid GUID")
		}

		resp, err := h.usecase.Signin(&auth_models.SigninParams{
			Guid:   guid,
			UserIp: ip,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(resp)
	}
}

func (h *handlers) Refresh() fiber.Handler {
	return func(c fiber.Ctx) error {
		return nil
	}
}
