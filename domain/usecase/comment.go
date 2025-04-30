package usecase

import (
	"comclas/domain/entity"
	"context"
)

type CommentUsecase interface {
	FetchComments(ctx context.Context) ([]*entity.Comment, error)
	FetchCommentByID(ctx context.Context, id int) (*entity.Comment, error)
	CommentClassification(ctx context.Context, comment entity.CommentDTO) (*entity.Classification, error)
	UpdateComment(ctx context.Context, id int, comment entity.CommentDTO) error
	RemoveComment(ctx context.Context, id int) error
}
