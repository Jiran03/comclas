package userservice

import (
	"comclas/domain/entity"
	"comclas/domain/repository"
	"comclas/domain/service"
	uuidpkg "comclas/pkg/uuid"
	"context"
)

type userService struct {
	repo repository.UserRepository
}

// AddUser implements service.UserService.
func (u *userService) AddUser(ctx context.Context, user entity.UserDTO) error {
	id := uuidpkg.GenerateID()
	user.ID = id
	if user.ChangeBy == "" {
		user.ChangeBy = user.ID
	}

	domain, err := entity.NewUser(user)
	if err != nil {
		return err
	}

	return u.repo.Create(domain)
}

// FetchUserByID implements service.UserService.
func (u *userService) FetchUserByID(ctx context.Context, id string) (*entity.User, error) {
	return u.repo.GetByID(id)
}

// FetchUsers implements service.UserService.
func (u *userService) FetchUsers(ctx context.Context) ([]*entity.User, error) {
	return u.repo.GetAll()
}

// RemoveUser implements service.UserService.
func (u *userService) RemoveUser(ctx context.Context, id string) error {
	return u.repo.Delete(id)
}

// UpdateUser implements service.UserService.
func (u *userService) UpdateUser(ctx context.Context, id string, user entity.UserDTO) error {
	user.ID = id
	domain, err := entity.NewUser(user)
	if err != nil {
		return err
	}

	return u.repo.Update(domain)
}

func NewUserService(wrap *repository.Wrapper) service.UserService {
	return &userService{repo: wrap.UserRepository}
}
