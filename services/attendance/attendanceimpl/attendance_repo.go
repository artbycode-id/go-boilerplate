package attendanceimpl

import (
	"artbycode.id/go-app/services/attendance"
	"gorm.io/gorm"
)

type GormAttendanceRepository struct {
	db *gorm.DB
}

func NewGormAttendanceRepository(db *gorm.DB) *GormAttendanceRepository {
	return &GormAttendanceRepository{db: db}
}

func (g *GormAttendanceRepository) SaveAttendance(attendance *attendance.Attendance) error {
	attendanceGorm := AttendanceGormModel{}
	attendanceGorm.FromEntity(attendance)
	return g.db.Create(&attendanceGorm).Error
}

func (g *GormAttendanceRepository) FindAttendanceByUserIDAndScheduleID(userID int, scheduleID int) (*attendance.Attendance, error) {
	var attendanceGorm AttendanceGormModel
	err := g.db.Where("user_id = ? AND schedule_id = ?", userID, scheduleID).First(&attendanceGorm).Error
	if err != nil {
		return &attendance.Attendance{}, err
	}
	entity := attendanceGorm.ToEntity()
	return &entity, nil
}

func (g *GormAttendanceRepository) UpdateAttendance(attendance *attendance.Attendance) error {
	attendanceGorm := AttendanceGormModel{}
	attendanceGorm.FromEntity(attendance)
	return g.db.Save(&attendanceGorm).Error
}
