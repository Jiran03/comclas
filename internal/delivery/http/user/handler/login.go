package userhandler

import (
	userrequest "comclas/internal/delivery/http/user/request"
	customresponse "comclas/pkg/custom_response"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(os.Getenv("SECRET_KEY_JWT"))

type Token struct {
	Token string `json:"token"`
}

func (h *Handler) LoginUser(c *gin.Context) {
	var loginRequest userrequest.Login

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		fmt.Println("req bang:c", loginRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	fmt.Println("req: ", loginRequest)

	// Dummy authentication logic
	if loginRequest.Username != "admin" || loginRequest.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	id := c.Param("id")
	fmt.Println("id: ", id)

	// Generate JWT token
	expirationTime := time.Now().Add(360 * time.Minute)
	claims := &Claims{
		UserID:   id, // Dummy user ID
		UserRole: "super_admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	tokenStruct := &Token{
		Token: tokenString,
	}

	customresponse.ParseResponseSuccess(c, *tokenStruct)
}
