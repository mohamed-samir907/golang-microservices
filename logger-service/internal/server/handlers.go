package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type requestPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (s *Server) WriteLog(c *fiber.Ctx) error {
	var req requestPayload
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// insert data
	err := s.Models.LogEntry.Inset(req.Name, req.Data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	return c.JSON(resp)
}
