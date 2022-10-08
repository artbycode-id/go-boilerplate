//go:build wireinject
// +build wireinject

package runner

import (
	"artbycode.id/go-app/services/api"
	"artbycode.id/go-app/services/api/apiimpl"
	"artbycode.id/go-app/services/attendance"
	"artbycode.id/go-app/services/attendance/attendanceimpl"
	"artbycode.id/go-app/services/config"
	"artbycode.id/go-app/services/config/configimpl"
	"artbycode.id/go-app/services/database/databaseimpl"
	"artbycode.id/go-app/services/user"
	"artbycode.id/go-app/services/user/userimpl"
	"github.com/google/wire"
)

var wireSetService = wire.NewSet(
	NewRunnerServer,
	apiimpl.NewApiService,
	configimpl.NewConfigService,
	userimpl.NewUserService,
	databaseimpl.NewDatabaseService,
	attendanceimpl.NewAttendanceService,
	attendanceimpl.NewScheduleService,
	wire.Bind(new(api.ApiService), new(*apiimpl.ApiService)),
	wire.Bind(new(user.UserService), new(*userimpl.UserService)),
	wire.Bind(new(attendance.AttendanceService), new(*attendanceimpl.AttendanceService)),
	wire.Bind(new(attendance.ScheduleService), new(*attendanceimpl.ScheduleService)),
	wire.Bind(new(config.ConfigService), new(*configimpl.ConfigService)),
)

func InitializeRunnerServer() *RunnerServer {
	wire.Build(wireSetService)
	return nil
}
