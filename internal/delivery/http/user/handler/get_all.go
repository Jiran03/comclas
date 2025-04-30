package userhandler

import (
	userresponse "comclas/internal/delivery/http/user/response"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchUsers(c *gin.Context) {
	var (
		role, roleExists = c.Get(middlewarepkg.UserRoleKey)
	)

	if !roleExists || (roleExists && (role.(string) != "super_admin" && role.(string) != "admin")) {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	users, err := h.uc.FetchUsers(c)
	if err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error get all user")
		return
	}

	userResponse := userresponse.ToListUserResponse(users)

	customresponse.ParseResponseSuccess(c, userResponse)
}
