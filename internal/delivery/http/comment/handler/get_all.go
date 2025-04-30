package commenthandler

import (
	commentresponse "comclas/internal/delivery/http/comment/response"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchComments(c *gin.Context) {
	var (
		role, roleExists = c.Get(middlewarepkg.UserRoleKey)
	)

	if !roleExists || (roleExists && (role.(string) != "super_admin" && role.(string) != "admin")) {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	comments, err := h.uc.FetchComments(c)
	if err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error get all comment")
		return
	}

	commentResponse := commentresponse.ToListCommentResponse(comments)

	customresponse.ParseResponseSuccess(c, commentResponse)
}
