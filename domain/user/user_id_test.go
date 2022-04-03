package user_test

import (
	"testing"

	"github.com/pocket7878/go-ddd-learning/domain/user"
)

func TestNewUserIDGenerateUniqueID(t *testing.T) {
	userIDFirst := user.NewUserID()
	userIDSecond := user.NewUserID()
	if userIDFirst.Value() == userIDSecond.Value() {
		t.Fatalf("NewUserID genrate same id twice")
	}
}
