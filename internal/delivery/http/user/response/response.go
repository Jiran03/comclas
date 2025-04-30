package userresponse

import (
	"comclas/domain/entity"
	roleresponse "comclas/internal/delivery/http/role/response"
)

type User struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	NIK      string             `json:"nik"`
	Role     *roleresponse.Role `json:"role"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	ChangeBy string             `json:"change_by"`
}

func ToUserResponse(data *entity.User) *User {
	return &User{
		ID:   data.GetID(),
		Name: data.GetName(),
		NIK:  data.GetNIK(),
		Role: &roleresponse.Role{
			ID:   1,
			Name: data.GetName(),
		},
		Username: data.GetUsername(),
		Password: data.GetPassword(),
		ChangeBy: data.GetChangeBy(),
	}
}

func ToListUserResponse(data []*entity.User) []*User {
	var userResponse []*User
	for _, val := range data {
		user := ToUserResponse(val)

		userResponse = append(userResponse, user)
	}

	return userResponse
}
