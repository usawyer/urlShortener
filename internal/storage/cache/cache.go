package ch

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/usawyer/urlShortener/models"
)

const cacheDuration = 6 * time.Hour

type Cache interface {
	AddCache(context.Context, models.Urls) error
	GetCache(context.Context, string) (string, error)
}

type chClient struct {
	rd *redis.Client
}

func (c *chClient) AddCache(ctx context.Context, urls models.Urls) error {
	return c.rd.Set(ctx, urls.Alias, urls.Url, cacheDuration).Err()
}

func (c *chClient) GetCache(ctx context.Context, key string) (string, error) {
	return c.rd.Get(ctx, key).Result()
}
