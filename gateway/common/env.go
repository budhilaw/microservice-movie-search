package common

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var Env env

type env struct {
	DbUser         string
	DbPassword     string
	DbName         string
	DbHost         string
	DbPort         int
	AppPort        int
	GrpcAdress     string
	GrpcPort       int
	ContextTimeout int
}

func LoadEnv() {
	godotenv.Load()

	Env.DbUser = os.Getenv("DB_USER")
	Env.DbPassword = os.Getenv("DB_PASS")
	Env.DbName = os.Getenv("DB_NAME")
	Env.DbHost = os.Getenv("DB_HOST")
	Env.DbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	Env.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	Env.GrpcAdress = os.Getenv("GRPC_ADDRESS")
	Env.GrpcPort, _ = strconv.Atoi(os.Getenv("GRPC_PORT"))
	Env.ContextTimeout, _ = strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
}
