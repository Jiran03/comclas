package userrepository

import "comclas/domain/entity"

func ModelToDomain(model User) (*entity.User, error) {
	user, err := entity.NewUser(entity.UserDTO{
		ID:       model.ID,
		Name:     model.Name,
		NIK:      model.NIK,
		RoleID:   model.RoleID,
		Username: model.Username,
		Password: model.Password,
		ChangeBy: model.ChangeBy,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func DomainToModel(domain *entity.User) *User {
	user := &User{
		ID:       domain.GetID(),
		Name:     domain.GetName(),
		NIK:      domain.GetNIK(),
		RoleID:   domain.GetRoleID(),
		Username: domain.GetUsername(),
		Password: domain.GetPassword(),
		ChangeBy: domain.GetChangeBy(),
	}

	return user
}
