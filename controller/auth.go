package controller

import (
	"io"
	"net/http"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/models"
	"github.com/VibuRoshin25/Go-Learner-Project/payload"
	"gorm.io/gorm"

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

	// Validate the request payload
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Password are required"})
		return
	}

	// Check Email Password with DB
	var user models.User
	if err := config.DB.Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Prepare payload for auth service
	requestPayload, _ := json.Marshal(payload.GenerateTokenPayload{
		Email: user.Email,
		Id:    user.ID,
	})
	resp, err := http.Post(config.AuthHost+"/token/generate", "application/json", bytes.NewBuffer(requestPayload))
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
