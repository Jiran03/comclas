package httpdelivery

import (
	"comclas/domain/usecase"
	commentdelivery "comclas/internal/delivery/http/comment"
	roledelivery "comclas/internal/delivery/http/role"
	userdelivery "comclas/internal/delivery/http/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, uc *usecase.Wrapper) {
	userdelivery.Routes(router, uc)
	roledelivery.Routes(router, uc)
	commentdelivery.Routes(router, uc)
}
