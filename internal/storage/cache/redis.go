package ch

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func New(zapLogger *zap.Logger) Cache {
	zapLogger = zapLogger.Named("Redis")
	host := getEnv("REDIS_HOST", "localhost:6379")
	password := getEnv("REDIS_PASSWORD", "redis")

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		zapLogger.Fatal("Error connecting to Redis")
	}
	zapLogger.Info("Connected to Redis")
	return &chClient{rd: client}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
