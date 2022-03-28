package user

import (
	"testing"
)

func TestInactiveUserStatusString(t *testing.T) {
	userStatus := UserStatus(Inactive)
	inactiveString := userStatus.String()
	if inactiveString != "inactive" {
		t.Fatalf("Inactive.String() should be \"inactive\", but got %s", inactiveString)
	}
}

func TestActiveUserStatusString(t *testing.T) {
	userStatus := UserStatus(Active)
	activeString := userStatus.String()
	if activeString != "active" {
		t.Fatalf("Inactive.String() should be \"active\", but got %s", activeString)
	}
}
