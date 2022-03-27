package user

import (
	"context"

	"github.com/pocket7878/go-ddd-learning/domain/user"
)

type UserCreateUseCase struct {
	userRepository user.UserRepository
}

func NewUserCreateUseCase(userRepository user.UserRepository) *UserCreateUseCase {
	return &UserCreateUseCase{userRepository}
}

func (u *UserCreateUseCase) CreateUser(ctx context.Context, name string) (*user.UserId, error) {
	userName := user.NewUserName(name)
	user := user.NewUser(*userName)
	err := u.userRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	userId := user.UserId()

	return userId, nil
}
