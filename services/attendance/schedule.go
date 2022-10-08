package attendance

import (
	"time"
)

type Schedule struct {
	ID       int
	Date     time.Time
	CheckIn  time.Time
	CheckOut time.Time
}

type ScheduleRepository interface {
	FindScheduleByDate(date string) (*Schedule, error)
}

type ScheduleService interface {
	GetScheduleToday() (*Schedule, error)
}
