package helper

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)


var ctx = context.Background()


func InitCache() (*redis.Client, error){
	// Connect to REDIS
	client := redis.NewClient(&redis.Options{
		Addr: Getenvordefault("REDIS_URL","localhost:6379"),
	})
	log.Println("Connected to Cache")
	return client, nil
}

func CloseCache(client *redis.Client) {
	if client != nil {
		client.Close()
	}
}

func GetCache(client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}