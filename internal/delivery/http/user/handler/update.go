package userhandler

import (
	"comclas/domain/entity"
	userrequest "comclas/internal/delivery/http/user/request"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUser(c *gin.Context) {
	var (
		user             userrequest.User
		id               = c.Param("id")
		role, roleExists = c.Get(middlewarepkg.UserRoleKey)
		userID, _        = c.Get(middlewarepkg.UserIDKey)
	)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !roleExists || (roleExists && (role.(string) != "super_admin" && role.(string) != "admin")) {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	if err := h.uc.UpdateUser(c, id, entity.UserDTO{
		Name:     user.Name,
		NIK:      user.NIK,
		RoleID:   user.RoleID,
		Username: user.Username,
		Password: user.Password,
		ChangeBy: userID.(string),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customresponse.ParseResponseSuccess(c, nil)
}
