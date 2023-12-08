package auth

type AuthLoginDTO struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type AuthService interface {
	Login(email, password string) (*AuthLoginDTO, error)
}
