package main

import (
	"GO-GIN-GORM-LEARNER-PROJECT/config"
	"GO-GIN-GORM-LEARNER-PROJECT/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run()
}
