package user

import "context"

type UserRepository interface {
	FindByID(ctx context.Context, id UserID) (*User, error)
	Save(ctx context.Context, user *User) error
}
