package user_test

import (
	"testing"

	"github.com/pocket7878/go-ddd-learning/domain/user"
)

func TestInactiveUserStatusString(t *testing.T) {
	userStatus := user.UserStatus(user.Inactive)
	inactiveString := userStatus.String()
	if inactiveString != "inactive" {
		t.Fatalf("Inactive.String() should be \"inactive\", but got %s", inactiveString)
	}
}

func TestActiveUserStatusString(t *testing.T) {
	userStatus := user.UserStatus(user.Active)
	activeString := userStatus.String()
	if activeString != "active" {
		t.Fatalf("Inactive.String() should be \"active\", but got %s", activeString)
	}
}
