package main

import (
	"log"
	"time"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	router := gin.New()

	time.Sleep(5 * time.Second)

	config.Connect()
	routes.UserRoute(router)
	router.Run()
}
