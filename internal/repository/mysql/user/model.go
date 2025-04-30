package userrepository

import "time"

type User struct {
	ID        string
	Name      string
	NIK       string
	RoleID    int
	Username  string
	Password  string
	ChangeBy  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (m User) GetTableName() string {
	return `users`
}
