package rolehandler

import (
	"comclas/domain/entity"
	rolerequest "comclas/internal/delivery/http/role/request"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateRole(c *gin.Context) {
	var (
		roleReq          rolerequest.Role
		id, err          = strconv.Atoi(c.Param("id"))
		role, roleExists = c.Get(middlewarepkg.UserRoleKey)
		userID, _        = c.Get(middlewarepkg.UserIDKey)
	)

	if err != nil {
		customresponse.ParseResponseError(c, http.StatusBadRequest, err, "error get id")
		return
	}

	if err = c.ShouldBindJSON(&roleReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !roleExists || (roleExists && (role.(string) != "super_admin" && role.(string) != "admin")) {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	if err := h.uc.UpdateRole(c, id, entity.RoleDTO{
		Name:     entity.RoleName(roleReq.Name),
		ChangeBy: userID.(string),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customresponse.ParseResponseSuccess(c, nil)
}
