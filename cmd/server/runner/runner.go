package runner

import (
	"os"

	"artbycode.id/go-app/internal/api"
	"artbycode.id/go-app/internal/config"
)

type RunnerServer struct {
	apiService    api.ApiService
	configService config.ConfigService
}

func NewRunnerServer(apiService api.ApiService, configService config.ConfigService) *RunnerServer {
	return &RunnerServer{
		apiService:    apiService,
		configService: configService,
	}
}

func (r *RunnerServer) Run() error {
	var cfgFile = func() string {
		if customFile := os.Getenv("APP_CONFIG_FILE"); customFile != "" {
			return customFile
		}
		return "./config/app.yml"
	}()

	if err := r.configService.Load(cfgFile); err != nil {
		return err
	}
	return r.apiService.Serve()
}
