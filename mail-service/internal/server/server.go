package server

import (
	"ms/mail-service/internal/mailer"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Mailer *mailer.Mail
}

func New() *Server {
	return &Server{
		Mailer: mailer.NewMail(),
	}
}

func (s *Server) Run(port string) {
	app := fiber.New()

	s.handleRoutes(app)

	app.Listen(port)
}
