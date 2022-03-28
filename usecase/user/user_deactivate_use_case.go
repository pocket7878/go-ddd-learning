package user

import (
	"context"

	"github.com/pocket7878/go-ddd-learning/domain/user"
)

type UserDeactivateUseCase struct {
	userRepository user.UserRepository
}

func NewUserDeactivateUseCase(userRepository user.UserRepository) *UserDeactivateUseCase {
	return &UserDeactivateUseCase{userRepository}
}

func (u *UserDeactivateUseCase) DeactivateUser(ctx context.Context, userId user.UserId) (*user.User, error) {
	user, err := u.userRepository.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	user.Deactivate()
	err = u.userRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
