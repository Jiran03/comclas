package userhandler

import (
	userresponse "comclas/internal/delivery/http/user/response"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchUserByID(c *gin.Context) {
	var (
		_, roleExists = c.Get(middlewarepkg.UserRoleKey)
		id            = c.Param("id")
	)

	if !roleExists {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	user, err := h.uc.FetchUserByID(c, id)
	if err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error get all user")
		return
	}

	userResponse := userresponse.ToUserResponse(user)

	customresponse.ParseResponseSuccess(c, userResponse)
}
