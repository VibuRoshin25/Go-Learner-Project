package main

import (
	"time"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	time.Sleep(5 * time.Second)

	config.Connect()
	routes.UserRoute(router)
	router.Run()
}
