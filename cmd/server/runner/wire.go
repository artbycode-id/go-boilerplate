//go:build wireinject
// +build wireinject

package runner

import (
	"artbycode.id/go-app/internal/api"
	"artbycode.id/go-app/internal/config"

	"artbycode.id/go-app/internal/services/auth"
	"artbycode.id/go-app/internal/services/auth/authimpl"

	"github.com/google/wire"
)

var wireSetService = wire.NewSet(
	NewRunnerServer,
	api.NewApiService,
	config.NewConfigService,
	authimpl.NewUserRepositoryImpl,
	wire.Bind(new(auth.UserRepository), new(*authimpl.UserRepositoryImpl)),
	authimpl.NewAuthServiceImpl,
	wire.Bind(new(auth.AuthService), new(*authimpl.AuthServiceImpl)),
)

func InitializeRunnerServer() *RunnerServer {
	wire.Build(wireSetService)
	return nil
}
