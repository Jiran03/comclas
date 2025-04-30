package classificationrepository

import "comclas/domain/entity"

func ModelToDomain(model Classification) (*entity.Classification, error) {
	classification, err := entity.NewClassification(entity.ClassificationDTO{
		ID:                 model.ID,
		CommentID:          model.CommentID,
		ClassificationType: entity.ClassificationType(model.Type),
		CreatedBy:          model.CreatedBy,
	})

	if err != nil {
		return nil, err
	}

	return classification, nil
}

func DomainToModel(domain *entity.Classification) *Classification {
	classification := &Classification{
		ID:        domain.GetID(),
		CommentID: domain.GetCommentID(),
		Type:      domain.GetType(),
		CreatedBy: domain.GetCreatedBy(),
	}

	return classification
}
