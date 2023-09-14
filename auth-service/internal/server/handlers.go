package server

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (s *Server) Auth(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Check if the user exists
	user, err := s.Models.User.GetByEmail(req.Email)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("invalid credentials")
	}

	// validate the password
	valid, err := user.PasswordMatches(req.Password)
	if err != nil || !valid {
		return c.Status(http.StatusUnauthorized).SendString("invalid credentials")
	}

	res := &jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	return c.JSON(res)
}
