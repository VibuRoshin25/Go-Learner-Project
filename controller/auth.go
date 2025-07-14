package controller

import (
	"net/http"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/logs"
	"github.com/VibuRoshin25/Go-Learner-Project/models"
	"github.com/VibuRoshin25/Go-Learner-Project/payload"
	"github.com/VibuRoshin25/Go-Learner-Project/proto/auth"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		logs.LogError(c, "Failed to bind JSON: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the request payload
	if req.Email == "" || req.Password == "" {
		logs.LogError(c, "Email and Password are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Password are required"})
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		logs.LogError(c, "User already exists with email: "+req.Email)
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Create new user
	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		logs.LogError(c, "Failed to create user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	logs.LogInfo(c, "User created successfully with email: "+req.Email)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func SignIn(c *gin.Context) {
	var req payload.SignInPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		logs.LogError(c, "Failed to bind JSON: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the request payload
	if req.Email == "" || req.Password == "" {
		logs.LogError(c, "Email and Password are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Password are required"})
		return
	}

	// Check Email Password with DB
	var user models.User
	if err := config.DB.Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logs.LogError(c, "Invalid email or password for email: "+req.Email)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		logs.LogError(c, "Database error: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	authResp, err := config.AuthClient.GenerateToken(c, &auth.GenerateTokenRequest{
		Email:  user.Email,
		UserId: int64(user.ID),
	})
	if err != nil {
		logs.LogError(c, "Failed to generate token: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	token := authResp.Token
	if token == "" {
		logs.LogError(c, "Token not found in auth service response")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not found in response"})
		return
	}

	logs.LogInfo(c, "User signed in successfully with email: "+req.Email)
	c.JSON(http.StatusOK, gin.H{
		"message": "SignIn successful",
		"token":   token,
	})
}
