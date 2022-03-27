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

func ReconstructUser(id UserId, name UserName, status UserStatus) *User {
	return &User{
		id:     id,
		name:   name,
		status: status,
	}
}

func (u *User) Deactivate() {
	u.status = Inactive
}

func (u *User) UserId() *UserId {
	return &u.id
}

func (u *User) UserName() *UserName {
	return &u.name
}

func (u *User) UserStatus() *UserStatus {
	return &u.status
}
