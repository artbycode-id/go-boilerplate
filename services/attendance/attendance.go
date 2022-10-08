package attendance

import (
	"errors"
	"time"

	"artbycode.id/go-app/services/user"
)

var (
	StatusAttendanceLate          = "late"
	StatusAttendanceOnTime        = "on_time"
	StatusAttendanceCheckoutEarly = "checkout_early"
)
var (
	ErrorAttendanceNotFound = errors.New("Attendance not found")
	ErrorUserNotFound       = errors.New("User not found")
	ErrorAlreadyCheckIn     = errors.New("User already check in")
	ErrorAlreadyCheckOut    = errors.New("User already check out")

	ErrorInvalidType = errors.New("Invalid type")
)

type Attendance struct {
	ID               int
	StatusAttendance string
	Schedule         *Schedule
	User             *user.User
	CheckInTime      time.Time
	CheckOutTime     time.Time
}

func (ua Attendance) CheckIn(attendance *Schedule, user *user.User) error {

	if attendance == nil {
		return ErrorAttendanceNotFound
	}

	if user == nil {
		return ErrorUserNotFound
	}

	// check if user already check in
	if ua.CheckInTime != (time.Time{}) {
		return ErrorAlreadyCheckIn
	}

	// check if user is late
	if attendance.CheckIn.After(user.GetUserTime()) {
		ua.StatusAttendance = StatusAttendanceLate
	} else {
		ua.StatusAttendance = StatusAttendanceOnTime
	}

	ua.CheckInTime = user.GetUserTime()
	ua.User = user

	return nil
}

func (ua Attendance) CheckOut(attendance *Schedule, user *user.User) error {

	if attendance == nil {
		return ErrorAttendanceNotFound
	}

	if user == nil {
		return ErrorUserNotFound
	}

	// check if user already check out
	if ua.CheckOutTime != (time.Time{}) {
		return ErrorAlreadyCheckOut
	}

	// check if user is early checkout
	if attendance.CheckOut.Before(user.GetUserTime()) {
		ua.StatusAttendance = StatusAttendanceCheckoutEarly
	} else {
		ua.StatusAttendance = StatusAttendanceOnTime
	}

	ua.CheckOutTime = user.GetUserTime()
	ua.Schedule = attendance
	ua.User = user

	return nil
}

var (
	TypeCheckIn  = "check_in"
	TypeCheckOut = "check_out"
)

type AttendanceCommand struct {
	Type string `json:"type"`
}

func (ac AttendanceCommand) Validate() error {
	if ac.Type != TypeCheckIn && ac.Type != TypeCheckOut {
		return ErrorInvalidType
	}
	return nil
}

func (ac AttendanceCommand) IsCheckIn() bool {
	return ac.Type == TypeCheckIn
}

func (ac AttendanceCommand) IsCheckOut() bool {
	return ac.Type == TypeCheckOut
}

type ResultSubmitAttendanceDTO struct {
	TypeAttendace    string `json:"type_attendance"`
	StatusAttendance string `json:"status_attendance"`
}

func NewResultSubmitAttendanceDTO(typeAttendance string, statusAttendance string) *ResultSubmitAttendanceDTO {
	return &ResultSubmitAttendanceDTO{
		TypeAttendace:    typeAttendance,
		StatusAttendance: statusAttendance,
	}
}

type AttendanceRepository interface {
	SaveAttendance(attendance *Attendance) error
	UpdateAttendance(attendance *Attendance) error
	FindAttendanceByUserIDAndScheduleID(scheduleId, userId int) (*Attendance, error)
}

type AttendanceService interface {
	SubmitAttendance(attendanceCommand *AttendanceCommand) (*ResultSubmitAttendanceDTO, error)
}
