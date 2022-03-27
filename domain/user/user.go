package user

type User struct {
	id     UserId
	name   UserName
	status UserStatus
}

func NewUser(id UserId, name UserName) *User {
	return &User{id, name, Inactive}
}

func (u *User) Deactivate() {
	u.status = Inactive
}
