package roleservice

import (
	"comclas/domain/entity"
	"comclas/domain/repository"
	"comclas/domain/service"
	"context"
)

type roleService struct {
	repo repository.RoleRepository
}

// AddRole implements service.RoleService.
func (r *roleService) AddRole(ctx context.Context, role entity.RoleDTO) error {
	domain, err := entity.NewRole(role)
	if err != nil {
		return err
	}

	return r.repo.Create(domain)
}

// FetchRoleByID implements service.RoleService.
func (r *roleService) FetchRoleByID(ctx context.Context, id int) (*entity.Role, error) {
	return r.repo.GetByID(id)
}

// FetchRoles implements service.RoleService.
func (r *roleService) FetchRoles(ctx context.Context) ([]*entity.Role, error) {
	return r.repo.GetAll()
}

// RemoveRole implements service.RoleService.
func (r *roleService) RemoveRole(ctx context.Context, id int) error {
	return r.repo.Delete(id)
}

// UpdateRole implements service.RoleService.
func (r *roleService) UpdateRole(ctx context.Context, id int, role entity.RoleDTO) error {
	role.ID = id
	domain, err := entity.NewRole(role)
	if err != nil {
		return err
	}

	return r.repo.Update(domain)
}

func NewRoleService(wrap *repository.Wrapper) service.RoleService {
	return &roleService{repo: wrap.RoleRepository}
}
