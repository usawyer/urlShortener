package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/usawyer/urlShortener/internal/storage"
	"github.com/usawyer/urlShortener/models"
	"go.uber.org/zap"
)

const aliasLenght = 8

type Service struct {
	storage *storage.Storage
	logger  *zap.Logger
}

func New(storage *storage.Storage, logger *zap.Logger) *Service {
	logger = logger.Named("Service")
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) ShortenUrl(ctx context.Context, url string) (string, error) {
	if !govalidator.IsURL(url) {
		return "", errors.New("invalid URL")
	}

	url = standardizeUrl(url)

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
		s.logger.Warn("Failed to set value in cache", zap.Error(err))
	}

	err = s.storage.Database.AddUrl(newData)
	if err != nil {
		return "", err
	}

	return alias, nil
}

func (s *Service) ResolveUrl(ctx context.Context, alias string) (string, error) {
	if len(alias) != aliasLenght {
		return "", errors.New("invalid alias input")
	}

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
		s.logger.Warn("Failed to set value in cache", zap.Error(err))
	}

	return longUrl, nil
}

func standardizeUrl(url string) string {
	if url[:4] != "http" {
		genURL := strings.Replace(url, "www.", "", 1)
		trimGenURL := strings.TrimSuffix(genURL, "/")
		return "http://" + trimGenURL
	}

	if url[:5] == "https" {
		genURL := strings.Replace(url, "www.", "", 1)
		trimGenURL := strings.TrimSuffix(genURL, "/")
		return strings.Replace(trimGenURL, "https", "http", 1)
	}

	genURL := strings.Replace(url, "www.", "", 1)
	trimGenURL := strings.TrimSuffix(genURL, "/")
	return trimGenURL
}
