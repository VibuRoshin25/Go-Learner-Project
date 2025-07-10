package controller

import (
	"net/http"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
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
	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}
