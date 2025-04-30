package service

import (
	"comclas/domain/entity"
	"context"
)

type ClassificationService interface {
	FetchClassifications(ctx context.Context) ([]*entity.Classification, error)
	FetchClassificationByID(ctx context.Context, id int) (*entity.Classification, error)
	AddClassification(ctx context.Context, classification entity.ClassificationDTO) (*entity.Classification, error)
	UpdateClassification(ctx context.Context, id int, classification entity.ClassificationDTO) error
	RemoveClassification(ctx context.Context, id int) error
}
