package server

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *Server) handleRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://*, http://*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization, Accept, X-CSRF-Token",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	app.Use("/ping", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString(time.Now().Format(time.DateTime))
	})

	app.Post("/", s.SendMail)
}
