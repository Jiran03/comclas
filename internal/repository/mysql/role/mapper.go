package rolerepository

import "comclas/domain/entity"

func ModelToDomain(model Role) (*entity.Role, error) {
	role, err := entity.NewRole(entity.RoleDTO{
		ID:       model.ID,
		Name:     entity.RoleName(model.Name),
		ChangeBy: model.ChangeBy,
	})

	if err != nil {
		return nil, err
	}

	return role, nil
}

func DomainToModel(domain *entity.Role) *Role {
	role := &Role{
		ID:       domain.GetID(),
		Name:     string(domain.GetName()),
		ChangeBy: domain.GetChangeBy(),
	}

	return role
}
