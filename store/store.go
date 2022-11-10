package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Store struct {
	client *redis.Client
}

var (
	ctx   = context.Background()
	store = &Store{}
)

func InitStoreClient(rdsHost string, rdsPwd string) *Store {
	client := redis.NewClient(&redis.Options{
		Addr:     rdsHost,
		Password: rdsPwd,
		DB:       1,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalln("Failed to connect to redis")
	}

	store.client = client
	return store
}

func CreateUrl(shortcode string, url string) {
	err := store.client.Set(ctx, shortcode, url, -1).Err()
	if err != nil {
		log.Fatalf("Failed to store url and shortcode - %s", err)
	}
	log.Println("Stored shortcode and url")
}

func RetrieveUrl(shortcode string) string {
	url, err := store.client.Get(ctx, shortcode).Result()
	if err != nil {
		return err.Error()
	}

	return url
}
