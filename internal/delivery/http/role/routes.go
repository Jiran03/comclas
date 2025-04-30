package roledelivery

import (
	"comclas/domain/usecase"
	rolehandler "comclas/internal/delivery/http/role/handler"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, uc *usecase.Wrapper) {
	var (
		prefix       = "/v1/role"
		handler      = rolehandler.Init(uc)
		pathListAuth = map[string]struct {
			path    string
			handler func(*gin.Context)
			method  string
		}{
			"createRole": {
				path:    "/create",
				handler: handler.CreateRole,
				method:  http.MethodPost,
			},
			"listRole": {
				path:    "",
				handler: handler.FetchRoles,
				method:  http.MethodGet,
			},
			"detailRole": {
				path:    "/:id",
				handler: handler.FetchRoleByID,
				method:  http.MethodGet,
			},
			"editRole": {
				path:    "update/:id",
				handler: handler.UpdateRole,
				method:  http.MethodPut,
			},
		}

		routesAuth = router.Group(prefix)
	)

	routesAuth.Use(middlewarepkg.AuthMiddleware())
	for _, val := range pathListAuth {
		routesAuth.Handle(val.method, val.path, val.handler)
	}
}
