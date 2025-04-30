package repository

import "comclas/domain/entity"

type UserRepository interface {
	GetAll() ([]*entity.User, error)
	GetByID(id string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id string) error
}
