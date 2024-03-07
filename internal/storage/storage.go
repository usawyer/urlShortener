package storage

import (
	ch "github.com/usawyer/urlShortener/internal/storage/cache"
	db "github.com/usawyer/urlShortener/internal/storage/database"
)

type Storage struct {
	database *db.Database
	cache    *ch.Cache
}

func New(db *db.Database, ch *ch.Cache) *Storage {
	return &Storage{
		database: db,
		cache:    ch,
	}
}
