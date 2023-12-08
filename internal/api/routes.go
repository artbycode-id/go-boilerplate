package api

import "github.com/gofiber/fiber/v2"

func (a *apiService) setRoutes(app *fiber.App) {

	v1 := app.Group("/v1")

	// Auth routes group
	auth := v1.Group("/auth")
	auth.Get("/", a.authLoginHandler)
	auth.Get("/register", a.authRegisterHandler)
}
