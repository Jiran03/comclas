package entity

type RoleName string

const (
	SuperAdminRole RoleName = "super_admin"
	AdminRole      RoleName = "admin"
	UserRole       RoleName = "user"
)

type Role struct {
	id       int
	name     RoleName
	changeby string
}

type RoleDTO struct {
	ID       int
	Name     RoleName
	ChangeBy string
}

func NewRole(dto RoleDTO) (*Role, error) {
	role := &Role{
		id:       dto.ID,
		name:     dto.Name,
		changeby: dto.ChangeBy,
	}

	return role, nil
}

func (r *Role) GetID() int {
	return r.id
}

func (r *Role) GetName() RoleName {
	return r.name
}

func (r *Role) GetChangeBy() string {
	return r.changeby
}
