package user

type User struct {
	id     UserID
	name   UserName
	status UserStatus
}

func NewUser(name UserName) *User {
	return &User{
		id:     *NewUserID(),
		name:   name,
		status: Inactive,
	}
}

func ReconstructUser(id UserID, name UserName, status UserStatus) *User {
	return &User{
		id:     id,
		name:   name,
		status: status,
	}
}

func (u *User) Deactivate() {
	u.status = Inactive
}

func (u *User) UserID() *UserID {
	return &u.id
}

func (u *User) UserName() *UserName {
	return &u.name
}

func (u *User) UserStatus() *UserStatus {
	return &u.status
}
