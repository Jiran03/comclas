package classificationservice

import (
	"comclas/domain/entity"
	"comclas/domain/repository"
	"comclas/domain/service"
	"context"
)

type classificationService struct {
	repo repository.ClassificationRepository
}

// AddClassification implements service.ClassificationService.
func (s *classificationService) AddClassification(ctx context.Context, classification entity.ClassificationDTO) (*entity.Classification, error) {
	domain, err := entity.NewClassification(classification)
	if err != nil {
		return nil, err
	}

	errCreateClassification := s.repo.Create(domain)
	if errCreateClassification != nil {
		return nil, errCreateClassification
	}

	return domain, nil
}

// FetchClassificationByID implements service.ClassificationService.
func (s *classificationService) FetchClassificationByID(ctx context.Context, id int) (*entity.Classification, error) {
	classification, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return classification, nil
}

// FetchClassifications implements service.ClassificationService.
func (s *classificationService) FetchClassifications(ctx context.Context) ([]*entity.Classification, error) {
	classifications, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return classifications, nil
}

// RemoveClassification implements service.ClassificationService.
func (s *classificationService) RemoveClassification(ctx context.Context, id int) error {
	return s.repo.Delete(id)
}

// UpdateClassification implements service.ClassificationService.
func (s *classificationService) UpdateClassification(ctx context.Context, id int, classification entity.ClassificationDTO) error {
	classification.ID = id
	domain, err := entity.NewClassification(classification)
	if err != nil {
		return err
	}

	return s.repo.Update(domain)
}

func NewClassificationService(wrap *repository.Wrapper) service.ClassificationService {
	return &classificationService{
		repo: wrap.ClassificationRepository,
	}
}
