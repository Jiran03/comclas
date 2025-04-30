package userhandler

import (
	"comclas/domain/entity"
	userrequest "comclas/internal/delivery/http/user/request"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var (
		user             userrequest.User
		role, roleExists = c.Get(middlewarepkg.UserRoleKey)
		userID, _        = c.Get(middlewarepkg.UserIDKey)
	)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("role: ", role)

	if !roleExists || (roleExists && (role.(string) != "super_admin" && role.(string) != "admin")) {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	if err := h.uc.AddUser(c, entity.UserDTO{
		Name:     user.Name,
		NIK:      user.NIK,
		RoleID:   user.RoleID,
		Username: user.Username,
		Password: user.Password,
		ChangeBy: userID.(string),
	}); err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error add user")
		return
	}

	customresponse.ParseResponseSuccess(c, nil)
}
