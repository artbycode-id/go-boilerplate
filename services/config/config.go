package config

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type ConfigService interface {
	GetConfigDatabase() *DatabaseConfig
	InitConfig() error
}
