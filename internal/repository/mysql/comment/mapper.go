package commentrepository

import "comclas/domain/entity"

func ModelToDomain(model Comment) (*entity.Comment, error) {
	comment, err := entity.NewComment(entity.CommentDTO{
		ID:        model.ID,
		Value:     model.Value,
		CreatedBy: model.CreatedBy,
	})

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func DomainToModel(domain *entity.Comment) *Comment {
	comment := &Comment{
		Value:     domain.GetValue(),
		CreatedBy: domain.GetCreatedBy(),
	}

	return comment
}
