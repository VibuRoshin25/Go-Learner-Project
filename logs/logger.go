package logs

import (
	"context"
	"log"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/proto/logger"
)

func LogError(ctx context.Context, message string) {

	resp, err := config.LogClient.Log(ctx, &logger.LogRequest{
		Message: message,
		Service: "Learner",
		Level:   "ERROR",
	})
	if err != nil {
		log.Println("Failed to log error: ", err)
	} else if !resp.Success {
		log.Println("Logger error: ", resp.Err)
	}
}

func LogInfo(ctx context.Context, message string) {
	resp, err := config.LogClient.Log(ctx, &logger.LogRequest{
		Message: message,
		Service: "Learner",
		Level:   "INFO",
	})
	if err != nil {
		log.Println("Failed to log error: ", err)
	} else if !resp.Success {
		log.Println("Logger error: ", resp.Err)
	}
}
