package main

import (
	"log"
	"time"

	"vibrox-core/config"
	"vibrox-core/routes"

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

	logClientConn, err := config.InitLoggerClient()
	if err != nil {
		log.Fatal("Failed to initialize logger client:", err)
	}

	authClientConn, err := config.InitAuthClient()
	if err != nil {
		log.Fatal("Failed to initialize auth client:", err)
	}

	defer func() {
		logClientConn.Close()
		authClientConn.Close()
	}()

	routes.UserRoute(router)
	router.Run()
}
