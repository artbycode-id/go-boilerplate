package databaseimpl

import (
	"artbycode.id/go-app/services/attendance/attendanceimpl"
	"artbycode.id/go-app/services/config"
	"artbycode.id/go-app/services/database"
	"artbycode.id/go-app/services/user/userimpl"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService struct {
	db         *gorm.DB
	cfgService config.ConfigService
}

func NewDatabaseService(configService config.ConfigService) database.DatabaseService {
	return &DatabaseService{
		cfgService: configService,
	}
}

func (d *DatabaseService) InitConnection() error {
	cfg := d.cfgService.GetConfigDatabase()
	dsn := "host=" + cfg.Host + " user=" + cfg.Username + " password=" + cfg.Password + " dbname=" + cfg.Database + " port=" + cfg.Port + " sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *DatabaseService) GetGormDB() *gorm.DB {
	return d.db
}

func (d *DatabaseService) Migrate() error {
	return d.db.AutoMigrate(
		&userimpl.UserGormModel{},
		&attendanceimpl.AttendanceGormModel{},
		&attendanceimpl.ScheduleGormModel{},
	)
}
