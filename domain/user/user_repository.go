package user

type UserRepository interface {
	FindById(id UserId) (*User, error)
	Save(user *User) error
}
