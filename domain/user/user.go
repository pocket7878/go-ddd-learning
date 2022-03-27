package user

type User struct {
	id     UserId
	name   UserName
	status UserStatus
}

func NewUser(name UserName) *User {
	return &User{
		id:     *NewUserId(),
		name:   name,
		status: Inactive,
	}
}

func (u *User) Deactivate() {
	u.status = Inactive
}

func (u *User) UserId() UserId {
	return u.id
}
