package user

import "fmt"

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if !isAlphaNumeric(value) {
		return nil, fmt.Errorf("user name should be alpha numeric")
	}
	return &UserName{value}, nil
}

func ReconstructUserName(value string) *UserName {
	return &UserName{value}
}

func (name *UserName) Value() string {
	return name.value
}

func isAlphaNumeric(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}
