package server

import (
	"fmt"
	"ms/mail-service/internal/mailer"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type requestPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (s *Server) SendMail(c *fiber.Ctx) error {
	var req requestPayload
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	msg := mailer.Message{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Data:    req.Message,
	}

	if err := s.Mailer.SendSMTPMessage(msg); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	resp := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("sent to %s", req.To),
	}

	return c.JSON(resp)
}
