package helper

import (
	"os"
)

func GetEnvOrDefault(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}