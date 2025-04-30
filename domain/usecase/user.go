package usecase

import (
	"comclas/domain/entity"
	"context"
)

type UserUsecase interface {
	FetchUsers(ctx context.Context) ([]*entity.User, error)
	FetchUserByID(ctx context.Context, id string) (*entity.User, error)
	AddUser(ctx context.Context, user entity.UserDTO) error
	UpdateUser(ctx context.Context, id string, user entity.UserDTO) error
	RemoveUser(ctx context.Context, id string) error
}
