package config

import (
	"os"

	"github.com/VibuRoshin25/Go-Learner-Project/proto/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var LoggerHost = os.Getenv("LOGGER_HOST")

var LogClient logger.LoggerClient

func InitLoggerClient() (*grpc.ClientConn, error) {

	conn, err := grpc.NewClient(LoggerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	LogClient = logger.NewLoggerClient(conn)

	return conn, nil
}
