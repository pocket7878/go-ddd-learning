package user_test

import (
	"testing"

	"github.com/pocket7878/go-ddd-learning/domain/user"
)

func TestUserNameShouldBeAlphaNumeric(t *testing.T) {
	_, err := user.NewUserName("_")
	if err == nil {
		t.Fatalf("NewUserName should return error, but got nil")
	}
	if err.Error() != "user name should be alpha numeric" {
		t.Fatalf("NewUserName should return error \"user name should be alpha numeric\", but got %s", err.Error())
	}
}
