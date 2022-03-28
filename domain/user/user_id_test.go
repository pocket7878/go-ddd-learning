package user

import (
	"testing"
)

func TestNewUserIdGenerateUniqueId(t *testing.T) {
	userIdFirst := NewUserId()
	userIdSecond := NewUserId()
	if userIdFirst.Value() == userIdSecond.Value() {
		t.Fatalf("NewUserId genrate same id twice")
	}
}
