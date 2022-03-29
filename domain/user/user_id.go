package user

import (
	"github.com/google/uuid"
)

type UserID struct {
	value string
}

func NewUserID() *UserID {
	uuid := uuid.New()
	return &UserID{uuid.String()}
}

func ReconstructUserID(value string) *UserID {
	return &UserID{value}
}

func (id *UserID) Value() string {
	return id.value
}
