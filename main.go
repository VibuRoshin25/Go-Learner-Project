package main

import (
	"GIN-CRUD-SAMPLE-PROJECT/config"
	"GIN-CRUD-SAMPLE-PROJECT/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run()
}
