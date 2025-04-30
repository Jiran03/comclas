package commentusecase

import (
	"comclas/domain/entity"
	"comclas/domain/service"
	"comclas/domain/usecase"
	"context"
)

type commentUsecase struct {
	svc      service.CommentService
	classSvc service.ClassificationService
}

// CommentClassification implements usecase.CommentUsecase.
func (u *commentUsecase) CommentClassification(ctx context.Context, comment entity.CommentDTO) (*entity.Classification, error) {
	return u.svc.CommentClassification(ctx, comment)
}

// FetchCommentByID implements usecase.CommentUsecase.
func (u *commentUsecase) FetchCommentByID(ctx context.Context, id int) (*entity.Comment, error) {
	return u.svc.FetchCommentByID(ctx, id)
}

// FetchComments implements usecase.CommentUsecase.
func (u *commentUsecase) FetchComments(ctx context.Context) ([]*entity.Comment, error) {
	return u.svc.FetchComments(ctx)
}

// RemoveComment implements usecase.CommentUsecase.
func (u *commentUsecase) RemoveComment(ctx context.Context, id int) error {
	return u.svc.RemoveComment(ctx, id)
}

// UpdateComment implements usecase.CommentUsecase.
func (u *commentUsecase) UpdateComment(ctx context.Context, id int, comment entity.CommentDTO) error {
	return u.svc.UpdateComment(ctx, id, comment)
}

func NewCommentUsecase(wrap *service.Wrapper) usecase.CommentUsecase {
	return &commentUsecase{svc: wrap.CommentService}
}
