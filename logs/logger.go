package logs

import (
	"context"
	"fmt"

	"github.com/VibuRoshin25/Go-Learner-Project/config"
	"github.com/VibuRoshin25/Go-Learner-Project/proto/logger"
)

func LogError(ctx context.Context, message string) error {

	resp, err := config.LogClient.Log(ctx, &logger.LogRequest{
		Message: message,
		Service: "Learner",
		Level:   "ERROR",
	})
	if err != nil {
		return err
	} else if !resp.Success {
		return fmt.Errorf("logging error: %s", resp.Err)
	}

	return nil
}

func LogInfo(ctx context.Context, message string) error {
	resp, err := config.LogClient.Log(ctx, &logger.LogRequest{
		Message: message,
		Service: "Learner",
		Level:   "INFO",
	})
	if err != nil {
		return err
	} else if !resp.Success {
		return fmt.Errorf("logging error: %s", resp.Err)
	}

	return nil
}
