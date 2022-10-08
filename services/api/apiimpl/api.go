package apiimpl

import (
	"artbycode.id/go-app/services/attendance"
	"artbycode.id/go-app/services/config"
	"artbycode.id/go-app/services/user"
	"github.com/gin-gonic/gin"
)

type ApiService struct {
	userService       user.UserService
	attendanceService attendance.AttendanceService
	scheduleService   attendance.ScheduleService
	configService     config.ConfigService
}

func NewApiService(userService user.UserService, attendanceService attendance.AttendanceService, scheduleService attendance.ScheduleService, configService config.ConfigService) *ApiService {
	return &ApiService{
		userService:       userService,
		attendanceService: attendanceService,
		scheduleService:   scheduleService,
		configService:     configService,
	}
}

func (a *ApiService) Run() error {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return router.Run()

}
