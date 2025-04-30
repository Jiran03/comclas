package classificationrepository

import (
	"comclas/domain/entity"
	"comclas/domain/repository"

	"gorm.io/gorm"
)

type classificationRepository struct {
	db *gorm.DB
}

// Create implements repository.ClassificationRepository.
func (r *classificationRepository) Create(classification *entity.Classification) error {
	classificationModel := DomainToModel(classification)
	return r.db.Create(classificationModel).Error
}

// Delete implements repository.ClassificationRepository.
func (r *classificationRepository) Delete(id int) error {
	return r.db.Delete(&Classification{}, "id = ?", id).Error
}

// GetAll implements repository.ClassificationRepository.
func (r *classificationRepository) GetAll() ([]*entity.Classification, error) {
	var (
		classificationModels []Classification
		classifications      []*entity.Classification
	)
	err := r.db.Find(&classificationModels).Error
	if err != nil {
		return nil, err
	}

	for _, v := range classificationModels {
		classification, err := ModelToDomain(v)
		if err != nil {
			return nil, err
		}

		classifications = append(classifications, classification)
	}

	return classifications, nil
}

// GetByID implements repository.ClassificationRepository.
func (r *classificationRepository) GetByID(id int) (*entity.Classification, error) {
	var classification Classification
	err := r.db.First(&classification, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return ModelToDomain(classification)
}

// Update implements repository.ClassificationRepository.
func (r *classificationRepository) Update(classification *entity.Classification) error {
	classificationModel := DomainToModel(classification)
	return r.db.Save(classificationModel).Error
}

func NewClassificationRepository(db *gorm.DB) repository.ClassificationRepository {
	return &classificationRepository{db}
}
