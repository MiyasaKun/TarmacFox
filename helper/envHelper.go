package helper

import (
	"log"
	"os"
)

func Getenvordefault(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		log.Printf("Environment variable not set: %s", key)
		return fallback
	}
	return value
}