package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/user/ip_warning", func(c fiber.Ctx) error {
		// zdarova ti voshel s novogo ip
		// imagine that there is a call to Google smtp
		return c.SendStatus(fiber.StatusOK)
	})

	panic(app.Listen(":6666"))
}
