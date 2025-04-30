package commentservice

import (
	"comclas/domain/entity"
	"comclas/domain/repository"
	"comclas/domain/service"
	"context"
)

type commentService struct {
	repo     repository.CommentRepository
	classSvc service.ClassificationService
}

// CommentClassification implements service.CommentService.
func (s *commentService) CommentClassification(ctx context.Context, comment entity.CommentDTO) (*entity.Classification, error) {
	domain, err := entity.NewComment(comment)
	if err != nil {
		return nil, err
	}

	commentID, errCreateComment := s.repo.Create(domain)
	if errCreateComment != nil {
		return nil, errCreateComment
	}

	classification, errAddClass := s.classSvc.AddClassification(ctx, entity.ClassificationDTO{
		CommentID: commentID,
	})
	if errAddClass != nil {
		return nil, errAddClass
	}

	classification.SetCommment(domain)

	return classification, nil
}

// FetchCommentByID implements service.CommentService.
func (s *commentService) FetchCommentByID(ctx context.Context, id int) (*entity.Comment, error) {
	comment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// FetchComments implements service.CommentService.
func (s *commentService) FetchComments(ctx context.Context) ([]*entity.Comment, error) {
	comments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// RemoveComment implements service.CommentService.
func (s *commentService) RemoveComment(ctx context.Context, id int) error {
	return s.repo.Delete(id)
}

// UpdateComment implements service.CommentService.
func (s *commentService) UpdateComment(ctx context.Context, id int, comment entity.CommentDTO) error {
	comment.ID = id
	domain, err := entity.NewComment(comment)
	if err != nil {
		return err
	}

	return s.repo.Update(domain)
}

func NewCommentService(wrap *repository.Wrapper) service.CommentService {
	return &commentService{
		repo: wrap.CommentRepository,
	}
}
