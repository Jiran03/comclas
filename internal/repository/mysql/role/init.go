package rolerepository

import (
	"comclas/domain/entity"
	"comclas/domain/repository"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

// Create implements repository.RoleRepository.
func (r *roleRepository) Create(role *entity.Role) error {
	roleModel := DomainToModel(role)
	return r.db.Create(roleModel).Error
}

// Delete implements repository.RoleRepository.
func (r *roleRepository) Delete(id int) error {
	return r.db.Delete(&Role{}, "id = ?", id).Error
}

// GetAll implements repository.RoleRepository.
func (r *roleRepository) GetAll() ([]*entity.Role, error) {
	var (
		roleModels []Role
		roles      []*entity.Role
	)
	err := r.db.Find(&roleModels).Error
	if err != nil {
		return nil, err
	}

	for _, v := range roleModels {
		role, err := ModelToDomain(v)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

// GetByID implements repository.RoleRepository.
func (r *roleRepository) GetByID(id int) (*entity.Role, error) {
	var role Role
	err := r.db.First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return ModelToDomain(role)
}

// Update implements repository.RoleRepository.
func (r *roleRepository) Update(role *entity.Role) error {
	roleModel := DomainToModel(role)
	return r.db.Save(roleModel).Error
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db}
}
