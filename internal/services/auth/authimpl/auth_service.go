package authimpl

import (
	"artbycode.id/go-app/internal/services/auth"
)

type AuthServiceImpl struct {
	userRepository auth.UserRepository
}

func NewAuthServiceImpl(userRepository auth.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (a *AuthServiceImpl) Login(email, password string) (*auth.AuthLoginDTO, error) {
	if password != "password" {
		return nil, auth.NewAuthError(auth.ErrInvalidPassword, auth.CodeErrorInvalidPassword)
	}
	user, err := a.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, auth.NewAuthError(auth.ErrUserEmailNotFound, auth.CodeErrorInvalidPassword)
	}

	return &auth.AuthLoginDTO{
		User:  *user,
		Token: "token",
	}, nil
}
