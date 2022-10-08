package userimpl

import (
	"artbycode.id/go-app/services/database"
	"artbycode.id/go-app/services/user"
)

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(databaseService database.DatabaseService) *UserService {
	return &UserService{
		userRepo: NewGormUserRepository(databaseService.GetGormDB()),
	}
}

func (u *UserService) GetUser(id int) (*user.User, error) {
	user, err := u.userRepo.FindByID(id)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
