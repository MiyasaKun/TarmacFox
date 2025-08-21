package helper

import (
	"log"
	"os"
)

func Getenvordefault(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		log.Panic("Environment variable not set: ", key)
		return fallback
	}
	return value
}