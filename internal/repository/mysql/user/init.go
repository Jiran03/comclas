package userrepository

import (
	"comclas/domain/entity"
	"comclas/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Create implements repository.UserRepository.
func (r *userRepository) Create(user *entity.User) error {
	userModel := DomainToModel(user)
	return r.db.Create(userModel).Error
}

// Delete implements repository.UserRepository.
func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

// GetAll implements repository.UserRepository.
func (r *userRepository) GetAll() ([]*entity.User, error) {
	var (
		userModels []User
		users      []*entity.User
	)
	err := r.db.Find(&userModels).Error
	if err != nil {
		return nil, err
	}

	for _, v := range userModels {
		user, err := ModelToDomain(v)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByID implements repository.UserRepository.
func (r *userRepository) GetByID(id string) (*entity.User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return ModelToDomain(user)
}

// Update implements repository.UserRepository.
func (r *userRepository) Update(user *entity.User) error {
	userModel := DomainToModel(user)
	return r.db.Save(userModel).Error
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}
