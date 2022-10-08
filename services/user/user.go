package user

import "time"

type User struct {
	ID   int
	Name string
}

func (u User) GetUserTime() time.Time {
	return time.Now()
}

type UserRepository interface {
	FindByID(id int) (User, error)
}

type UserService interface {
	GetUser(id int) (*User, error)
}
