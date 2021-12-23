package common

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var Env env

type env struct {
	AppPort int
}

func LoadEnv() {
	godotenv.Load()

	Env.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
}
