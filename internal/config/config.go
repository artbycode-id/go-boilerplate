// internal/config/config.go

package config

import (
	"github.com/spf13/viper"
)

// Config holds the application configuration
type ConfigService interface {
	Load(filename string) error
	Get() *Config
}

type configService struct {
	config *Config
}

func NewConfigService() ConfigService {
	return &configService{}
}

func (c *configService) Load(filename string) error {
	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("PORT", ":9991")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	c.config = &config
	return nil
}

func (c *configService) Get() *Config {
	return c.config
}

type Config struct {
	Port string `mapstructure:"PORT"`
}
