package user

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive            = "inactive"
)

func (s *UserStatus) String() string {
	switch *s {
	case Active:
		return "active"
	case Inactive:
		return "inactive"
	default:
		panic("Unknown UserStatus")
	}
}
