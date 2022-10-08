package runner

import (
	"log"

	"artbycode.id/go-app/services/api"
	"artbycode.id/go-app/services/config"
	"artbycode.id/go-app/services/database"
	"github.com/joho/godotenv"
)

type RunnerServer struct {
	databaseService database.DatabaseService
	apiService      api.ApiService
	configService   config.ConfigService
}

func NewRunnerServer(databaseService database.DatabaseService, apiService api.ApiService, configService config.ConfigService) *RunnerServer {
	return &RunnerServer{
		databaseService: databaseService,
		apiService:      apiService,
		configService:   configService,
	}
}

func (r *RunnerServer) Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// init configuration
	if err := r.configService.InitConfig(); err != nil {
		return err
	}

	// init database
	if err := r.databaseService.InitConnection(); err != nil {
		return err
	}

	// auto migrate databse
	r.databaseService.Migrate()

	return r.apiService.Run()
}
