package routes

import (
	"github.com/VibuRoshin25/Go-Learner-Project/controller"
	"github.com/VibuRoshin25/Go-Learner-Project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/health", controller.GetHealth)
	router.GET("/", controller.GetUsers)
	router.GET("/:id", controller.GetById)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", middleware.AuthMiddleware(), controller.DeleteUser)
	router.PUT("/:id", middleware.AuthMiddleware(), controller.UpdateUser)
}
