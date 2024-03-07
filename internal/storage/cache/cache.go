package ch

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/usawyer/urlShortener/models"
	"log"
	"os"
	"time"
)

// const cacheDuration = 6 * time.Hour
const cacheDuration = 2 * time.Minute

type Cache interface {
	AddCache(context.Context, models.Urls) error
	GetCache(context.Context, string) (string, error)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

type chClient struct {
	rd *redis.Client
}

func New() Cache {
	host := getEnv("REDIS_HOST", "localhost:6379")
	password := getEnv("REDIS_PASSWORD", "redis")
	//db, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		//DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		fmt.Println("Error connecting to Redis:", err)
	}

	log.Println("REDDIS IS PINGED")

	return &chClient{rd: client}
}

func (c *chClient) AddCache(ctx context.Context, urls models.Urls) error {
	return c.rd.Set(ctx, urls.Alias, urls.Url, cacheDuration).Err()
}

func (c *chClient) GetCache(ctx context.Context, key string) (string, error) {
	return c.rd.Get(ctx, key).Result()
}
