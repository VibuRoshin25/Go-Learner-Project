package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/logs"
	"github.com/VibuRoshin25/Go-Learner-Project/payload"

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

		// Prepare JSON body
		payload := payload.ValidateTokenPayload{Token: token}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			logs.LogError(c, "Failed to marshal JSON: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		// Create HTTP request
		resp, err := http.Post("http://"+config.AuthHost+"/token/validate", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			logs.LogError(c, "Failed to send request to auth service: "+err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token validation failed"})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		// Check status code
		if resp.StatusCode != http.StatusOK {
			logs.LogError(c, "Auth service returned non-200 status code: "+resp.Status)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
