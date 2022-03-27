package user

type UserStatus int

const (
	Active UserStatus = iota + 1
	Inactive
)

func (s *UserStatus) String() string {
	switch *s {
	case Active:
		return "Active"
	case Inactive:
		return "Inactive"
	default:
		panic("Unknown UserStatus")
	}
}
