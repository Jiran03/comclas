package service

import (
	"comclas/domain/repository"
	"comclas/domain/service"
	classificationservice "comclas/internal/service/classification"
	commentservice "comclas/internal/service/comment"
	roleservice "comclas/internal/service/role"
	userservice "comclas/internal/service/user"
)

func Init(repo *repository.Wrapper) *service.Wrapper {
	return &service.Wrapper{
		UserService:           userservice.NewUserService(repo),
		RoleService:           roleservice.NewRoleService(repo),
		CommentService:        commentservice.NewCommentService(repo),
		ClassificationService: classificationservice.NewClassificationService(repo),
	}
}
