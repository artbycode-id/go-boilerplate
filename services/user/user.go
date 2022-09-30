package user

type User struct {
	ID   int
	Name string
}

type CreateUserCommand struct {
	Name string
}

type UserService interface {
	SearchUser(id int) (*User, error)
	CreateUserAdmin(userCmd *CreateUserCommand) error
}

type UserRepository interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) error
}
