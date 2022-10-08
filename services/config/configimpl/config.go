package configimpl

import (
	"os"

	"artbycode.id/go-app/services/config"
)

type ConfigService struct {
	DatabaseConfig *config.DatabaseConfig
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (c *ConfigService) InitConfig() error {
	c.DatabaseConfig = &config.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	return nil
}

func (c *ConfigService) GetConfigDatabase() *config.DatabaseConfig {
	return c.DatabaseConfig
}
