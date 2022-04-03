package user_test

import (
	"testing"

	"github.com/pocket7878/go-ddd-learning/domain/user"
)

func TestNewUserCreateInactiveUser(t *testing.T) {
	userName, err := user.NewUserName("test")
	if err != nil {
		t.Fatalf("NewUserName should not return error, but got %s", err.Error())
	}
	u := user.NewUser(*userName)
	if *u.UserStatus() != user.Inactive {
		t.Fatalf("NewUser should create inactive user, but got %s", u.UserStatus())
	}
}

func TestDeactivateUser(t *testing.T) {
	userName, err := user.NewUserName("test")
	if err != nil {
		t.Fatalf("NewUserName should not return error, but got %s", err.Error())
	}
	u := user.NewUser(*userName)
	u.Deactivate()
	if *u.UserStatus() != user.Inactive {
		t.Fatalf("NewUser should create inactive user, but got %s", u.UserStatus())
	}
}
