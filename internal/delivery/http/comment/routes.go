package commentdelivery

import (
	"comclas/domain/usecase"
	commenthandler "comclas/internal/delivery/http/comment/handler"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, uc *usecase.Wrapper) {
	var (
		prefix       = "/v1/comment-class"
		handler      = commenthandler.Init(uc)
		pathListAuth = map[string]struct {
			path    string
			handler func(*gin.Context)
			method  string
		}{
			"createCommentClassification": {
				path:    "/create",
				handler: handler.CreateCommentClassification,
				method:  http.MethodPost,
			},
			"listComment": {
				path:    "",
				handler: handler.FetchComments,
				method:  http.MethodGet,
			},
			"detailComment": {
				path:    "/:id",
				handler: handler.FetchCommentByID,
				method:  http.MethodGet,
			},
			"editComment": {
				path:    "update/:id",
				handler: handler.UpdateComment,
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
