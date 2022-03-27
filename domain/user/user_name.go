package user

type UserName struct {
	value string
}

func NewUserName(value string) *UserName {
	// TODO: Validate domain logic
	return &UserName{value}
}

func ReconstructUserName(value string) *UserName {
	return &UserName{value}
}

func (name *UserName) Value() string {
	return name.value
}
