package attendanceimpl

import (
	"time"

	"artbycode.id/go-app/services/attendance"
	"artbycode.id/go-app/services/database"
)

type ScheduleService struct {
	repoSchedule attendance.ScheduleRepository
}

func NewScheduleService(databaseService database.DatabaseService) *ScheduleService {
	return &ScheduleService{
		repoSchedule: NewGormScheduleRepository(databaseService.GetGormDB()),
	}
}

func (s *ScheduleService) GetScheduleToday() (*attendance.Schedule, error) {
	toDayFormat := time.Now().Format("2006-01-02")
	schedule, err := s.repoSchedule.FindScheduleByDate(toDayFormat)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
