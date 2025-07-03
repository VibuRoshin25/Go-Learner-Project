package controller

import (
	"GO-GIN-GORM-LEARNER-PROJECT/config"
	"GO-GIN-GORM-LEARNER-PROJECT/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func GetById(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Find(&user)
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
