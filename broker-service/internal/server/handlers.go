package server

import (
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
	Log    logPayload  `json:"log,omitempty"`
	Mail   mailPayload `json:"mail,omitempty"`
}

type authPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type logPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type mailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
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
	if err := c.BodyParser(&reqPayload); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	switch reqPayload.Action {
	case "auth":
		return authenticate(c, &reqPayload.Auth)

	case "log":
		return logItem(c, &reqPayload.Log)

	case "mail":
		return sendMail(c, &reqPayload.Mail)

	default:
		return c.Status(http.StatusBadRequest).SendString("unknown action")
	}
}

func sendMail(c *fiber.Ctx, m *mailPayload) error {
	response, err := SendPostRequest("http://mail-service", m)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusOK {
		return c.Status(http.StatusBadRequest).SendString("error calling mail service")
	}

	jsonRes, err := ParseReponseBody(response)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(jsonRes)
}

func logItem(c *fiber.Ctx, l *logPayload) error {
	response, err := SendPostRequest("http://logger-service", l)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusOK {
		return c.Status(http.StatusBadRequest).SendString("error calling logger service")
	}

	jsonRes, err := ParseReponseBody(response)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(jsonRes)
}

func authenticate(c *fiber.Ctx, a *authPayload) error {
	response, err := SendPostRequest("http://auth-service/auth", a)
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

	jsonRes, err := ParseReponseBody(response)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	payload := &jsonResponse{
		Error:   false,
		Message: "Authenticated",
		Data:    jsonRes.Data,
	}

	return c.JSON(payload)
}
