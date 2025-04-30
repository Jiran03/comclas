package rolehandler

import (
	roleresponse "comclas/internal/delivery/http/role/response"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchRoles(c *gin.Context) {
	var (
		role, roleExists = c.Get(middlewarepkg.UserRoleKey)
	)

	if !roleExists || (roleExists && (role.(string) != "super_admin" && role.(string) != "admin")) {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	roles, err := h.uc.FetchRoles(c)
	if err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error get all role")
		return
	}

	roleResponse := roleresponse.ToListRoleResponse(roles)

	customresponse.ParseResponseSuccess(c, roleResponse)
}
