package middleware

import (
	"net/http"

	"vibrox-core/config"
	"vibrox-core/logs"
	"vibrox-core/proto/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			logs.LogError(c, "Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		resp, err := config.AuthClient.ValidateToken(c, &auth.ValidateTokenRequest{Token: token})
		if err != nil {
			logs.LogError(c, "Failed to validate token: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if !resp.Valid {
			logs.LogError(c, "Invalid token: "+resp.Error)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
