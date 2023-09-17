package server

import (
	"ms/logger-service/internal/data"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Models data.Models
}

func (s *Server) Run(port string) {
	app := fiber.New()

	s.handleRoutes(app)

	app.Listen(port)
}
