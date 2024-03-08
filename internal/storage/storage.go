package storage

import (
	ch "github.com/usawyer/urlShortener/internal/storage/cache"
	db "github.com/usawyer/urlShortener/internal/storage/database"
)

type Storage struct {
	Database db.Database
	Cache    ch.Cache
}

func New(ch ch.Cache, db db.Database) *Storage {
	return &Storage{
		Database: db,
		Cache:    ch,
	}
}
