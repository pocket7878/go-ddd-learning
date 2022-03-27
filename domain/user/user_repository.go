package user

import "context"

type UserRepository interface {
	FindById(ctx context.Context, id UserId) (*User, error)
	Save(ctx context.Context, user *User) error
}
