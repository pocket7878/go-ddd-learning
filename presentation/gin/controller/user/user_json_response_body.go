package user

import (
	"github.com/pocket7878/go-ddd-learning/domain/user"
)

type UserJsonResponseBody struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func NewUserJsonResponseBody(user *user.User) *UserJsonResponseBody {
	return &UserJsonResponseBody{
		ID:     user.UserID().Value(),
		Name:   user.UserName().Value(),
		Status: user.UserStatus().String(),
	}
}
