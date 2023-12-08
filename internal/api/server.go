package api

import (
	"artbycode.id/go-app/internal/config"
	"artbycode.id/go-app/internal/services/auth"
	"github.com/gofiber/fiber/v2"
)

type ApiService interface {
	Serve() error
}

type apiService struct {
	configService config.ConfigService
	authService   auth.AuthService
}

func NewApiService(
	configService config.ConfigService,
	authSapiService auth.AuthService,
) ApiService {
	return &apiService{
		configService: configService,
		authService:   authSapiService,
	}
}

func (a *apiService) Serve() error {
	app := fiber.New()
	a.setRoutes(app)
	return app.Listen(":" + a.configService.Get().Port)
}
