package attendanceimpl

import (
	"artbycode.id/go-app/services/attendance"
	"gorm.io/gorm"
)

type GormScheduleRepository struct {
	db *gorm.DB
}

func NewGormScheduleRepository(db *gorm.DB) *GormScheduleRepository {
	return &GormScheduleRepository{db: db}
}

func (g *GormScheduleRepository) FindScheduleByDate(date string) (*attendance.Schedule, error) {
	var scheduleGorm ScheduleGormModel
	err := g.db.Where("date = ?", date).First(&scheduleGorm).Error
	if err != nil {
		return &attendance.Schedule{}, err
	}
	entity := scheduleGorm.ToEntity()
	return &entity, nil
}
