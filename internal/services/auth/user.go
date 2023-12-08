package auth

type User struct {
	ID       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
}
