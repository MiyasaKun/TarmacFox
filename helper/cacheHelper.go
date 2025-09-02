package helper

import (
	"github.com/redis/go-redis/v9"
)

func InitCache() (*redis.Client, error){
	// Connect to REDIS
	client := redis.NewClient(&redis.Options{
		Addr: Getenvordefault("REDIS_URL","localhost:6379"),
	})
	return client, nil
}

func CloseCache(client *redis.Client) {
	if client != nil {
		client.Close()
	}
}