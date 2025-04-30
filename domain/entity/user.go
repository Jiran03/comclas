package entity

type User struct {
	id       string
	name     string
	nik      string
	roleID   int
	username string
	password string
	changeby string
}

type UserDTO struct {
	ID       string
	Name     string
	NIK      string
	RoleID   int
	Username string
	Password string
	ChangeBy string
}

func NewUser(dto UserDTO) (*User, error) {
	user := &User{
		id:       dto.ID,
		name:     dto.Name,
		nik:      dto.NIK,
		roleID:   dto.RoleID,
		username: dto.Username,
		password: dto.Password,
		changeby: dto.ChangeBy,
	}

	return user, nil
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetNIK() string {
	return u.nik
}
func (u *User) GetRoleID() int {
	return u.roleID
}
func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetChangeBy() string {
	return u.changeby
}
