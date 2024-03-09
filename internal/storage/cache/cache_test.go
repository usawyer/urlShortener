package ch

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/usawyer/urlShortener/models"
	"testing"
)

func TestChClient(t *testing.T) {
	srv, err := miniredis.Run()
	if err != nil {
		t.Fatalf("error starting miniredis: %v", err)
	}
	defer srv.Close()

	client := redis.NewClient(&redis.Options{
		Addr: srv.Addr(),
	})
	defer client.Close()

	cache := &chClient{rd: client}

	testUrl := models.Urls{
		Alias: "test",
		Url:   "https://example.com",
	}

	err = cache.AddCache(context.Background(), testUrl)
	assert.NoError(t, err)

	cachedURL, err := cache.GetCache(context.Background(), testUrl.Alias)
	assert.NoError(t, err)
	assert.Equal(t, testUrl.Url, cachedURL)
}
