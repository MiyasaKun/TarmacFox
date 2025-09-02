package helper

import (
	"tarmac-fox/helper"

	"github.com/redis/go-redis/v9"
)

func InitCache() (cl redis.Client, err error){
	// Connect to REDIS
	client := redis.NewClient(&redis.Options{
		Addr: helper.Getenvordefault("REDIS_URL","localhost:6379"),
	})
	return client,nil
}

func CloseCache() {

}