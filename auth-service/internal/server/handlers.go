package server

import (
	"bytes"
	"encoding/json"
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

	if err = logRequest("auth", fmt.Sprintf("%s logged in", user.Email)); err != nil {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	res := &jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	return c.JSON(res)
}

func logRequest(name, data string) error {
	var payload struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	payload.Name = name
	payload.Data = data

	// convert struct into json string
	jsonData, _ := json.MarshalIndent(payload, "", "\t")

	// call the auth service
	request, err := http.NewRequest(
		http.MethodPost,
		"http://logger-service",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	return nil
}
