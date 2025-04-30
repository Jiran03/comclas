package userhandler

import "comclas/domain/usecase"

var (
	InvalidRoleMsg = "invalid role"
)

type Handler struct {
	uc usecase.UserUsecase
}

func Init(wrap *usecase.Wrapper) *Handler {
	return &Handler{
		uc: wrap.UserUsecase,
	}
}
