package authimpl

import "artbycode.id/go-app/internal/services/auth"

type UserRepositoryImpl struct {
	users []auth.User
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: []auth.User{
			{
				ID:       1,
				UUID:     "2d3dw1",
				Email:    "test@email.com",
				Name:     "name",
				Password: "password",
			},
		},
	}
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*auth.User, error) {
	for _, user := range u.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}
