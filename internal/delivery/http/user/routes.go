package userdelivery

import (
	"comclas/domain/usecase"
	userhandler "comclas/internal/delivery/http/user/handler"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, uc *usecase.Wrapper) {
	var (
		prefix       = "/v1/user"
		handler      = userhandler.Init(uc)
		pathListAuth = map[string]struct {
			path    string
			handler func(*gin.Context)
			method  string
		}{
			"createUser": {
				path:    "/create",
				handler: handler.CreateUser,
				method:  http.MethodPost,
			},
			"listUser": {
				path:    "",
				handler: handler.FetchUsers,
				method:  http.MethodGet,
			},
			"detailUser": {
				path:    "/:id",
				handler: handler.FetchUserByID,
				method:  http.MethodGet,
			},
		}

		pathList = map[string]struct {
			path    string
			handler func(*gin.Context)
			method  string
		}{
			"loginUser": {
				path:    "/login/:id",
				handler: handler.LoginUser,
				method:  http.MethodPost,
			},
		}

		routesAuth = router.Group(prefix)
		routes     = router.Group(prefix)
	)

	routesAuth.Use(middlewarepkg.AuthMiddleware())
	for _, val := range pathListAuth {
		routesAuth.Handle(val.method, val.path, val.handler)
	}

	for _, val := range pathList {
		routes.Handle(val.method, val.path, val.handler)
	}
}
