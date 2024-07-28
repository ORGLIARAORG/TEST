package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jarqvi/gin-crud/utils"
)

func Authenticate(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	tokenParts := strings.Split(auth, " ")
	if len(tokenParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	userID, err := utils.VerifyToken(tokenParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	c.Set("userID", userID)

	c.Next()
}
