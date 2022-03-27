package user

import (
	"github.com/pocket7878/go-ddd-learning/domain/user"
)

type UserCreateUseCase struct {
	userRepository user.UserRepository
}

func NewUserCreateUseCase(userRepository user.UserRepository) *UserCreateUseCase {
	return &UserCreateUseCase{userRepository}
}

func (u *UserCreateUseCase) CreateUser(name string) (*user.UserId, error) {
	userName := user.NewUserName(name)
	user := user.NewUser(*userName)
	err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	userId := user.UserId()

	return &userId, nil
}
