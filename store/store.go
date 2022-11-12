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
		log.Fatalf("Failed to connect to redis- %s", err)
	}

	store.client = client
	return store
}

func CreateUrl(shortcode string, url string) error {
	err := store.client.Set(ctx, shortcode, url, -1).Err()
	if err != nil {
		log.Panicf("Failed to store url and shortcode - %s", err)
		return err
	}
	return nil
}

func RetrieveUrl(shortcode string) (string, error) {
	url, err := store.client.Get(ctx, shortcode).Result()
	if err != nil {
		log.Printf("Failed to retrieve by shortcode - %s", shortcode)
	}

	return url, err
}
