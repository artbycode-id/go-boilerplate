package attendanceimpl

import (
	"artbycode.id/go-app/services/attendance"
	"artbycode.id/go-app/services/database"
	"artbycode.id/go-app/services/user"
	"artbycode.id/go-app/services/user/userimpl"
)

type AttendanceService struct {
	attendanceRepo  attendance.AttendanceRepository
	scheduleService attendance.ScheduleService
	userService     user.UserService
}

func NewAttendanceService(databaseService database.DatabaseService) *AttendanceService {
	return &AttendanceService{
		attendanceRepo:  NewGormAttendanceRepository(databaseService.GetGormDB()),
		scheduleService: NewScheduleService(databaseService),
		userService:     userimpl.NewUserService(databaseService),
	}
}

func (a *AttendanceService) SubmitAttendance(attendanceCommand *attendance.AttendanceCommand) (*attendance.ResultSubmitAttendanceDTO, error) {

	errCmd := attendanceCommand.Validate()
	if errCmd != nil {
		return nil, errCmd
	}

	schedule, err := a.scheduleService.GetScheduleToday()
	if err != nil {
		return nil, err
	}

	user, err := a.userService.GetUser(1)
	if err != nil {
		return nil, err
	}

	attendanceResult, err := a.attendanceRepo.FindAttendanceByUserIDAndScheduleID(user.ID, schedule.ID)

	if err != nil {
		return nil, err
	}

	if attendanceCommand.IsCheckIn() {
		checkIn := attendanceResult.CheckIn(schedule, user)
		if checkIn != nil {
			return nil, checkIn
		}

		err := a.attendanceRepo.SaveAttendance(attendanceResult)
		if err != nil {
			return nil, err
		}
	}

	checkOut := attendanceResult.CheckOut(schedule, user)
	if checkOut != nil {
		return nil, checkOut
	}

	err = a.attendanceRepo.UpdateAttendance(attendanceResult)
	if err != nil {
		return nil, err
	}
	return attendance.NewResultSubmitAttendanceDTO(attendanceCommand.Type, attendanceResult.StatusAttendance), nil

}
