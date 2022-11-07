package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Database struct {
	client *redis.Client
}

var (
	ctx      = context.Background()
	database = &Database{}
)

func InitRedisClient(rdsHost string, rdsPwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     rdsHost,
		Password: rdsPwd,
		DB:       1,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalln("Failed to connect to redis")
	}

	database.client = client
	log.Println("Connected to redis")
}

func CreateUrl(shortcode string, url string) {
	err := database.client.Set(ctx, shortcode, url, 1)
	if err != nil {
		return
	}
	log.Println("Stored shortcode and url")
}

func RetrieveUrl(shortcode string) string {
	result, err := database.client.Get(ctx, shortcode).Result()
	if err != nil {
		return err.Error()
	}

	return result
}
