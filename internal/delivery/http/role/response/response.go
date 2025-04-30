package roleresponse

import (
	"comclas/domain/entity"
)

type Role struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ChangeBy string `json:"change_by"`
}

func ToRoleResponse(data *entity.Role) *Role {
	return &Role{
		ID:       data.GetID(),
		Name:     string(data.GetName()),
		ChangeBy: data.GetChangeBy(),
	}
}

func ToListRoleResponse(data []*entity.Role) []*Role {
	var roleResponse []*Role
	for _, val := range data {
		role := ToRoleResponse(val)

		roleResponse = append(roleResponse, role)
	}

	return roleResponse
}
