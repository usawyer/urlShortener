package service

import (
	"crypto/sha256"
	"encoding/hex"
)

type Service struct {
	storage int
}

func New(storage int) *Service {
	return &Service{storage: storage}
}

func (s *Service) ShortenUrl(url string) (string, error) {

	// сделать проверку, есть ли url  уже в сторадже или bd

	hash := sha256.Sum256([]byte(url))
	shortCode := hex.EncodeToString(hash[:])[:8]

	// проверка, есть ли такой шорт уже в кэше или в бд
	// вставить в кэш, в бд

	return shortCode, nil
}

func (s *Service) ResolveUrl(alias string) (string, error) {
	// вернуть из кэша, если нет - из бд
	//

	return "", nil
}
