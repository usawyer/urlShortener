package ch

import "github.com/redis/go-redis/v9"

type Cache interface {
	AddCache(string) (string, error)
	GetCache(string) (string, error)
}

type chClient struct {
	rd *redis.Client
}

func (c *chClient) AddCache(string) (string, error) {
	return "", nil
}

func (c *chClient) GetCache(string) (string, error) {
	return "", nil
}
