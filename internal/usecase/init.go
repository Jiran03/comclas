package usecase

import (
	"comclas/domain/service"
	"comclas/domain/usecase"
	commentusecase "comclas/internal/usecase/comment"
	roleusecase "comclas/internal/usecase/role"
	userusecase "comclas/internal/usecase/user"
)

func Init(svc *service.Wrapper) *usecase.Wrapper {
	return &usecase.Wrapper{
		UserUsecase:    userusecase.NewUserUsecase(svc),
		RoleUsecase:    roleusecase.NewRoleUsecase(svc),
		CommentUsecase: commentusecase.NewCommentUsecase(svc),
	}
}
