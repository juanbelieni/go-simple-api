package app

import (
	"github.com/gofiber/fiber"
)

func GetApp() *fiber.App {
	app := fiber.New()
	app.Post("/signin", Signin)
	app.Post("/login", Login)

	return app
}
