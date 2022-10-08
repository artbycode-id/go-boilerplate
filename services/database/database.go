package database

import "gorm.io/gorm"

type DatabaseService interface {
	GetGormDB() *gorm.DB
	Migrate() error
	InitConnection() error
}
