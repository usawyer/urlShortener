package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/usawyer/urlShortener/internal/storage"
	"github.com/usawyer/urlShortener/models"
	"log"
)

const aliasLenght = 8

type Service struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) ShortenUrl(url string) (string, error) {
	if !govalidator.IsURL(url) {
		return "", errors.New("invalid URL")
	}

	ctx := context.Background()

	if data, ok := s.storage.Database.FindUrl(url, "url"); ok {
		return data.Alias, nil
	}

	hash := sha256.Sum256([]byte(url))
	alias := hex.EncodeToString(hash[:])[:aliasLenght]

	if _, ok := s.storage.Database.FindUrl(alias, "alias"); ok {
		return "", errors.New("impossible to create new alias")
	}

	newData := models.Urls{
		Alias: alias,
		Url:   url,
	}

	err := s.storage.Cache.AddCache(ctx, newData)
	if err != nil {
		log.Printf("Failed to set value in cache: %v\n\n", err)
	}

	err = s.storage.Database.AddUrl(newData)
	if err != nil {
		return "", err
	}

	return alias, nil
}

func (s *Service) ResolveUrl(alias string) (string, error) {
	if len(alias) != aliasLenght {
		return "", errors.New("invalid alias input")
	}

	ctx := context.Background()
	longUrl, err := s.storage.Cache.GetCache(ctx, alias)
	if err == nil {
		return longUrl, nil
	}

	longUrl, err = s.storage.Database.GetUrl(alias)
	if err != nil {
		return "", errors.Wrap(err, "failed to get URL from database")
	}

	err = s.storage.Cache.AddCache(ctx, models.Urls{Alias: alias, Url: longUrl})
	if err != nil {
		log.Printf("Failed to set value in cache: %v\n", err)
	}

	return longUrl, nil
}
