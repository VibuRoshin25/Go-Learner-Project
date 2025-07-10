package controller

import (
	"io"
	"net/http"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/payload"

	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var req payload.SignInPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Prepare payload for auth service
	requestPayload, _ := json.Marshal(req)
	resp, err := http.Post(config.AuthHost+"/signin", "application/json", bytes.NewBuffer(requestPayload))
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var authResp payload.ValidateTokenPayload
	if err := json.Unmarshal(body, &authResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from auth service"})
		return
	}

	token := authResp.Token
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not found in response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SignIn successful",
		"token":   token,
	})
}
