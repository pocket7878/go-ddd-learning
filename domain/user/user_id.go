package user

import (
	"github.com/google/uuid"
)

type UserId struct {
	value string
}

func NewUserId() *UserId {
	uuid := uuid.New()
	return &UserId{uuid.String()}
}

func ReconstructUserId(value string) *UserId {
	return &UserId{value}
}

func (id *UserId) Value() string {
	return id.value
}
