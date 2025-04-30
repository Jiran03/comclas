package rolehandler

import (
	"comclas/domain/entity"
	rolerequest "comclas/internal/delivery/http/role/request"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRole(c *gin.Context) {
	var (
		roleReq          rolerequest.Role
		role, roleExists = c.Get(middlewarepkg.UserRoleKey)
		userID, _        = c.Get(middlewarepkg.UserIDKey)
	)
	if err := c.ShouldBindJSON(&roleReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("role: ", role)

	if !roleExists || (roleExists && (role.(string) != "super_admin" && role.(string) != "admin")) {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	if err := h.uc.AddRole(c, entity.RoleDTO{
		Name:     entity.RoleName(roleReq.Name),
		ChangeBy: userID.(string),
	}); err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error add role")
		return
	}

	customresponse.ParseResponseSuccess(c, nil)
}
