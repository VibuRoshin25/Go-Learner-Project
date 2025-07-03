package main

import (
	"GO-GIN-GORM-LEARNER-PROJECT/config"
	"GO-GIN-GORM-LEARNER-PROJECT/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	time.Sleep(5 * time.Second)

	config.Connect()
	routes.UserRoute(router)
	router.Run()
}
