package service

import (
	"comclas/domain/entity"
	"context"
)

type RoleService interface {
	FetchRoles(ctx context.Context) ([]*entity.Role, error)
	FetchRoleByID(ctx context.Context, id int) (*entity.Role, error)
	AddRole(ctx context.Context, role entity.RoleDTO) error
	UpdateRole(ctx context.Context, id int, role entity.RoleDTO) error
	RemoveRole(ctx context.Context, id int) error
}
