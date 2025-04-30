package middlewarepkg

import (
	customresponse "comclas/pkg/custom_response"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	UserIDKey   = "userID"
	UserRoleKey = "userRole"
)

type Claims struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}

var jwKey = []byte(os.Getenv("SECRET_KEY_JWT"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("masuk kah? ")
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		fmt.Println("token: ", tokenString)
		cleanedToken := strings.TrimPrefix(tokenString, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(cleanedToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwKey, nil
		})

		fmt.Println("token clean:", cleanedToken)

		if err != nil || !token.Valid {
			fmt.Println(err)
			customresponse.ParseResponseError(c, http.StatusUnauthorized, err, "Invalid token")
			c.Abort()
			return
		}
		fmt.Println("claims: ", claims)

		if claims.UserID == "" {
			customresponse.ParseResponseError(c, http.StatusUnauthorized, err, "invalid user id")
			c.Abort()
			return
		}

		// Pass user ID to the context
		c.Set(UserIDKey, claims.UserID)
		c.Set(UserRoleKey, claims.UserRole)
		c.Next()
	}
}
