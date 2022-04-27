package server

import (
	"fmt"
)

func ValidateServerConfig(configuration *Config) (err error) {
	err = ValidateEmptyString(configuration.Port, "port")
	if err != nil {
		return
	}
	err = ValidateEmptyString(configuration.JWTSecret, "jwt-secret")
	if err != nil {
		return
	}
	err = ValidateEmptyString(configuration.DatabaseUrl, "database-url")
	if err != nil {
		return
	}
	return
}

func ValidateEmptyString(s string, params interface{}) (err error) {
	if s == "" {
		err = fmt.Errorf("%s must be not empty", params)
	}
	return
}
