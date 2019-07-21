package middlewares

import (
	"net/http"
	"strings"
	"user/utils/jwt"

	"github.com/gin-gonic/gin"
)

const (
	//UserID :
	UserID = "UserID"
)

//AuthRequired :
func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenHeader := c.GetHeader("Authorization")
		tokenPart := strings.Replace(tokenHeader, "Bearer ", "", 1)

		response := gin.H{"message": "invalid token", "status": -1}

		token, err := jwt.ParseToken(tokenPart)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwt.Claims); ok {
			c.Set(UserID, claims.ID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
		}
	}
}
