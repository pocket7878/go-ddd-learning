package user

import (
	userUseCase "github.com/pocket7878/go-ddd-learning/usecase/user"
)

type UserController struct {
	userCreateUseCase     userUseCase.UserCreateUseCase
	userDeactivateUseCase userUseCase.UserDeactivateUseCase
}

func NewUserController(userCreateUseCase userUseCase.UserCreateUseCase, userDeactivateUseCase userUseCase.UserDeactivateUseCase) *UserController {
	return &UserController{
		userCreateUseCase:     userCreateUseCase,
		userDeactivateUseCase: userDeactivateUseCase,
	}
}
