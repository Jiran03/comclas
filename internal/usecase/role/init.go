package roleusecase

import (
	"comclas/domain/entity"
	"comclas/domain/service"
	"comclas/domain/usecase"
	"context"
)

type roleUsecase struct {
	svc service.RoleService
}

// AddRole implements usecase.RoleUsecase.
func (r *roleUsecase) AddRole(ctx context.Context, role entity.RoleDTO) error {
	return r.svc.AddRole(ctx, role)
}

// FetchRoleByID implements usecase.RoleUsecase.
func (r *roleUsecase) FetchRoleByID(ctx context.Context, id int) (*entity.Role, error) {
	return r.svc.FetchRoleByID(ctx, id)
}

// FetchRoles implements usecase.RoleUsecase.
func (r *roleUsecase) FetchRoles(ctx context.Context) ([]*entity.Role, error) {
	return r.svc.FetchRoles(ctx)
}

// RemoveRole implements usecase.RoleUsecase.
func (r *roleUsecase) RemoveRole(ctx context.Context, id int) error {
	return r.svc.RemoveRole(ctx, id)
}

// UpdateRole implements usecase.RoleUsecase.
func (r *roleUsecase) UpdateRole(ctx context.Context, id int, role entity.RoleDTO) error {
	return r.svc.UpdateRole(ctx, id, role)
}

func NewRoleUsecase(wrap *service.Wrapper) usecase.RoleUsecase {
	return &roleUsecase{svc: wrap.RoleService}
}
