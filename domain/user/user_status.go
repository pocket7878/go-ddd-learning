package user

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive            = "inactive"
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
