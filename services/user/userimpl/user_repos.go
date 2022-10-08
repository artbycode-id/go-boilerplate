package userimpl

import (
	"artbycode.id/go-app/services/user"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (g *GormUserRepository) FindByID(id int) (user.User, error) {
	var userGorm UserGormModel
	err := g.db.First(&userGorm, id).Error
	if err != nil {
		return user.User{}, err
	}
	return userGorm.ToEntity(), nil
}
