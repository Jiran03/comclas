package repository

import "comclas/domain/entity"

type ClassificationRepository interface {
	GetAll() ([]*entity.Classification, error)
	GetByID(id int) (*entity.Classification, error)
	Create(classification *entity.Classification) error
	Update(classification *entity.Classification) error
	Delete(id int) error
}
