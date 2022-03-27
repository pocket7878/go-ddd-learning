package user

import (
	"github.com/pocket7878/go-ddd-learning/domain/user"
)

type UserDeactivateUseCase struct {
	userRepository user.UserRepository
}

func NewUserDeactivateUseCase(userRepository user.UserRepository) *UserDeactivateUseCase {
	return &UserDeactivateUseCase{userRepository}
}

func (u *UserDeactivateUseCase) DeactivateUser(userId user.UserId) error {
	user, err := u.userRepository.FindById(userId)
	if err != nil {
		return err
	}

	user.Deactivate()

	return u.userRepository.Save(user)
}
