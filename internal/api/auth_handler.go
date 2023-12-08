package api

import (
	"artbycode.id/go-app/internal/services/auth"
	"github.com/gofiber/fiber/v2"
)

func (a *apiService) authLoginHandler(c *fiber.Ctx) error {

	// get data from params
	param := c.Query("pw")

	r, err := a.authService.Login("test@email.com", param)
	if err != nil {
		if authErr, ok := err.(*auth.AuthErrorDTO); ok {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(ResponseErrorBody{
				Code:     authErr.Code,
				Message:  authErr.Message,
				Detail:   authErr.Detail,
				UniqueId: authErr.UniqueId,
			})
		}
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(errorInternalServerResponseBody(err))
	}
	return c.JSON(r)
}

func (a *apiService) authRegisterHandler(c *fiber.Ctx) error {
	return c.SendString("Register")
}
