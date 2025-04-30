package rolerepository

import "time"

type Role struct {
	ID        int
	Name      string
	ChangeBy  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (m Role) GetTableName() string {
	return `roles`
}
