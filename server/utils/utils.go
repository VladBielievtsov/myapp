package utils

import "os"

func IsEnvProd() bool {
	return os.Getenv("ENV") == "production"
}

func AppEnv() string {
	return os.Getenv("ENV")
}
