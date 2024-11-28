package main

import (
	"medods-test/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	logger := logger.NewLogger()

	app.Get("/user/ip_warning", func(c fiber.Ctx) error {
		// imagine that there is a call to Google smtp
		logger.Info("Mail sended!")
		return c.SendStatus(fiber.StatusOK)
	})

	panic(app.Listen(":6666"))
}
