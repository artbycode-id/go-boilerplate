package userimpl

import (
	"time"

	"artbycode.id/go-app/services/user"
)

type UserGormModel struct {
	ID        int `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u UserGormModel) TableName() string {
	return "users"
}

func (u UserGormModel) ToEntity() user.User {
	return user.User{
		ID:   u.ID,
		Name: u.Name,
	}
}
