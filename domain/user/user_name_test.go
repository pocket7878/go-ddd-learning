package user

import (
	"testing"
)

func TestUserNameShouldBeAlphaNumeric(t *testing.T) {
	_, err := NewUserName("_")
	if err == nil {
		t.Fatalf("NewUserName should return error, but got nil")
	}
	if err.Error() != "user name should be alpha numeric" {
		t.Fatalf("NewUserName should return error \"user name should be alpha numeric\", but got %s", err.Error())
	}
}
