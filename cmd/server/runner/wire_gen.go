// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitializeRunnerServer() *RunnerServer {
	configService := configimpl.NewConfigService()
	databaseService := databaseimpl.NewDatabaseService(configService)
	userService := userimpl.NewUserService(databaseService)
	attendanceService := attendanceimpl.NewAttendanceService(databaseService)
	scheduleService := attendanceimpl.NewScheduleService(databaseService)
	apiService := apiimpl.NewApiService(userService, attendanceService, scheduleService, configService)
	runnerServer := NewRunnerServer(databaseService, apiService, configService)
	return runnerServer
}

// wire.go:

var wireSetService = wire.NewSet(
	NewRunnerServer, apiimpl.NewApiService, configimpl.NewConfigService, userimpl.NewUserService, databaseimpl.NewDatabaseService, attendanceimpl.NewAttendanceService, attendanceimpl.NewScheduleService, wire.Bind(new(api.ApiService), new(*apiimpl.ApiService)), wire.Bind(new(user.UserService), new(*userimpl.UserService)), wire.Bind(new(attendance.AttendanceService), new(*attendanceimpl.AttendanceService)), wire.Bind(new(attendance.ScheduleService), new(*attendanceimpl.ScheduleService)), wire.Bind(new(config.ConfigService), new(*configimpl.ConfigService)),
)