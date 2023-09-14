package server

import (
	"database/sql"
	"ms/auth-service/internal/data"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	DB     *sql.DB
	Models data.Models
}

func (s *Server) Run(port string) {
	app := fiber.New()

	s.handleRoutes(app)

	app.Listen(port)
}
