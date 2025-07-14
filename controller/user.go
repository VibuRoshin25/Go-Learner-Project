package controller

import (
	"net/http"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/logs"
	"github.com/VibuRoshin25/Go-Learner-Project/models"
	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Service is up and running",
	})
}

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	if len(users) == 0 {
		logs.LogError(c, "No users available in the database")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No Users Available",
		})
		return
	}
	c.JSON(http.StatusOK, &users)
}

func GetById(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Find(&user)
	if user.CreatedAt.IsZero() {
		logs.LogError(c, "User not found with ID: "+c.Param("id"))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	if err := config.DB.Create(&user).Error; err != nil {
		logs.LogError(c, "Failed to create user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&user).Error; err != nil {
		logs.LogError(c, "Failed to delete user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		logs.LogError(c, "Failed to find user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}
	c.BindJSON(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		logs.LogError(c, "Failed to update user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, &user)
}
