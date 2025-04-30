package repository

import "comclas/domain/entity"

type RoleRepository interface {
	GetAll() ([]*entity.Role, error)
	GetByID(id int) (*entity.Role, error)
	Create(role *entity.Role) error
	Update(role *entity.Role) error
	Delete(id int) error
}
