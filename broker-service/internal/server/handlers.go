package server

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type requestPayload struct {
	Action string      `json:"action"`
	Auth   authPayload `json:"auth,omitempty"`
}

type authPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Broker(c *fiber.Ctx) error {
	payload := &jsonResponse{
		Error:   false,
		Message: "Done",
	}

	return c.JSON(payload)
}

func HandleSubmission(c *fiber.Ctx) error {
	var reqPayload requestPayload
	err := c.ParamsParser(&reqPayload)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	switch reqPayload.Action {
	case "auth":
		return authenticate(c, &reqPayload.Auth)

	default:
		return c.Status(http.StatusBadRequest).SendString("unknown action")
	}
}

func authenticate(c *fiber.Ctx, a *authPayload) error {
	// convert struct into json string
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the auth service
	request, err := http.NewRequest(
		http.MethodPost,
		"http://auth-service/auth",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		return c.Status(http.StatusUnauthorized).SendString("invalid credentials")
	} else if response.StatusCode != http.StatusOK {
		return c.Status(http.StatusBadRequest).SendString("error calling auth service")
	}

	var jsonRes jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonRes)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if jsonRes.Error {
		return c.Status(http.StatusUnauthorized).SendString(jsonRes.Message)
	}

	payload := &jsonResponse{
		Error:   false,
		Message: "Authenticated",
		Data:    jsonRes.Data,
	}

	return c.JSON(payload)
}
