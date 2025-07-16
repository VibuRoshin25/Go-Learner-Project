package config

import (
	"os"

	"vibrox-core/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthHost = os.Getenv("AUTH_HOST")

var AuthClient auth.TokenClient

func InitAuthClient() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(AuthHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	AuthClient = auth.NewTokenClient(conn)

	return conn, nil
}
