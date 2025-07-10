package routes

import (
	"github.com/VibuRoshin25/Go-Learner-Project/controller"
	"github.com/VibuRoshin25/Go-Learner-Project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/health", controller.GetHealth)
	router.GET("/users", middleware.AuthMiddleware(), controller.GetUsers)
	router.GET("/users/:id", middleware.AuthMiddleware(), controller.GetById)
	router.POST("/users", middleware.AuthMiddleware(), controller.CreateUser)
	router.DELETE("/users/:id", middleware.AuthMiddleware(), controller.DeleteUser)
	router.PUT("/users/:id", middleware.AuthMiddleware(), controller.UpdateUser)
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)
}
