package utils

import (
	"os"

	"github.com/braejan/platzi-go-rest-websockets/entities"
)

const (
	PORT_ENV_KEY       = "PORT"
	JWT_SECRET_ENV_KEY = "JWT_SECRET"
	DATABASE_URL       = "DATABASE_URL"
)

func GetEnvValue(key string) (value string) {
	value = os.Getenv(key)
	return
}

func GetConfigEnvVariables() (confEnvVariables *entities.ConfEnvVariables) {
	confEnvVariables = &entities.ConfEnvVariables{
		Port:        GetEnvValue(PORT_ENV_KEY),
		JWTSecret:   GetEnvValue(JWT_SECRET_ENV_KEY),
		DatabaseUrl: GetEnvValue(DATABASE_URL),
	}
	return
}
