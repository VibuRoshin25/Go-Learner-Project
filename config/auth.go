package config

import "os"

var AuthHost = os.Getenv("LEARNER_AUTH_SERVICE_HOST")
