package userusecase

import (
	"comclas/domain/entity"
	"comclas/domain/service"
	"comclas/domain/usecase"
	"context"
)

type userUsecase struct {
	svc service.UserService
}

// AddUser implements usecase.UserUsecase.
func (u *userUsecase) AddUser(ctx context.Context, user entity.UserDTO) error {
	return u.svc.AddUser(ctx, user)
}

// FetchUserByID implements usecase.UserUsecase.
func (u *userUsecase) FetchUserByID(ctx context.Context, id string) (*entity.User, error) {
	return u.svc.FetchUserByID(ctx, id)
}

// FetchUsers implements usecase.UserUsecase.
func (u *userUsecase) FetchUsers(ctx context.Context) ([]*entity.User, error) {
	return u.svc.FetchUsers(ctx)
}

// RemoveUser implements usecase.UserUsecase.
func (u *userUsecase) RemoveUser(ctx context.Context, id string) error {
	return u.svc.RemoveUser(ctx, id)
}

// UpdateUser implements usecase.UserUsecase.
func (u *userUsecase) UpdateUser(ctx context.Context, id string, user entity.UserDTO) error {
	return u.svc.UpdateUser(ctx, id, user)
}

func NewUserUsecase(wrap *service.Wrapper) usecase.UserUsecase {
	return &userUsecase{svc: wrap.UserService}
}
