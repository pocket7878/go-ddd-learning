package user

import (
	"github.com/google/uuid"
)

type UserId struct {
	value string
}

func NewUserId(value int) *UserId {
	uuid := uuid.New()
	return &UserId{uuid.String()}
}

func (id *UserId) Value() string {
	return id.value
}
