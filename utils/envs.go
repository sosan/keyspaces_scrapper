package utils

import (
	"os"
	"github.com/joho/godotenv"
)

func LoadEnvs() {
	godotenv.Load()
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

