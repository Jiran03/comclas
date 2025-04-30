package rolehandler

import (
	roleresponse "comclas/internal/delivery/http/role/response"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchRoleByID(c *gin.Context) {
	var (
		_, roleExists = c.Get(middlewarepkg.UserRoleKey)
		id, err       = strconv.Atoi(c.Param("id"))
	)

	if err != nil {
		customresponse.ParseResponseError(c, http.StatusBadRequest, err, "error get id")
		return
	}

	if !roleExists {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	role, err := h.uc.FetchRoleByID(c, id)
	if err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error get all role")
		return
	}

	roleResponse := roleresponse.ToRoleResponse(role)

	customresponse.ParseResponseSuccess(c, roleResponse)
}
