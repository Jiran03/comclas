package rolehandler

import "comclas/domain/usecase"

var (
	InvalidRoleMsg = "invalid role"
)

type Handler struct {
	uc usecase.RoleUsecase
}

func Init(wrap *usecase.Wrapper) *Handler {
	return &Handler{
		uc: wrap.RoleUsecase,
	}
}
