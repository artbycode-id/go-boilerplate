package attendanceimpl

import (
	"time"

	"artbycode.id/go-app/services/attendance"
	"artbycode.id/go-app/services/user/userimpl"
)

type AttendanceGormModel struct {
	ID               int `gorm:"primary_key"`
	StatusAttendance string
	UserID           int
	User             userimpl.UserGormModel `gorm:"foreignKey:UserID"`
	ScheduleID       int
	Schedule         ScheduleGormModel `gorm:"foreignKey:ScheduleID"`
	CheckInTime      time.Time
	CheckOutTime     time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (a AttendanceGormModel) TableName() string {
	return "attendances"
}

func (a AttendanceGormModel) ToEntity() attendance.Attendance {
	user := a.User.ToEntity()
	schedule := a.Schedule.ToEntity()
	return attendance.Attendance{
		StatusAttendance: a.StatusAttendance,
		User:             &user,
		Schedule:         &schedule,
		CheckInTime:      a.CheckInTime,
		CheckOutTime:     a.CheckOutTime,
	}
}

func (a *AttendanceGormModel) FromEntity(attendance *attendance.Attendance) {
	a.ID = attendance.ID
	a.StatusAttendance = attendance.StatusAttendance
	a.UserID = attendance.User.ID
	a.ScheduleID = attendance.Schedule.ID
	a.CheckInTime = attendance.CheckInTime
	a.CheckOutTime = attendance.CheckOutTime
}

type ScheduleGormModel struct {
	ID        uint `gorm:"primary_key"`
	Date      time.Time
	CheckIn   time.Time
	CheckOut  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s ScheduleGormModel) TableName() string {
	return "schedules"
}

func (s ScheduleGormModel) ToEntity() attendance.Schedule {
	return attendance.Schedule{
		Date:     s.Date,
		CheckIn:  s.CheckIn,
		CheckOut: s.CheckOut,
	}
}
