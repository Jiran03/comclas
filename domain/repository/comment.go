package repository

import "comclas/domain/entity"

type CommentRepository interface {
	GetAll() ([]*entity.Comment, error)
	GetByID(id int) (*entity.Comment, error)
	Create(comment *entity.Comment) (int, error)
	Update(comment *entity.Comment) error
	Delete(id int) error
}
