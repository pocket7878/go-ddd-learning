package user

import (
	"testing"
)

func TestNewUserIDGenerateUniqueID(t *testing.T) {
	userIDFirst := NewUserID()
	userIDSecond := NewUserID()
	if userIDFirst.Value() == userIDSecond.Value() {
		t.Fatalf("NewUserID genrate same id twice")
	}
}
