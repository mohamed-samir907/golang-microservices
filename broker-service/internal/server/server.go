package server

import "github.com/gofiber/fiber/v2"

func Run(port string) {
	app := fiber.New()

	handleRoutes(app)

	app.Listen(port)
}
