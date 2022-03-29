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

func (u *UserDeactivateUseCase) DeactivateUser(ctx context.Context, userID user.UserID) (*user.User, error) {
	user, err := u.userRepository.FindByID(ctx, userID)
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
